package logging

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"

	"trader/internal/config"
)

type ConsoleHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type ConsoleHandler struct {
	slog.Handler
	l     *log.Logger
	attrs []slog.Attr
}

func (h *ConsoleHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"
	msg := color.CyanString(r.Message)

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.GreenString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
		msg = color.YellowString(r.Message)
	case slog.LevelError:
		level = color.RedString(level)
		msg = color.RedString(r.Message)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})
	var b []byte

	b, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[MST 02-01-2006 15:04:05.000]")

	h.l.Println(timeStr, level, msg, string(b))

	return nil
}

func (h *ConsoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ConsoleHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

func NewConsoleHandler(out io.Writer, opts ConsoleHandlerOptions) *ConsoleHandler {
	h := &ConsoleHandler{
		Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}

func SetupLogging(loggerCfg *config.Logger) *slog.Logger {

	var level slog.Level

	switch loggerCfg.Level {
	case "DEBUG":
		level = slog.LevelDebug
	case "INFO":
		level = slog.LevelInfo
	case "WARN":
		level = slog.LevelWarn
	default:
		level = slog.LevelInfo
	}

	opts := ConsoleHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: level,
		},
	}
	var logger *slog.Logger
	if loggerCfg.Json {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &opts.SlogOpts))
	} else {
		logger = slog.New(NewConsoleHandler(os.Stdout, opts))
	}
	slog.SetDefault(logger)
	return logger
}
