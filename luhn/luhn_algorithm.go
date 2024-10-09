package luhn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type WrongCardFormatError struct{}
type InternalError struct{}

func (w *WrongCardFormatError) Error() string {
	return "Your card have wrong format"
}

func (i *InternalError) Error() string {
	return "Something went wrong on our side"
}

func Validate(cardValue string) (bool, error) {
	cardValue = strings.ReplaceAll(cardValue, " ", "")
	if matched, _ := regexp.MatchString("^[\\d]{16}$", cardValue); !matched {
		return false, &WrongCardFormatError{}
	}
	sArr := strings.Split(cardValue, "")
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
	fmt.Println(sum)
	return sum%10 == 0, nil
}
