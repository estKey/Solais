package utils

import (
	"runtime"
	"time"
)

var start = time.Now()

// calculateUptime 计算运行时间
func CalculateUptime() interface{} {
	// return time.Since(start).Nanoseconds() 
	return time.Since(start).Nanoseconds()
}

// getNumGoroutins 当前 goroutine 数量
func GetNumGoroutins() interface{} {
	return runtime.NumGoroutine()
}

// currentGoVersion 当前 Golang 版本
func CurrentGoVersion() interface{} {
  return runtime.Version()
}

// getNumCPUs 获取 CPU 核心数量
func GetNumCPUs() interface{} {
  return runtime.NumCPU()
}

// getGoOS 当前系统类型
func GetGoOS() interface{} {
  return runtime.GOOS
}
