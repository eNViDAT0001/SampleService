package external

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type DateTime struct {
	OriginValue string
	DateTime    *time.Time
}

const (
	LayoutDateTimeISO = "2006/01/02 15:04:05"
)

func (d *DateTime) UnmarshalJSON(b []byte) (err error) {
	if d == nil {
		return nil
	}

	s := strings.Trim(string(b), "\"")
	d.OriginValue = s

	if s == "null" || len(s) == 0 {
		return nil
	}

	t, err := time.Parse(LayoutDateTimeISO, s)
	if err != nil {
		// không trả về lỗi khi bind
		return nil
	}
	d.DateTime = &t

	return
}

// implement Stringer interface
func (d *DateTime) String() string {
	if d == nil {
		return ""
	}
	return d.OriginValue
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	if d == nil || d.DateTime == nil {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", d.DateTime.Format(LayoutDateTimeISO))), nil
}

// Scan scan value into Date, implements sql.Scanner interface
func (d *DateTime) Scan(value interface{}) error {
	if d == nil {
		return nil
	}

	t, ok := value.(time.Time)
	if ok {
		d.DateTime = &t
	}

	return nil
}

// Value return Date value, implement driver.Valuer interface
func (d DateTime) Value() (driver.Value, error) {
	if d.DateTime == nil {
		return nil, nil
	}
	return *d.DateTime, nil
}

func NowDateTime() DateTime {
	now := time.Now()
	return DateTime{DateTime: &now}
}
