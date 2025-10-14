package datetime

import (
	"database/sql/driver"
	"errors"
	"time"
)

// 年月日格式
const (
	RFC_FORMAT       = "2006-01-02T15:04:05Z07:00"
	YEAR_FORMAT      = "2006"
	YEAR_FORMAT0     = "2006年"
	MONTH_FORMAT0    = "200601"
	MONTH_FORMAT1    = "2006-01"
	MONTH_FORMAT2    = "2006年01月"
	DATE_FORMAT0     = "2006-01-02"
	DATE_FORMAT1     = "20060102"
	DATE_FORMAT2     = "01022006"
	DATE_FORMAT3     = "02012006"
	DATE_FORMAT4     = "2006年01月02日"
	DATETIME_FORMAT0 = "2006-01-02 15:04:05"
	DATETIME_FORMAT1 = "2006日01月02日 15时04分05秒"
	DATETIME_FORMAT2 = "20060102150405"
	TIME_FORMAT0     = "15:04:05"
	TIME_FORMAT1     = "15时04分05秒"
)

// Date 日期
type Date struct {
	time.Time
}

// MarshalJSON 编码为自定义的Json格式
func (d Date) MarshalJSON() ([]byte, error) {

	// 时间为零返回null
	if d.IsZero() {
		return []byte("null"), nil
	}

	return []byte("\"" + d.Format(DATE_FORMAT0) + "\""), nil
}

// UnmarshalJSON 将Json格式解码
func (d *Date) UnmarshalJSON(data []byte) error {

	var err error

	if len(data) == 2 || string(data) == "null" {
		return err
	}

	var now time.Time

	// 自定义格式解析
	if now, err = time.ParseInLocation(DATE_FORMAT0, string(data), time.Local); err == nil {
		*d = Date{now}
		return err
	}

	// 带引号的自定义格式解析
	if now, err = time.ParseInLocation("\""+DATE_FORMAT0+"\"", string(data), time.Local); err == nil {
		*d = Date{now}
		return err
	}

	return err
}

// Value 转换为数据库值
func (d Date) Value() (driver.Value, error) {

	if d.IsZero() {
		return nil, nil
	}

	return d.Time, nil
}

// Scan 数据库值转换为Date
func (d *Date) Scan(value interface{}) error {

	if val, ok := value.(time.Time); ok {
		*d = Date{Time: val}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}
