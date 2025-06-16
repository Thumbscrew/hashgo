package main

import (
	"context"
	"fmt"
	"github.com/Thumbscrew/hashgo/pkg/hasher"
	"github.com/schollz/progressbar/v3"
	"net/mail"
	"os"
	"slices"
	"strings"

	"github.com/urfave/cli/v3"
)

var (
	Version    = "dev"
	Algorithms = []string{"md5", "sha1", "sha224", "sha256", "sha384", "sha512"}
)

func main() {
	app := &cli.Command{
		Name:    "hashgo",
		Usage:   "Hash files (with a progress bar)",
		Version: Version,
		Authors: []any{
			mail.Address{
				Name:    "Thumbscrew",
				Address: "thumbscrw@pm.me",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "path",
				TakesFile: true,
				Aliases:   []string{"f"},
				Sources:   cli.EnvVars("HASHGO_PATH"),
				Required:  true,
				Usage:     "path to file to hash",
			},
			&cli.StringFlag{
				Name:     "algorithm",
				Aliases:  []string{"a"},
				Sources:  cli.EnvVars("HASHGO_ALGORITHM"),
				Required: true,
				Usage:    "hashing algorithm to use (md5, sha1, sha224, sha256, sha384, sha512)",
				Action: func(_ context.Context, _ *cli.Command, s string) error {
					if !slices.Contains(Algorithms, strings.ToLower(s)) {
						return fmt.Errorf("invalid hashing algorithm: %s", s)
					}

					return nil
				},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var (
				path      = cmd.String("path")
				algorithm = strings.ToLower(cmd.String("algorithm"))
			)

			fi, err := os.Stat(path)
			if err != nil {
				return err
			}

			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer func(f *os.File) {
				_ = f.Close()
			}(f)

			bar := progressbar.DefaultBytes(fi.Size(), fmt.Sprintf("Hashing file %s", path))
			c := &hasher.Client{
				Progress: bar,
			}

			var h string
			switch algorithm {
			case "md5":
				h, err = c.Md5(f)
			case "sha1":
				h, err = c.Sha1(f)
			case "sha224":
				h, err = c.Sha224(f)
			case "sha256":
				h, err = c.Sha256(f)
			case "sha384":
				h, err = c.Sha384(f)
			case "sha512":
				h, err = c.Sha512(f)
			default:
				return fmt.Errorf("invalid hashing algorithm: %s", algorithm)
			}
			if err != nil {
				return err
			}

			_ = bar.Close()

			fmt.Println(h)

			return nil
		},
	}

	if err := app.Run(context.TODO(), os.Args); err != nil {
		fmt.Printf("fatal: %s\n", err)
		os.Exit(1)
	}
}
