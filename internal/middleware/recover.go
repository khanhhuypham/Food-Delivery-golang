package middleware

import (
	"Food-Delivery/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (m *middlewareManager) Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			fmt.Println(err)
			if err != nil {
				c.Header("Content-Type", "text/html; charset=utf-8")
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.Status, appErr)
					return
				}
				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.Status, appErr)
				return
			}
		}()
		c.Next()
	}
}
