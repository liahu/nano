package utils

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"nano/src/utils/mlog"
)

var (
	// Log 全局日志对象
	Log *mlog.Logger
)

// InitLog 初始化日志系统
func InitLog() error {
	// 配置日志
	config := &mlog.LogCfg{
		LogLevel: "info",
		LogDir:   "./logs",
		LogFile:  "server.log",
	}

	// 使用mlog的NewLogger
	logger, err := mlog.NewLogger(config)
	if err != nil {
		return err
	}

	Log = logger
	return nil
}

// WithTraceID 添加上下文traceID
func WithTraceID(ctx context.Context, userID string) context.Context {
	// 生成与userID相关的traceID
	traceID := generateTraceID(userID)
	return context.WithValue(ctx, "traceID", traceID)
}

// getTraceIDArgs 从上下文中获取traceID并添加到参数中
func getTraceIDArgs(ctx context.Context, args ...any) []any {
	if ctx != nil {
		if traceID, ok := ctx.Value("traceID").(string); ok {
			// 将traceID添加到参数中
			if len(args) > 0 {
				if argsMap, ok := args[0].(map[string]interface{}); ok {
					argsMap["traceID"] = traceID
				} else {
					// 如果不是map，创建一个新的map
					newArgs := map[string]interface{}{
						"traceID": traceID,
						"args":    args,
					}
					args = []any{newArgs}
				}
			} else {
				args = []any{map[string]interface{}{"traceID": traceID}}
			}
		}
	}
	return args
}

// Info 信息级别日志
func Info(ctx context.Context, msg string, args ...any) {
	args = getTraceIDArgs(ctx, args...)
	if Log != nil {
		Log.Info(msg, args...)
	}
}

// Debug 调试级别日志
func Debug(ctx context.Context, msg string, args ...any) {
	args = getTraceIDArgs(ctx, args...)
	if Log != nil {
		Log.Debug(msg, args...)
	}
}

// Error 错误级别日志
func Error(ctx context.Context, msg string, args ...any) {
	args = getTraceIDArgs(ctx, args...)
	if Log != nil {
		Log.Error(msg, args...)
	}
}

// generateTraceID 生成与userID相关的traceID
func generateTraceID(userID string) string {
	// 使用时间戳 + userID的哈希生成traceID
	timestamp := time.Now().UnixNano()
	data := userID + "_" + strconv.FormatInt(timestamp, 10)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
