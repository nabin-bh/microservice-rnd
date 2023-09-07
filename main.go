package microservicernd

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/send-email", func(c *gin.Context) {
		//
	})

	r.Run(":8080")
}
