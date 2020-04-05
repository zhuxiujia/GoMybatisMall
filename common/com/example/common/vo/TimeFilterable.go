package vo

import (
	"time"
)

type TimeRangeable struct {
	TimeStartStr string `json:"time_start"`
	TimeEndStr   string `json:"time_end"`
}

func (it TimeRangeable) New(time_start string, time_end string) TimeRangeable {
	it.TimeStartStr = time_start
	it.TimeEndStr = time_end
	return it
}

func (it *TimeRangeable) TimeStart() *time.Time {
	return decodeTimeString(it.TimeStartStr)
}
func (it *TimeRangeable) TimeEnd() *time.Time {
	return decodeTimeString(it.TimeEndStr)
}

func decodeTimeString(arg string) *time.Time {
	if arg == "" {
		return nil
	}
	parse_str_time, e := time.Parse("2006-01-02 15:04:05", arg)
	if e != nil {
		return nil
	}
	return &parse_str_time
}
