package _time

import (
	"errors"
	"strconv"
	"time"
)

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func ParseMillisecondTimestamp(ts string) (time.Time, error) {
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, errors.New("无效的毫秒时间戳格式: " + err.Error())
	}
	return time.Unix(0, timestamp*int64(time.Millisecond)), nil
}

func ParseSecondTimestamp(ts string) (time.Time, error) {
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return time.Time{}, errors.New("无效的秒时间戳格式: " + err.Error())
	}
	return time.Unix(timestamp, 0), nil
}
