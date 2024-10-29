package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/gotd/td/telegram/tljson"
	"github.com/iyear/tdl/extension"
	"github.com/mattn/go-colorable"
	"github.com/neilotoole/jsoncolor"
	"github.com/spf13/pflag"
)

var verbose bool

func init() {
	pflag.BoolVarP(&verbose, "verbose", "v", false, "verbose output including User, Config, AppConfig")

	pflag.Parse()
}

func main() {
	extension.New(extension.Options{
		UpdateHandler: nil,
		Middlewares:   nil,
		Logger:        nil,
	})(func(ctx context.Context, e *extension.Extension) error {
		self, err := e.Client().Self(ctx)
		if err != nil {
			return errors.Wrap(err, "get self")
		}

		if !verbose {
			color.Green("You are %s. ID: %d", name(self.FirstName, self.LastName, self.Username), self.ID)
			return nil
		}

		// verbose output
		output("User:", self, "")

		appConfig, err := e.Client().API().HelpGetAppConfig(ctx, 0)
		if err != nil {
			return errors.Wrap(err, "get auth status")
		}

		if cfg, ok := appConfig.AsModified(); ok {
			appCfg := &jx.Encoder{}
			tljson.Encode(cfg.Config, appCfg)

			m := make(map[string]interface{})
			if err = json.Unmarshal(appCfg.Bytes(), &m); err != nil {
				return errors.Wrap(err, "unmarshal app config")
			}
			output("AppConfig:", m, "see https://core.telegram.org/api/config#client-configuration for more info on the result.")
		}

		config, err := e.Client().API().HelpGetConfig(ctx)
		if err != nil {
			return errors.Wrap(err, "get config")
		}

		output("Config:", config, "")

		return nil
	})
}

func output(header string, v any, footer string) {
	if header != "" {
		color.Blue(header)
	}

	enc := jsoncolor.NewEncoder(colorable.NewColorable(os.Stdout))

	clrs := jsoncolor.DefaultColors()
	clrs.Key = []byte("\x1b[35m")  // magenta
	clrs.Bool = []byte("\x1b[33m") // yellow
	enc.SetColors(clrs)
	enc.SetIndent("", "  ")

	if err := enc.Encode(v); err != nil {
		fmt.Printf("%+v\n", v) // fallback
	}

	if footer != "" {
		color.Blue(footer)
	}

	fmt.Println()
}

func name(first, last, username string) string {
	if first == "" && last == "" {
		return username
	}

	if first == "" {
		return last
	}

	if last == "" {
		return first
	}

	return fmt.Sprintf("%s %s", first, last)
}
