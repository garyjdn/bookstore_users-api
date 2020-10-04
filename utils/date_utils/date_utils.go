package date_utils

import "time"

const (
	dateFormatLayout = "2006-01-02T15:04:05Z"
	apiDBLayout      = "2006-01-02 15:04:05"
)

func Now() time.Time {
	return time.Now().UTC()
}

func NowString() string {
	return Now().Format(dateFormatLayout)
}

func GetNowDBFormat() string {
	return Now().Format(apiDBLayout)
}
