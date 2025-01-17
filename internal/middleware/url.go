package middleware

import "regexp"

func ValidateUrl(url string) (result bool) {
	regexCom := "^https?:\\/\\/[a-z0-9]+(?:[-.][a-z0-9]+)*(?::[0-9]{1,5})?(?:\\/[^\\/\\r\\n]+)*\\.[a-z]{2,5}(?:[?#]\\S*)?$"
	result, _ = regexp.MatchString(regexCom, url)
	return
}
