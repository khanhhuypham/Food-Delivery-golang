package middleware

import (
	"Food-Delivery/entity/model"
	"Food-Delivery/pkg/common"
	"github.com/gin-gonic/gin"
)

func (m *MiddlewareManager) RequiredRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*model.User)

		for i := range roles {
			if user.GetUserRole() == roles[i] {
				c.Next()
				return
			}
		}
		panic(common.ErrForbidden(nil))

	}
}
