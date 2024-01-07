package controller

import (
	"fmt"
	"net/http"

	"github.com/eron97/api-mvc.git/src/config/error_messages"
	"github.com/eron97/api-mvc.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errorMessage := error_messages.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()),
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	c.JSON(http.StatusOK, userRequest)
}
