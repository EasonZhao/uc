package util

import (
	"regexp"
)

const (
	PHONE_REG = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
	EMAIL_REG = "^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+"
)

func CheckPhoneNum(num string) bool {
	reg := regexp.MustCompile(PHONE_REG)
	return reg.MatchString(num)
}

func CheckPassword(pw string) bool {
	return true
}

func CheckEmail(email string) bool {
	if m, _ := regexp.MatchString(EMAIL_REG, email); !m {
		return false
	}
	return true
}
