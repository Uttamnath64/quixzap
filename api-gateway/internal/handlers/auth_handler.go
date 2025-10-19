package handlers

import "github.com/Uttamnath64/quixzap/app/appcontext"

type Auth struct {
	appCtx *appcontext.AppContext
	// authService *services.Auth
}

func NewAuth(appCtx *appcontext.AppContext) *Auth {
	return &Auth{
		appCtx: appCtx,
		// authService: services.NewAuth(container),
	}
}

// func (handler *Auth) Login(c *gin.Context) {

// 	rctx, ok := getRequestContext(c)
// 	if !ok {
// 		return
// 	}

// 	var payload requests.LoginRequest
// 	if !bindAndValidateJson(c, &payload) {
// 		return
// 	}

// 	serviceResponse := handler.authService.Login(rctx, payload, c.Request.UserAgent(), c.ClientIP())
// 	if isErrorResponse(c, serviceResponse) {
// 		return
// 	}

// 	c.JSON(http.StatusOK, responses.ApiResponse{
// 		Status:   true,
// 		Message:  serviceResponse.Message,
// 		Metadata: serviceResponse.Data,
// 	})
// }

// func (handler *Auth) Token(c *gin.Context) {

// 	rctx, ok := getRequestContext(c)
// 	if !ok {
// 		return
// 	}

// 	var payload requests.TokenRequest
// 	if !bindAndValidateJson(c, &payload) {
// 		return
// 	}

// 	// Get token
// 	serviceResponse := handler.authService.GetToken(rctx, payload, c.Request.UserAgent(), c.ClientIP())
// 	if isErrorResponse(c, serviceResponse) {
// 		return
// 	}

// 	c.JSON(http.StatusOK, responses.ApiResponse{
// 		Status:   true,
// 		Message:  serviceResponse.Message,
// 		Metadata: serviceResponse.Data,
// 	})

// }
