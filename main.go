package main

import (
	"fmt"
	"github.com/Thumbscrew/hashgo/pkg/hasher"
	"github.com/schollz/progressbar/v3"
	"os"
	"slices"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	Algorithms = []string{"md5", "sha1", "sha224", "sha256", "sha384", "sha512"}
)

func main() {
	app := &cli.App{
		Name:  "hashgo",
		Usage: "Hash files (with a progress bar)",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "path",
				Aliases:  []string{"f"},
				EnvVars:  []string{"HASHGO_PATH"},
				Required: true,
				Usage:    "path to file to hash",
			},
			&cli.StringFlag{
				Name:     "algorithm",
				Aliases:  []string{"a"},
				EnvVars:  []string{"HASHGO_ALGORITHM"},
				Required: true,
				Usage:    "hashing algorithm to use (md5, sha1, sha224, sha256, sha384, sha512)",
				Action: func(_ *cli.Context, s string) error {
					if !slices.Contains(Algorithms, strings.ToLower(s)) {
						return fmt.Errorf("invalid hashing algorithm: %s", s)
					}

					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			var (
				path = ctx.Path("path")
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

			h, err := c.Sha256(f)
			if err != nil {
				return err
			}
			_ = bar.Close()

			fmt.Println(h)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("fatal: %s\n", err)
		os.Exit(1)
	}
}
