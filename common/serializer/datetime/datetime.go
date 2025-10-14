package datetime

import (
	"database/sql/driver"
	"errors"
	"time"
)

// Datetime 日期时间
type Datetime struct {
	time.Time
}

// MarshalJSON 编码为自定义的Json格式
func (d Datetime) MarshalJSON() ([]byte, error) {

	// 时间为零返回null
	if d.IsZero() {
		return []byte("null"), nil
	}

	return []byte("\"" + d.Format(DATETIME_FORMAT0) + "\""), nil
}

// UnmarshalJSON 将Json格式解码
func (d *Datetime) UnmarshalJSON(data []byte) error {

	var err error

	if len(data) == 2 || string(data) == "null" {
		return err
	}

	var now time.Time

	// 自定义格式解析
	if now, err = time.ParseInLocation(DATETIME_FORMAT0, string(data), time.Local); err == nil {
		*d = Datetime{now}
		return err
	}

	// 带引号的自定义格式解析
	if now, err = time.ParseInLocation("\""+DATETIME_FORMAT0+"\"", string(data), time.Local); err == nil {
		*d = Datetime{now}
		return err
	}

	// 默认格式解析
	if now, err = time.ParseInLocation(time.RFC3339, string(data), time.Local); err == nil {
		*d = Datetime{now}
		return err
	}

	if now, err = time.ParseInLocation("\""+time.RFC3339+"\"", string(data), time.Local); err == nil {
		*d = Datetime{now}
		return err
	}

	return err
}

// Value 转换为数据库值
func (d Datetime) Value() (driver.Value, error) {

	if d.IsZero() {
		return nil, nil
	}

	return d.Time, nil
}

// Scan 数据库值转换为Datetime
func (d *Datetime) Scan(value interface{}) error {

	if val, ok := value.(time.Time); ok {
		*d = Datetime{Time: val}
		return nil
	}

	return errors.New("无法将值转换为时间戳")
}
