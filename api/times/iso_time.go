package times

import (
	"fmt"
	"strings"
	"time"
)

// ISOTime 标准库Time的别名类型。用于格式化为ISO8601标准的时间（形如'2015-09-09 09:09:09'）。
type ISOTime time.Time

const (
	// DateTimeLayout 日期时间转换模式
	DateTimeLayout = "2006-01-02 15:04:05"
	// DateLayout 日期转换模式
	DateLayout = "2006-01-02"
)

// MarshalJSON 序列化，实现json.Marshaller接口
func (t ISOTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte(fmt.Sprintf("%q", "")), nil
	}

	// 注意：序列化成的时间字符串必须包含在双引号当中
	return []byte(fmt.Sprintf("%q", time.Time(t).In(time.Local).Format(DateTimeLayout))), nil
}

// UnmarshalJSON 反序列化,实现json.Unmarshaler接口
func (t *ISOTime) UnmarshalJSON(b []byte) error {
	sTime := string(b)

	sTime = strings.TrimPrefix(sTime, "\"")
	sTime = strings.TrimSuffix(sTime, "\"")

	tmpT, err := time.Parse(DateTimeLayout, sTime)
	if err != nil {
		return err
	}

	*t = ISOTime(tmpT)
	return nil
}

// MarshalYAML 序列化成YAML
func (t ISOTime) MarshalYAML() (interface{}, error) {
	return time.Time(t).In(time.Local).Format(DateTimeLayout), nil
}

func (t ISOTime) String() string {
	return time.Time(t).In(time.Local).Format(DateTimeLayout)
}
