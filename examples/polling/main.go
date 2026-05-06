package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync/atomic"
	"time"

	"github.com/aws-contrib/aws-cli/awss3"
	"github.com/urfave/cli/v3"
)

type Config struct {
	DBHost   string `json:"db_host"`
	DBPort   int    `json:"db_port"`
	LogLevel string `json:"log_level"`
}

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Sources: cli.NewValueSourceChain(
					awss3.Object("my-bucket", "config.json"),
				),
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var current atomic.Pointer[Config]

			var cfg Config
			if err := json.Unmarshal([]byte(cmd.String("config")), &cfg); err != nil {
				return err
			}
			current.Store(&cfg)

			go poll(ctx, cmd, &current, time.Minute)

			// Use current.Load() wherever the latest config is needed.
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func poll(ctx context.Context, cmd *cli.Command, current *atomic.Pointer[Config], every time.Duration) {
	ticker := time.NewTicker(every)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for _, f := range cmd.Flags {
				sf, ok := f.(*cli.StringFlag)
				if !ok {
					continue
				}
				raw, ok := sf.Sources.Lookup()
				if !ok {
					continue
				}
				var next Config
				if err := json.Unmarshal([]byte(raw), &next); err != nil {
					log.Printf("refresh %s: %v", sf.Name, err)
					continue
				}
				current.Store(&next)
			}
		}
	}
}
