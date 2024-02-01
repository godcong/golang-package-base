package slog

import (
	"io"
	"log/slog"
	"strings"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func SplitLog() io.Writer {
	return &lumberjack.Logger{
		Filename:   "/path/file.log", // 日志文件的位置
		MaxSize:    10,               // 文件最大尺寸（以MB为单位）
		MaxBackups: 3,                // 保留的最大旧文件数量
		MaxAge:     28,               // 保留旧文件的最大天数
		Compress:   true,             // 是否压缩/归档旧文件
		LocalTime:  true,             // 使用本地时间创建时间戳
	}
}

func LogOutputType(w io.Writer, handlerType string, opts *slog.HandlerOptions) *slog.Logger {
	switch handlerType {
	case "text":
		return slog.New(slog.NewTextHandler(w, opts))
	case "json":
		return slog.New(slog.NewTextHandler(w, opts))
	}
	return slog.Default()
}

func SetLogLevel(opts *slog.HandlerOptions, level string) *slog.HandlerOptions {
	opts.Level = stringToLevel(level)
	return opts
}

func stringToLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}
