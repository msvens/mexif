package json

import (
	"errors"
	"log"
	"time"
)

type JsonType uint

type JSONObject map[string]interface{}
type JSONArray []interface{}

const (
	JString JsonType = iota + 100
	JNumber
	JBool
	JArr
	JObj
	JNull
	JUnknown
)

const (
	ExifDateTime       = "2006:01:02 15:04:05"
	ExifDateTimeOffset = "2006:01:02 15:04:05 -07:00"
	ExifDate           = "2006:01:02"
	ExifTime           = "15:04:05"
)

var IncorrectType = errors.New("Incorrect Type")
var ValueNotFound = errors.New("Value Not Found")

func GetUInt(field string, obj JSONObject) (uint, error) {
	n, err := GetNumber(field, obj)
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, IncorrectType
	}
	return uint(n), nil
}

func ScanUInt(field string, obj JSONObject, val *uint) error {
	if v, err := GetUInt(field, obj); err == nil {
		*val = v
		return nil
	} else {
		return err
	}
}

func GetInt(field string, obj JSONObject) (int, error) {
	n, err := GetNumber(field, obj)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func GetFloat32(field string, obj JSONObject) (float32, error) {
	n, err := GetNumber(field, obj)
	if err != nil {
		return 0, err
	}
	return float32(n), nil
}

func ScanFloat32(field string, obj JSONObject, val *float32) error {
	if v, err := GetFloat32(field, obj); err == nil {
		*val = v
		return nil
	} else {
		return err
	}
}

func GetFloat64(field string, obj JSONObject) (float64, error) {
	n, err := GetNumber(field, obj)
	if err != nil {
		return 0, err
	}
	return float64(n), nil
}

func ScanFloat64(field string, obj JSONObject, val *float64) error {
	if v, err := GetFloat64(field, obj); err == nil {
		*val = v
		return nil
	} else {
		return err
	}
}

func GetDateTime(dtField string, offsetField string, obj JSONObject) (time.Time, error) {
	dt, err := GetString(dtField, obj)
	if err != nil {
		return time.Time{}, err
	}
	offset, _ := GetString(offsetField, obj)
	return ParseDateTime(dt, offset)
}

func ScanDateTime(dtField string, offsetField string, obj JSONObject, val *time.Time) error {
	if v, err := GetDateTime(dtField, offsetField, obj); err == nil {
		*val = v
		return nil
	} else {
		return err
	}

}

func ParseDateTime(dt string, offset string) (time.Time, error) {
	if offset == "" {
		return time.Parse(ExifDateTime, dt)
	} else {
		return time.Parse(ExifDateTimeOffset, dt+" "+offset)
	}
}

func GetBool(field string, obj JSONObject) (bool, error) {
	if v, f := obj[field]; f {
		if s, t := v.(bool); t {
			return s, nil
		} else {
			return false, IncorrectType
		}
	}
	return false, ValueNotFound
}

func ScanBool(field string, obj JSONObject, val *bool) error {
	if v, err := GetBool(field, obj); err == nil {
		*val = v
		return nil
	} else {
		return err
	}
}

func GetArray(field string, obj JSONObject) (JSONArray, error) {
	if v, f := obj[field]; f {
		if s, t := v.([]interface{}); t {
			return s, nil
		} else {
			return JSONArray{}, IncorrectType
		}
	}
	return JSONArray{}, ValueNotFound
}

func ScanArray(field string, obj JSONObject, val *JSONArray) error {
	if v, err := GetArray(field, obj); err == nil {
		*val = v
		return nil
	} else {
		return err
	}
}

func GetObject(field string, obj JSONObject) (JSONObject, error) {
	if v, f := obj[field]; f {
		if s, t := v.(map[string]interface{}); t {
			return s, nil
		} else {
			return JSONObject{}, IncorrectType
		}
	}
	return JSONObject{}, ValueNotFound
}

func ScanObject(field string, obj JSONObject, val *JSONObject) error {
	if v, err := GetObject(field, obj); err == nil {
		*val = v
		return nil
	} else {
		return err
	}
}

func GetNumber(field string, obj JSONObject) (float64, error) {
	if v, f := obj[field]; f {
		if s, t := v.(float64); t {
			return s, nil
		} else {
			return 0, IncorrectType
		}
	}
	return 0, ValueNotFound
}

func ScanNumber(field string, obj JSONObject, val *float64) error {
	if v, err := GetNumber(field, obj); err == nil {
		*val = v
		return nil
	} else {
		return err
	}
}

func GetString(field string, obj JSONObject) (string, error) {
	if v, f := obj[field]; f {
		if s, t := v.(string); t {
			return s, nil
		} else {
			return "", IncorrectType
		}
	}
	return "", ValueNotFound
}

func ScanString(field string, obj JSONObject, val *string) error {
	if str, err := GetString(field, obj); err == nil {
		*val = str
		return nil
	} else {
		return err
	}
}

func IsType(val interface{}, jsonType JsonType) bool {
	return TypeOf(val) == jsonType
}

func TypeOf(i interface{}) JsonType {

	if i == nil {
		return JNull
	}
	switch t := i.(type) {
	case string:
		return JString
	case float64:
		return JNumber
	case bool:
		return JBool
	case []interface{}:
		return JArr
	case map[string]interface{}:
		return JObj
	default:
		log.Println(t)
		return JUnknown
	}
}
