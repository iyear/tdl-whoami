package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-faster/errors"
	"github.com/iyear/tdl/extension"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	conf := zap.NewDevelopmentConfig()
	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	extension.New(extension.Options{
		UpdateHandler: nil,
		Middlewares:   nil,
		// override default logger, which write to extension data dir
		Logger: zap.Must(conf.Build(zap.AddCaller())),
	})(func(ctx context.Context, e *extension.Extension) error {
		e.Log().Info("start",
			zap.String("name", e.Name()),
			zap.String("data_dir", e.Config().DataDir),
			zap.String("proxy", e.Config().Proxy),
			zap.Bool("debug", e.Config().Debug))

		// and call the API via e.Client().API()
		self, err := e.Client().Self(ctx)
		if err != nil {
			return errors.Wrap(err, "get self")
		}

		e.Log().Info("get self",
			zap.Int64("id", self.ID),
			zap.String("username", self.Username))

		b, err := json.MarshalIndent(self, "", "  ")
		if err != nil {
			return errors.Wrap(err, "marshal self to json")
		}

		fmt.Println(string(b))

		return nil
	})
}
