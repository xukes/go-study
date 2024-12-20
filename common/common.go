package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
)

func GetLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
}

func ShowHandleTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)
}

func HandlerError(c *gin.Context) {
	c.Next()
	// 检查是否有发生错误
	err := c.Errors.Last()
	if err != nil {
		// 输出错误信息
		// 返回统一的错误响应
		c.JSON(http.StatusOK, err.Meta)
	}
}
