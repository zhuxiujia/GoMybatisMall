package utils

import "time"

func DecodeTimeString(arg string) *time.Time {
	if arg == "" {
		return nil
	}
	parse_str_time, e := time.Parse("2006-01-02 15:04:05", arg)
	if e != nil {
		return nil
	}
	return &parse_str_time
}

func DecodeDateString(arg string) *time.Time {
	if arg == "" {
		return nil
	}
	parse_str_time, e := time.Parse("2006-01-02", arg)
	if e != nil {
		return nil
	}
	return &parse_str_time
}
