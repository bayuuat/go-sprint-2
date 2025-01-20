package utils

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	enLocale := en.New()
	uni = ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	enTranslations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterValidation("accessibleuri", validateAccessibleURI)
	validate.RegisterValidation("rfc3339", validateRFC3339)
}

func Validate[T any](data T) map[string]string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	res := map[string]string{}
	if err != nil {
		fmt.Print(err)
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = v.Translate(trans)
		}
	}
	return res
}

func validateAccessibleURI(fl validator.FieldLevel) bool {
	uri := fl.Field().String()

	parsedURL, err := url.Parse(uri)
	if err != nil {
		return false
	}

	if !strings.HasPrefix(parsedURL.Scheme, "http") {
		return false
	}

	if parsedURL.Host == "" {
		return false
	}

	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

	host := parsedURL.Host
	if strings.Contains(host, ":") {
		host, _, err = net.SplitHostPort(parsedURL.Host)
		if err != nil {
			return false
		}
	}

	return domainRegex.MatchString(host)

}

func validateRFC3339(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}
