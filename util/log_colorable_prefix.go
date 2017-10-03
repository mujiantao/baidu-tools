package baiduUtil

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

var (
	// ErrorColor 设置输出错误的颜色
	ErrorColor = color.New(color.FgRed).SprintFunc()
)

// 自定义log writer
type logWriter struct{}

func (logWriter) Write(bytes []byte) (int, error) {
	return fmt.Fprintf(color.Output, "["+BeijingTimeOption("Refer")+"] "+string(bytes))
}

// SetLogPrefix 设置日志输出的时间前缀
func SetLogPrefix() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}