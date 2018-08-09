package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"bytes"
)

func JSONToStruct(buf *bytes.Buffer, padding string, in []byte) {
	var m = make(map[string]interface{})

	if err := json.Unmarshal(in, &m); err != nil {
		return
	}

	mapToStruct(buf, padding, m)
}

func interfaceToType(buf *bytes.Buffer, padding string, s interface{}) {
	switch  v := s.(type) {
	case string:
		buf.WriteString("string")
		return
	case int:
		buf.WriteString("int")
		return
	case int8:
		buf.WriteString("int8")
		return
	case int16:
		buf.WriteString("int16")
		return
	case int32:
		buf.WriteString("int32")
		return
	case int64:
		buf.WriteString("int64")
		return
	case uint:
		buf.WriteString("uint")
		return
	case uint8:
		buf.WriteString("uint8")
		return
	case uint16:
		buf.WriteString("uint16")
		return
	case uint32:
		buf.WriteString("uint32")
		return
	case uint64:
		buf.WriteString("uint64")
		return
	case float32:
		buf.WriteString(float32To(v))
		return
	case float64:
		buf.WriteString(float64To(v))
		return
	case bool:
		buf.WriteString("bool")
		return
	}

	vt := reflect.TypeOf(s)
	vv := reflect.ValueOf(s)
	switch vt.Kind() {
	case reflect.Slice:
		buf.WriteString("[]")
		interfaceToType(buf, padding, vv.Index(0).Interface())
		return
	case reflect.Map:
		mapToStruct(buf, padding, vv.Interface().(map[string]interface{}))
		return
	}

	panic(fmt.Sprintf("%v:%v to json err", reflect.TypeOf(s).Kind(), s))
}
