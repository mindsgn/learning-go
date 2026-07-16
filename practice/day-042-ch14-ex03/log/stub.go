package log

import (
	"context"
	"net/http"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

func ContextWithLevel(ctx context.Context, level Level) context.Context { return ctx }
func LevelFromContext(ctx context.Context) (Level, bool)                { return "", false }
func Log(ctx context.Context, level Level, message string)              {}
func Middleware(h http.Handler) http.Handler                            { return h }
