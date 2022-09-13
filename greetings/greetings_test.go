// 以 _test.go 结尾的文件名告诉go test命令该文件包含测试函数
package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName调用Hello函数，传递一个name值，函数应该能够使用该值返回有效的响应消息。如果调用返回错误或意外响应消息（不包括您传入的名称），则使用t参数的 方法将消息打印到控制台并结束执行。
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty Hello使用空字符串调用函数。此测试旨在确认您的错误处理是否有效。如果调用返回非空字符串或没有错误，则使用t参数的Fatalf 方法将消息打印到控制台并结束执行。
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}