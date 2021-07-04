package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

// 打印结构体为字符串
func PrintObjectJson(obj interface{}) string {
	objByte, err := json.Marshal(obj)
	if err != nil {
		return fmt.Sprintf("json marshal failed, err: %v", err)
	}
	return fmt.Sprintf("%s", objByte)
}

// 生成指定位数的随机字符串
func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// 转换map为string
func ConvertMapToStr(toConvertMap map[string]string) string {
	str := make([]string, len(toConvertMap))
	i := 0
	for name, value := range toConvertMap {
		str[i] = fmt.Sprintf("%s=%s", name, value)
		i++
	}
	sort.Strings(str)
	return strings.Join(str, ",")
}

// 判断元素是否在列表中
func IsInList(str string, list []string) bool {
	listMap := make(map[string]bool)
	for _, key := range list {
		listMap[key] = true
	}
	if _, ok := listMap[str]; ok {
		return true
	}
	return false
}
