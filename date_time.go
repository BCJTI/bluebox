package bluebox

import (
	"time"
)

const (
	DateOnlyFormat  = `"2006-01-02"`
	TimestampFormat = `"2006-01-02T15:04:05Z"`
)

type DateOnly time.Time
type Timestamp time.Time

func NewDateOnly(t time.Time) DateOnly {
	return DateOnly(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()))
}

func NewTimestamp(t time.Time) Timestamp {
	return Timestamp(time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location()))
}

func (dt DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(dt).Format(DateOnlyFormat)), nil
}

func (ts Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(ts).Format(TimestampFormat)), nil
}

func (dt *DateOnly) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(DateOnlyFormat, string(data))
	if err == nil {
		*dt = DateOnly(t)
	}
	return err
}

func (ts *Timestamp) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(TimestampFormat, string(data))
	if err == nil {
		*ts = Timestamp(t)
	}
	return err
}

func (dt DateOnly) Pointer() *DateOnly {
	return &dt
}

func (ts Timestamp) Pointer() *Timestamp {
	return &ts
}

func (dt DateOnly) Time() time.Time {
	return time.Time(dt)
}

func (ts Timestamp) Time() time.Time {
	return time.Time(ts)
}

func (dt DateOnly) String() string {
	return dt.Time().String()
}

func (ts Timestamp) String() string {
	return ts.Time().String()
}
