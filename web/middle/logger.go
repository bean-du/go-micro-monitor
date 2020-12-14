package middle

import (
	"github.com/gin-gonic/gin"
)

func Log(ctx *gin.Context) {

	ctx.Next()

}
