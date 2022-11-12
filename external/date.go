package external

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Date struct {
	OriginValue string
	DateTime    *time.Time
}

const (
	LayoutISO = "2006/01/02"
	MinTime   = "1900/01/01"
	MaxTime   = "2999/12/31"
)

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	if d == nil {
		return nil
	}

	s := strings.Trim(string(b), "\"")
	d.OriginValue = s

	if s == "null" || len(s) == 0 {
		return nil
	}

	t, err := time.Parse(LayoutISO, s)
	if err != nil {
		// không trả về lỗi khi bind
		return nil
	}
	d.DateTime = &t

	return
}

// implement Stringer interface
func (d *Date) String() string {
	if d == nil {
		return ""
	}
	return d.OriginValue
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d == nil || d.DateTime == nil {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", d.DateTime.Format(LayoutISO))), nil
}

// Scan scan value into Date, implements sql.Scanner interface
func (d *Date) Scan(value interface{}) error {
	if d == nil {
		return nil
	}

	rValue := reflect.ValueOf(value)

	if rValue.Kind() == reflect.Ptr {
		t, ok := value.(*Date)
		if ok {
			d.DateTime = t.DateTime
		} else {
			t, ok := value.(*time.Time)
			if ok {
				d.DateTime = t
			}
		}
	} else {
		t, ok := value.(Date)
		if ok {
			d.DateTime = t.DateTime
		} else {
			t, ok := value.(time.Time)
			if ok {
				d.DateTime = &t
			}
		}
	}

	return nil
}

// Value return Date value, implement driver.Valuer interface
func (d Date) Value() (driver.Value, error) {
	if d.DateTime == nil {
		return nil, nil
	}
	return *d.DateTime, nil
}

func Now() Date {
	now := time.Now()
	return Date{DateTime: &now}
}
