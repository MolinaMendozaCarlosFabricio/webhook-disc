package infrastructure

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {

	routes := router.Group("webhook")
	{
		routes.POST("/events", HandlePullRequestEvent)
		routes.POST("/actions", HandleGithubActionEvent)
	}

}
