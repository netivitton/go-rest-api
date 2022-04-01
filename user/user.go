package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func String() string {
	return "ta da! \n"
}

func GetTags(c *gin.Context) {
	fmt.Println("GetTags")
}
