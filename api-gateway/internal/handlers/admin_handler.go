package handlers

import "github.com/Uttamnath64/quixzap/app/appcontext"

type Admin struct {
	appCtx *appcontext.AppContext
	// service   *services.Admin
}

func NewAdmin(appCtx *appcontext.AppContext) *Admin {
	return &Admin{
		appCtx: appCtx,
		// service:   services.NewAdmin(container),
	}
}

// func (handler *Admin) Get(c *gin.Context) {

// 	rctx, ok := getRequestContext(c)
// 	if !ok {
// 		return
// 	}

// 	var payload requests.LoginRequest
// 	if !bindAndValidateJson(c, &payload) {
// 		return
// 	}

// 	serviceResponse := handler.service.Login(rctx, payload, c.Request.UserAgent(), c.ClientIP())
// 	if isErrorResponse(c, serviceResponse) {
// 		return
// 	}

// 	c.JSON(http.StatusOK, responses.ApiResponse{
// 		Status:   true,
// 		Message:  serviceResponse.Message,
// 		Metadata: serviceResponse.Data,
// 	})
// }

// func (handler *Admin) GetAll(c *gin.Context) {

// 	rctx, ok := getRequestContext(c)
// 	if !ok {
// 		return
// 	}

// 	var payload requests.LoginRequest
// 	if !bindAndValidateJson(c, &payload) {
// 		return
// 	}

// 	serviceResponse := handler.service.Login(rctx, payload, c.Request.UserAgent(), c.ClientIP())
// 	if isErrorResponse(c, serviceResponse) {
// 		return
// 	}

// 	c.JSON(http.StatusOK, responses.ApiResponse{
// 		Status:   true,
// 		Message:  serviceResponse.Message,
// 		Metadata: serviceResponse.Data,
// 	})
// }

// func (handler *Admin) Craete(c *gin.Context) {

// 	rctx, ok := getRequestContext(c)
// 	if !ok {
// 		return
// 	}

// 	var payload requests.LoginRequest
// 	if !bindAndValidateJson(c, &payload) {
// 		return
// 	}

// 	serviceResponse := handler.service.Login(rctx, payload, c.Request.UserAgent(), c.ClientIP())
// 	if isErrorResponse(c, serviceResponse) {
// 		return
// 	}

// 	c.JSON(http.StatusOK, responses.ApiResponse{
// 		Status:   true,
// 		Message:  serviceResponse.Message,
// 		Metadata: serviceResponse.Data,
// 	})
// }

// func (handler *Admin) Update(c *gin.Context) {

// 	rctx, ok := getRequestContext(c)
// 	if !ok {
// 		return
// 	}

// 	var payload requests.LoginRequest
// 	if !bindAndValidateJson(c, &payload) {
// 		return
// 	}

// 	serviceResponse := handler.service.Login(rctx, payload, c.Request.UserAgent(), c.ClientIP())
// 	if isErrorResponse(c, serviceResponse) {
// 		return
// 	}

// 	c.JSON(http.StatusOK, responses.ApiResponse{
// 		Status:   true,
// 		Message:  serviceResponse.Message,
// 		Metadata: serviceResponse.Data,
// 	})
// }

// func (handler *Admin) Block(c *gin.Context) {

// 	rctx, ok := getRequestContext(c)
// 	if !ok {
// 		return
// 	}

// 	var payload requests.LoginRequest
// 	if !bindAndValidateJson(c, &payload) {
// 		return
// 	}

// 	serviceResponse := handler.service.Login(rctx, payload, c.Request.UserAgent(), c.ClientIP())
// 	if isErrorResponse(c, serviceResponse) {
// 		return
// 	}

// 	c.JSON(http.StatusOK, responses.ApiResponse{
// 		Status:   true,
// 		Message:  serviceResponse.Message,
// 		Metadata: serviceResponse.Data,
// 	})
// }
