package validations

import (
	"fmt"
	"net/http"
	"pos/models/base"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func BaseResponse(code int, message string, data interface{}) interface{} {
	return base.BaseResponseData{code, message, data}
}

func CustomValidation(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s wajib diisi",
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s bukan email yang valid",
						err.Field())
				case "min":
					report.Message = fmt.Sprintf("%s kurang dari minimal jumlah karakter yang diizinkan",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s nilai harus lebih besar daripada %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s nilai harus lebih kecil daripada %s",
						err.Field(), err.Param())
				default:
					report.Message = fmt.Sprint("Undefined Error",
						err.Field(), err.Param())
				}

			}
		}

		c.Logger().Error(report)
		c.JSON(report.Code, BaseResponse(
			report.Code,
			"Error Validation Field",
			report,
		))
	}
}
