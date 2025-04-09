package user_rpc

//
//type userHandler struct {
//	userService UserService
//	hasher      *utils.Hasher
//}
//
//func (ctrl *UserGinHttp) IntrospectTokenRPC(c *gin.Context) {
//	var bodyData struct {
//		Token string `json:"token"`
//	}
//
//	if err := c.ShouldBindJSON(&bodyData); err != nil {
//		panic(common.ErrBadRequest(err))
//	}
//
//	user, err := ctrl.introspectCmdHdl.Execute(c.Request.Context(), &userservice.IntrospectCommand{Token: bodyData.Token})
//
//	if err != nil {
//		panic(err)
//	}
//
//	c.JSON(http.StatusOK, common.Response(user))
//}
