package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotlog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

var (
	filePath string
	linkName string
)

func Logger() gin.HandlerFunc {
	// 指定环境变量，日志代码的路径
	logRootPath := os.Getenv("GO_BLOG_LOG_ROOT")
	if logRootPath != "" {
		// 文件路径
		filePath = logRootPath + "/gin-blog.log"
		linkName = logRootPath + "/latest_log.log"
	} else {
		// 文件路径
		filePath = "log/gin-blog.log"
		linkName = "log/latest_log.log"
	}

	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err: ", err)
	}

	logger := logrus.New()

	// 输入日志
	logger.Out = src

	// 日志级别 分割
	logWriter, _ := rotlog.New(
		filePath+"%Y%m%d.log",
		rotlog.WithMaxAge(7*24*time.Hour),
		rotlog.WithRotationTime(24*time.Hour),
		rotlog.WithLinkName(linkName), // 软连接 需要管理员权限
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
		logrus.TraceLevel: logWriter,
	}

	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // 需要深入了解中间件的流程 洋葱模型 从左到右 中心为Next
		durationTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(durationTime.Nanoseconds()/1000000.0))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown host name"
		}

		statusCode := c.Writer.Status()    // 状态码
		clientIp := c.ClientIP()           // 客户端的IP
		userAgent := c.Request.UserAgent() // 客户端的信息
		dataSize := c.Writer.Size()        // 文件的大小
		method := c.Request.Method         // 请求的方法
		path := c.Request.RequestURI       // 请求的路径

		if dataSize < 0 {
			dataSize = 0 // 格式化
		}

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"IP":        clientIp,
			"Method":    method,
			"path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		// gin 框架 内部的错误
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		// 状态码
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
