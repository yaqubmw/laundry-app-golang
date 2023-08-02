package middleware

import (
	"enigma-laundry-apps/config"
	"enigma-laundry-apps/model"
	"enigma-laundry-apps/utils/exceptions"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogRequestMiddleware(log *logrus.Logger) gin.HandlerFunc {

	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	exceptions.CheckErr(err)
	log.SetOutput(file)

	return func(c *gin.Context) {

		startTime := time.Now()

		c.Next()

		endTime := time.Since(startTime)

		requestLog := model.RequestLog{
			StartTime:  startTime,
			EndTime:    endTime,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			UserAgent:  c.Request.UserAgent(),
		}

		switch {
		case c.Writer.Status() >= 500:
			log.Error(requestLog)
		case c.Writer.Status() >= 400:
			log.Warn(requestLog)
		default:
			log.Info(requestLog)
		}

	}
}
