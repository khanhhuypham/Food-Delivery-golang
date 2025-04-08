package middleware

import (
	"Food-Delivery/pkg/common"
	"Food-Delivery/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func extractTokenFromHeader(r *http.Request) (accessToken string, err error) {
	//bearerToken: Bearer zxcxzcz...
	bearerToken := r.Header.Get("Authorization")
	splitedParts := strings.Split(bearerToken, " ")
	if splitedParts[0] != "Bearer" || len(splitedParts) < 2 || strings.TrimSpace(splitedParts[1]) == "" {
		return "", utils.ErrInvalidToken
	}
	accessToken = splitedParts[1]
	return accessToken, nil
}

// middleware kiểm tra xem token có hợp lệ hay không
// B1:Get Token from header
// B2: validate token and get payload
// B3: từ payload, dùng email để tìm user trong
func (middleware *MiddlewareManager) RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := extractTokenFromHeader(ctx.Request)
		if err != nil {
			panic(err)
		}
		tokenPayload, err := utils.ValidateJwt(accessToken, middleware.configuration)
		if err != nil {
			panic(err)
		}
		user, err := middleware.userRepo.FindDataWithCondition(ctx, map[string]any{"email": tokenPayload.Email})
		if err != nil {
			panic(err)
		}

		ctx.Set(common.KeyRequester, user)
		ctx.Next()
	}
}
