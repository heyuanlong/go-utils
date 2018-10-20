package utils

import (
	"math/rand"
	"time"
	"github.com/Jeffail/gabs"
	"errors"

	"fmt"
	"bytes"
	"sort"
	"strconv"
)


func GetRandomString(lens int) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetJsonString(body []byte,path string)(string ,error){
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return "", err
	}
	v, ok := jsonParsed.Path(path).Data().(string)
	if !ok {
		return "", errors.New("get value fail")
	}
	return v, nil
}
func GetJsonInt64(body []byte,path string)(int64 ,error){
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return 0, err
	}
	v, ok := jsonParsed.Path(path).Data().(int64)
	if !ok {
		return 0, errors.New("get value fail")
	}
	return v, nil
}
func GetJsonFloat64(body []byte,path string)(float64 ,error){
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return 0, err
	}
	v, ok := jsonParsed.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}
	return v, nil
}



func ChangeMapToURLParam(param map[string]interface{}) (string ,error){
	paramM := make(map[string]string)
	for k, v := range param {
		switch val := v.(type) {
		case string:
			paramM[k] = val
		case bool:
			paramM[k] = strconv.FormatBool(val)
		case int:
			paramM[k] = strconv.FormatInt(int64(val), 10)
		case int32:
			paramM[k] = strconv.FormatInt(int64(val), 10)
		case int64:
			paramM[k] = strconv.FormatInt(int64(val), 10)
		case float32:
			paramM[k] = strconv.FormatFloat(float64(val), 'f', -1, 64)
		case float64:
			paramM[k] = strconv.FormatFloat(float64(val), 'f', -1, 64)
		default:
			//klog.Warn.Println(k, v)
			return "",errors.New("not find value type")
		}
	}
	lens :=  len(paramM)
	keys := make([]string, 0,lens)
	for k := range paramM {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b bytes.Buffer

	for i, k := range keys {
		if i == (lens - 1){
			fmt.Fprintf(&b, "%s=%s", k, paramM[k])
		}else{
			fmt.Fprintf(&b, "%s=%s&", k, paramM[k])
		}
	}
	return b.String(),nil
}