package validation

import (
	"encoding/json"
	"errors"

	"github.com/felipeversiane/go-starter/internal/infra/config/response"
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

func ValidateError(
	validation_err error,
) *response.ErrorResponse {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return response.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []response.Cause{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := response.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return response.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return response.NewBadRequestError("Error trying to convert fields")
	}
}
