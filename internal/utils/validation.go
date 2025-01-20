package utils

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"net"
	"net/url"
	"regexp"
	"strings"
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
  validate.RegisterValidation("isodate", validateIsoDate)
}

func Validate[T any](data T) map[string]string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	res := map[string]string{}

	fmt.Println(err)
	for _, v := range err.(validator.ValidationErrors) {
		res[v.StructField()] = v.Translate(trans)
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

func validateIsoDate(tl validator.FieldLevel) bool {
	ISO8601DateRegexString := "^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])(?:T|\\s)(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])?(Z)?$"
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)

	return ISO8601DateRegex.MatchString(tl.Field().String())
}

func validateRFC3339(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}
