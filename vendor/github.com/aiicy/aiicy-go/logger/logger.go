//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that can
// be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, August 2019
//

package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var S *zap.SugaredLogger

type Logger = zap.SugaredLogger

type Field = zap.Field

func Error(err error) Field {
	return zap.Error(err)
}

func String(key string, val string) Field {
	return zap.String(key, val)
}

func init() {
	// Print log when start
	S = New(LogConfig{Level: "debug"})
}

// NewEncoderConfig creates logger config for debug mode
func NewEncoderConfig() zapcore.EncoderConfig {
	conf := zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	return conf
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// ParseLevel parses string to zap level
func ParseLevel(level string) (zapcore.Level, error) {
	switch level {
	case "debug":
		return zap.DebugLevel, nil
	case "info":
		return zap.InfoLevel, nil
	case "warn":
		return zap.WarnLevel, nil
	case "error":
		return zap.ErrorLevel, nil
	default:
		return zap.WarnLevel, fmt.Errorf("level %s not exist", level)
	}
}

// New create a new Sugared logger
func New(c LogConfig, fields ...string) *zap.SugaredLogger {
	var (
		format zapcore.Encoder
		write  zapcore.WriteSyncer
	)
	logLevel, err := ParseLevel(c.Level)
	if err != nil {
		S.Warnf("failed to parse log level (%s), use default level (info)", c.Level)
	}

	if c.Format == "json" {
		format = zapcore.NewJSONEncoder(NewEncoderConfig())
	} else {
		format = zapcore.NewConsoleEncoder(NewEncoderConfig())
	}

	if c.Mode == "file" {
		write = zapcore.AddSync(&lumberjack.Logger{
			Filename:   c.Path,
			MaxAge:     c.Age.Max,  //days
			MaxSize:    c.Size.Max, // megabytes
			MaxBackups: c.Backup.Max,
		})
	} else {
		write = os.Stdout
	}
	core := zapcore.NewCore(
		format,
		write,
		logLevel,
	)
	var options []zap.Option
	if len(fields) > 0 && len(fields)%2 == 0 {
		zapFields := []zap.Field{}
		for index := 0; index < len(fields)-1; index = index + 2 {
			zapFields = append(zapFields, zap.String(fields[index], fields[index+1]))
		}
		options = append(options, zap.Fields(zapFields...))
	}
	if logLevel == zap.DebugLevel {
		options = append(options, zap.AddCaller())
	}
	return zap.New(core, options...).Sugar()
}
