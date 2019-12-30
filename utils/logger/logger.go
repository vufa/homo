//
// Copyright (c) 2019-present Codist <countstarlight@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// Written by Codist <countstarlight@gmail.com>, August 2019
//

package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var L *zap.Logger
var S *zap.SugaredLogger

func init() {
	// Print log when start
	L = NewLogger()
	S = L.Sugar()
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

// StartLogService set global logger output to file
func StartLogService(filePath string) {
	L = NewLoggerToFile(filePath)
	S = L.Sugar()
}

// NewLogService creates logger
func NewLogger(fields ...string) *zap.Logger {
	var core zapcore.Core
	core = zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewEncoderConfig()),
		os.Stdout,
		zap.DebugLevel,
	)
	if len(fields) > 0 {
		zapFields := []zap.Field{}
		for index := 0; index < len(fields)-1; index = index + 2 {
			zapFields = append(zapFields, zap.String(fields[index], fields[index+1]))
		}
		return zap.New(core, zap.AddCaller(), zap.Fields(zapFields...))
	}
	return zap.New(core, zap.AddCaller())
}

// NewLogServiceToFile creates logger output to file
func NewLoggerToFile(filePath string, fields ...string) *zap.Logger {
	var core zapcore.Core
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	conf := zap.NewProductionEncoderConfig()
	conf.EncodeTime = TimeEncoder
	core = zapcore.NewCore(
		zapcore.NewJSONEncoder(conf),
		w,
		zap.WarnLevel,
	)
	if len(fields) > 0 && len(fields)%2 == 0 {
		zapFields := []zap.Field{}
		for index := 0; index < len(fields)-1; index = index + 2 {
			zapFields = append(zapFields, zap.String(fields[index], fields[index+1]))
		}
		return zap.New(core, zap.Fields(zapFields...))
	}
	return zap.New(core)
}
