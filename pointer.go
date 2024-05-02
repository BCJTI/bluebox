package bluebox

import "time"

func Bool2Pnt(value bool) *bool {
	return &value
}

func Int2Pnt(value int) *int {
	return &value
}

func Float2Pnt(value float64) *float64 {
	return &value
}

func Str2Pnt(value string) *string {
	return &value
}

func Time2Pnt(value time.Time) *time.Time {
	return &value
}
