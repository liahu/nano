package utils

import (
	"context"
	"testing"
)

// TestLogger 测试日志系统是否正常工作
func TestLogger(t *testing.T) {
	// 初始化日志系统
	err := InitLog()
	if err != nil {
		t.Fatalf("日志初始化失败: %v", err)
	}

	// 测试基本日志
	t.Log("测试基本日志...")
	Info(context.Background(), "hello world")
	Debug(context.Background(), "hello world")
	Error(context.Background(), "hello world")

	// 测试带traceID的日志
	t.Log("测试带traceID的日志...")
	ctx := WithTraceID(context.Background(), "test_user_123")
	Info(ctx, "带traceID的信息日志", map[string]interface{}{
		"userID": "test_user_123",
		"action": "login",
	})

	Debug(ctx, "带traceID的调试日志")
	Error(ctx, "带traceID的错误日志", map[string]interface{}{
		"error": "test error",
		"code":  500,
	})

	t.Log("日志测试完成")
}
