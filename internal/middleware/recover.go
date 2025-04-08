package middleware

import (
	"Food-Delivery/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func (m *MiddlewareManager) Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {

			if err := recover(); err != nil {
				c.Header("Content-Type", "text/html; charset=utf-8")
				if appErr, ok := err.(*common.AppError); ok {

					c.AbortWithStatusJSON(appErr.Status, appErr)
					return
				}

				appErr := common.ErrInternal(err.(error))
				env := os.Getenv("ENV")

				if strings.ToLower(env) == "production" {
					c.AbortWithStatusJSON(appErr.Status, appErr)
				} else {
					c.AbortWithStatusJSON(appErr.Status, appErr.WithDebug(fmt.Sprintf("%s", err)))
					panic(err)
				}

				return
			}
		}()
		c.Next()
	}
}
