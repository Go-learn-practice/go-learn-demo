package cache

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	//默认大小为100MB
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)

	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = KB * num
	case "MB":
		byteNum = MB * num
	case "GB":
		byteNum = GB * num
	case "TB":
		byteNum = TB * num
	case "PB":
		byteNum = PB * num
	default:
		num = 0
		byteNum = 0
	}

	if num == 0 {
		log.Println("ParseSize 仅支持 B KB MB GB TB PB 单位")
		num = 100
		byteNum = MB * num
		unit = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unit

	return byteNum, sizeStr
}

func GetValueSize(val interface{}) int64 {
	return 0
}
