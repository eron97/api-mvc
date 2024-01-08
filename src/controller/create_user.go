package controller

import (
	"net/http"

	"github.com/eron97/api-mvc.git/src/config/error_messages"
	"github.com/eron97/api-mvc.git/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessage := error_messages.NewBadRequestError("Erro de validação")

			for _, fieldError := range ve {
				cause := error_messages.CausesMessages{
					Field:   fieldError.Field(),
					Message: fieldError.Translate(transl),
				}
				errorMessage.Causes = append(errorMessage.Causes, cause)
			}

			c.JSON(http.StatusBadRequest, gin.H{"Erro de validação": errorMessage})
			return
		}

		if err := c.ShouldBindJSON(&userRequest); err != nil {
			errorMessage := error_messages.NewUnmarshalError("Campos preenchidos não correspondem ao seus tipos atribuídos")
			c.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário criado com sucesso", "user": userRequest})
}
