package logger

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// InitLogger 初始化日志系统，支持输出到文件
func InitLogger() {
	Log = logrus.New()

	// 设置日志级别
	Log.SetLevel(logrus.InfoLevel)

	// 添加默认格式化器
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 获取项目根目录（这里使用相对于当前文件的路径）
	currentDir, _ := os.Getwd()
	projectRoot := currentDir
	// 如果当前目录不是项目根目录，尝试向上查找
	if !strings.HasSuffix(projectRoot, "init_order") {
		// 直接使用绝对路径作为备选方案
		projectRoot = "/home/hjl/web3project/init_order"
	}
	
	// 获取日志目录路径
	logDir := filepath.Join(projectRoot, "log")
	
	// 创建日志目录（如果不存在）
	if err := os.MkdirAll(logDir, 0755); err != nil {
		// 如果无法创建日志目录，仅输出到标准输出
		Log.SetOutput(os.Stdout)
	} else {
		// 创建日志文件
		logFileName := time.Now().Format("2006-01-02") + ".log"
		logFilePath := filepath.Join(logDir, logFileName)
		
		// 打开日志文件（如果不存在则创建）
		logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// 如果无法打开日志文件，仅输出到标准输出
			Log.SetOutput(os.Stdout)
		} else {
			// 只输出到文件，不输出到终端
			Log.SetOutput(logFile)
		}
	}

	Log.Info("日志系统初始化完成")
}

// Debug 记录调试日志
func Debug(args ...interface{}) {
	if Log != nil {
		Log.Debug(args...)
	}
}

// Debugf 记录格式化的调试日志
func Debugf(format string, args ...interface{}) {
	if Log != nil {
		Log.Debugf(format, args...)
	}
}

// Info 记录信息日志
func Info(args ...interface{}) {
	if Log != nil {
		Log.Info(args...)
	}
}

// Infof 记录格式化的信息日志
func Infof(format string, args ...interface{}) {
	if Log != nil {
		Log.Infof(format, args...)
	}
}

// Warn 记录警告日志
func Warn(args ...interface{}) {
	if Log != nil {
		Log.Warn(args...)
	}
}

// Warnf 记录格式化的警告日志
func Warnf(format string, args ...interface{}) {
	if Log != nil {
		Log.Warnf(format, args...)
	}
}

// Error 记录错误日志
func Error(args ...interface{}) {
	if Log != nil {
		Log.Error(args...)
	}
}

// Errorf 记录格式化的错误日志
func Errorf(format string, args ...interface{}) {
	if Log != nil {
		Log.Errorf(format, args...)
	}
}

// Fatal 记录致命错误日志并退出程序
func Fatal(args ...interface{}) {
	if Log != nil {
		Log.Fatal(args...)
	} else {
		os.Exit(1)
	}
}

// Fatalf 记录格式化的致命错误日志并退出程序
func Fatalf(format string, args ...interface{}) {
	if Log != nil {
		Log.Fatalf(format, args...)
	} else {
		os.Exit(1)
	}
}
