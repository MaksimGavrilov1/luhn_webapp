package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

type wrongCardFormatError struct{}
type internalError struct{}

const (
	emptyString = ""
	digitRegex  = "^[\\d]{16}$"
	whiteSpace  = " "
)

func (w *wrongCardFormatError) Error() string {
	return "Your card have wrong format"
}

func (i *internalError) Error() string {
	return "Something went wrong on our side"
}

func Validate(cardValue string) (bool, error) {
	cardValue = strings.ReplaceAll(cardValue, whiteSpace, emptyString)
	if matched, _ := regexp.MatchString(digitRegex, cardValue); !matched {
		return false, &wrongCardFormatError{}
	}
	sArr := strings.Split(cardValue, emptyString)
	var iElem, sum int
	var err error
	for index, elem := range sArr {
		iElem, err = strconv.Atoi(elem)
		if err != nil {
			return false, err
		}
		if index%2 != 0 {
			iElem = iElem * 2
			if iElem > 9 {
				sum += iElem - 9
			} else {
				sum += iElem
			}
		} else {
			sum += iElem
		}
	}
	return sum%10 == 0, nil
}
