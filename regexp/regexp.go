package _regexp

import (
	"regexp"
	"strings"

	gopinyin "github.com/mozillazg/go-pinyin"
)

func ReplaceMultipleSpaces(s string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(s, " ")
}

func RemoveCountryCode(phoneNumber string) string {
	re := regexp.MustCompile(`^\+86`)
	return re.ReplaceAllString(phoneNumber, "")
}

// StringToPinyin 字符串（包含汉字）转拼音
func StringToPinyin(s string) string {
	re := regexp.MustCompile("[\u4e00-\u9fa5]")
	return re.ReplaceAllStringFunc(s, chinese2Pinyin)
}

func chinese2Pinyin(s string) string {
	pinyinList := gopinyin.LazyPinyin(s, gopinyin.NewArgs())
	return strings.Join(pinyinList, "")
}
