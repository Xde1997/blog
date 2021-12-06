package sqlutils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func GetRow(row *sql.Rows) map[int]map[string]string {
	col, _ := row.Columns()
	vals := make([][]byte, len(col))
	scans := make([]interface{}, len(col))
	for k := range col {
		scans[k] = &vals[k]
	}
	result := make(map[int]map[string]string)
	i := 0
	for row.Next() {
		row.Scan(scans...)
		row := make(map[string]string)
		for k, v := range vals {
			key := col[k]
			row[key] = string(v)
		}
		result[i] = row
		i++
	}
	return result
}

func DataToStructBySqlTag(data map[string]string, ptr interface{}) {
	pvalue := reflect.ValueOf(ptr).Elem()
	ptype := reflect.TypeOf(ptr).Elem()
	for i := 0; i < pvalue.NumField(); i++ {
		value := data[ptype.Field(i).Tag.Get("sql")]
		val := reflect.ValueOf(value)
		//获取需要映射的字段名
		key := ptype.Field(i).Name
		//获取需要映射的字段的类型
		ktype := ptype.Field(i).Type
		ktypename := fmt.Sprint(ktype)
		//获取数据的类型
		vtype := reflect.TypeOf(value)
		if vtype != ktype {
			val, _ = ConversionType(value, ktypename)
		}
		pvalue.FieldByName(key).Set(val)
	}
}

func ConversionType(value string, ktype string) (reflect.Value, error) {
	if ktype == "string" {
		return reflect.ValueOf(ktype), nil
	} else if ktype == "int64" {
		buf, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(buf), err
	} else if ktype == "int32" {
		buf, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int32(buf)), err
	} else if ktype == "int8" {
		buf, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(buf)), err
	} else if ktype == "int" {
		buf, err := strconv.Atoi(value)
		return reflect.ValueOf(buf), err
	} else if ktype == "float32" {
		buf, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(buf)), err
	} else if ktype == "float64" {
		buf, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(buf), err
	} else if ktype == "bool" {
		buf, err := strconv.ParseBool(value)
		return reflect.ValueOf(buf), err
	} else if ktype == "time.Time" {
		var buf time.Time
		if len(value) < 10 {
			buf = time.Time{}
			return reflect.ValueOf(buf), nil
		}
		var temp = value[0:10] + " " + value[11:19]
		buf, err := time.ParseInLocation("2006-01-02 15:04:05", temp, time.Local)
		return reflect.ValueOf(buf), err
	} else if ktype == "Time" {
		var buf time.Time
		if len(value) < 10 {
			buf = time.Time{}
			return reflect.ValueOf(buf), nil
		}
		var temp = value[0:10] + " " + value[11:19]
		buf, err := time.ParseInLocation("2006-01-02 15:04:05", temp, time.Local)
		return reflect.ValueOf(buf), err
	} else if ktype == "map[string]interface {}" {
		buf, err := JsonToMap(value)
		return reflect.ValueOf(buf), err
	} else {
		return reflect.Value{}, errors.New("no such type")
	}
}

// Convert json string to map
func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
