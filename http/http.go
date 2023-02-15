package http

import "github.com/gin-gonic/gin"

// TODO - ADD HTTP ROUTES INTO THIS PACKAGE

func New() *gin.Engine {
	return gin.Default()
}

func Listen(ge *gin.Engine) {
	ge.Run() // listen and serve by default on 0.0.0.0:8080, but port will be override if set in .env variables
}
