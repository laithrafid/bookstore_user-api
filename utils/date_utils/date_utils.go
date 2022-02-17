package date_utils

import "time"

//String returns the time formatted using the format string
//	"2006-01-02 15:04:05.999999999 -0700 MST"
const (
	apiDateLayout = "2006-01-02T15:04:05Z-0700:MST"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
