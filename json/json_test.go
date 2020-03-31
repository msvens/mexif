package json

import (
	"fmt"
	"math"
	"testing"
	"time"
)

type NoType struct {
	v int
}

var testuint uint = 1
var testfloat float32 = 2.2
var testnum float64 = math.MaxFloat64
var teststr string = "teststring"
var testdatetime string = "2020:01:01 15:01:01"
var testdtwrong string = "2020/01/01 15:01:01"
var testdatetimeoffset string = "2020:01:01 15:01:01 +01:00"
var testoffset string = "+01:00"
var testbool bool = true
var testnotype NoType = NoType{}
var testarr []interface{} = []interface{}{teststr, teststr}
var testobj = map[string]interface{}{
	"val1": "val1",
	"val2": "val2",
}

var rootObj = map[string]interface{}{
	"num":     testnum,
	"uint":    float64(testuint),
	"float":   float64(testfloat),
	"str":     teststr,
	"dt":      testdatetime,
	"offset":  testoffset,
	"dtwrong": testdtwrong,
	"bool":    testbool,
	"notype":  testnotype,
	"array":   testarr,
	"obj":     testobj,
	"null":    nil,
}

func TestMissingField(t *testing.T) {
	_, err := GetString("missing", rootObj)
	if err == nil || err != ValueNotFound {
		t.Errorf("expected ValueNotFound error got %v", err)
	}
}

func TestTypeOf(t *testing.T) {
	var tt JsonType
	if tt = TypeOf(testnum); tt != JNumber {
		t.Errorf("expected jnumber got %v", tt)
	}
	if tt = TypeOf(testbool); tt != JBool {
		t.Errorf("expected jbool got %v", tt)
	}
	if tt = TypeOf(teststr); tt != JString {
		t.Errorf("expected jstring got %v", tt)
	}
	if tt = TypeOf(testarr); tt != JArr {
		t.Errorf("expected jarr got %v", tt)
	}
	if tt = TypeOf(rootObj); tt != JObj {
		t.Errorf("expected jobj got %v", tt)
	}
	if tt = TypeOf(nil); tt != JNull {
		t.Errorf("exxpected jnull got %v", tt)
	}
	if tt = TypeOf(testnotype); tt != JUnknown {
		t.Errorf("expected junknown got %v", tt)
	}
}

func TestGetArray(t *testing.T) {
	if f, err := GetArray("array", rootObj); err != nil {
		t.Errorf("expected bool got error %v", err)
	} else if len(testarr) != len(f) {
		t.Errorf("length of array does not match")
	}
	if _, err := GetArray("notype", rootObj); err == nil || err != IncorrectType {
		t.Errorf("expected IncorrectType error got %v", err)
	}
}

func TestScanArray(t *testing.T) {
	var a JSONArray
	if err := ScanArray("array", rootObj, &a); err != nil {
		t.Errorf("unexpected error %v", err)
	} else if len(testarr) != len(a) {
		t.Errorf("length of array does not match")
	}
}

func TestGetDataTime(t *testing.T) {
	//without offset
	if dt, err := GetDateTime("dt", "", rootObj); err != nil {
		t.Errorf("expected obj get error %v", err)
	} else if v := dt.Format(ExifDateTime); v != testdatetime {
		t.Errorf("expected %v got %v", testdatetime, v)
	}
	//test with offset
	if dt, err := GetDateTime("dt", "offset", rootObj); err != nil {
		t.Errorf("expected obj get error %v", err)
	} else if v := dt.Format(ExifDateTimeOffset); v != testdatetimeoffset {
		t.Errorf("expected %v got %v", testdatetime, v)
	}
	//test wrong format
	if dt, err := GetDateTime("dtwrong", "", rootObj); err == nil {
		t.Errorf("expected date format error, got %v,", dt)
	}
}

func ScanGetTime(t *testing.T) {
	var dt time.Time
	if err := ScanDateTime("dt", "", rootObj, &dt); err != nil {
		fmt.Errorf("unexpected error: %v", err)
	} else if v := dt.Format(ExifDateTime); v != testdatetime {
		fmt.Errorf("expected %v got %v", testdatetime, v)
	}
}

func TestGetObject(t *testing.T) {
	if f, err := GetObject("obj", rootObj); err != nil {
		t.Errorf("expected obj got error %v", err)
	} else if _, exists := f["val1"]; !exists {
		fmt.Errorf("expected key val1")
	}
	if _, err := GetObject("notype", rootObj); err == nil || err != IncorrectType {
		t.Errorf("expected IncorrectType error got %v", err)
	}
}

func TestScanObject(t *testing.T) {
	var o JSONObject
	if err := ScanObject("obj", rootObj, &o); err != nil {
		t.Errorf("unexpected error %v", err)
	} else if _, exists := o["val1"]; !exists {
		t.Errorf("expected key val 1")
	}
}

func TestGetNumber(t *testing.T) {
	f, err := GetNumber("num", rootObj)
	if err != nil {
		t.Errorf("expected number got error %v", err)
	}
	if f != testnum {
		t.Errorf("unexpected %v got %v", testnum, f)
	}
	f, err = GetNumber("notype", rootObj)
	if err == nil || err != IncorrectType {
		t.Errorf("expected IncorrectType error got %v", err)
	}
}

func TestScanNumber(t *testing.T) {
	var n float64 = 0
	err := ScanNumber("num", rootObj, &n)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if n != testnum {
		t.Errorf("expected %v, got %v", testnum, n)
	}
}

func TestGetBool(t *testing.T) {
	if f, err := GetBool("bool", rootObj); err != nil {
		t.Errorf("expected bool got error %v", err)
	} else if f != testbool {
		t.Errorf("unexpected %v got %v", testbool, f)
	}
	if _, err := GetBool("notype", rootObj); err == nil || err != IncorrectType {
		t.Errorf("expected IncorrectType error got %v", err)
	}
}

func TestScanBool(t *testing.T) {
	var b bool
	if err := ScanBool("bool", rootObj, &b); err != nil {
		t.Errorf("unexpected error %v", err)
	} else if b != testbool {
		t.Errorf("expected %v, got %v", testbool, b)
	}
}

func TestGetString(t *testing.T) {
	if f, err := GetString("str", rootObj); err != nil {
		t.Errorf("expected sting got error %v", err)
	} else if f != teststr {
		t.Errorf("unexpected %v got %v", teststr, f)
	}
	if _, err := GetString("notype", rootObj); err == nil || err != IncorrectType {
		t.Errorf("expected IncorrectType error got %v", err)
	}
}

func TestScanString(t *testing.T) {
	var s string
	if err := ScanString("str", rootObj, &s); err != nil {
		t.Errorf("unexpected error %v", err)
	} else if s != teststr {
		t.Errorf("expected %v, got %v", teststr, s)
	}
}

func TestGetUInt(t *testing.T) {
	f, err := GetUInt("uint", rootObj)
	if err != nil {
		t.Errorf("expected number got error %v", err)
	}
	if f != testuint {
		t.Errorf("unexpected %v got %v", testuint, f)
	}
	f, err = GetUInt("notype", rootObj)
	if err == nil || err != IncorrectType {
		t.Errorf("expected IncorrectType error got %v", err)
	}
}

func TestScanUint(t *testing.T) {
	var n uint = 0
	err := ScanUInt("uint", rootObj, &n)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if n != testuint {
		t.Errorf("expected %v, got %v", testuint, n)
	}
}
