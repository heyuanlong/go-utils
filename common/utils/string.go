package utils

import (
	"errors"
	"math/rand"
	"time"

	"github.com/Jeffail/gabs"

	"bytes"
	"fmt"
	"sort"
	"strconv"
)

//0：数字+大小写字母，1：数字+小写字母，2：数字+大写字母，3：数字
func GetRandomString(lens int, types int) string {
	var str string
	if types == 1 {
		str = "0123456789abcdefghijklmnopqrstuvwxyz"
	} else if types == 2 {
		str = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if types == 3 {
		str = "0123456789"
	} else {
		str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//data.str
func GetJsonString(body []byte, path string) (string, error) {
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

//data.int64
func GetJsonInt64(body []byte, path string) (int64, error) {
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return 0, err
	}
	v, ok := jsonParsed.Path(path).Data().(float64)
	if !ok {
		return 0, errors.New("get value fail")
	}
	return int64(v), nil
}

//data.float64
func GetJsonFloat64(body []byte, path string) (float64, error) {
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

//map转为url参数，带升序排序
func ChangeMapToURLParam(param map[string]interface{}) (string, error) {
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
			return "", errors.New("not find value type")
		}
	}
	lens := len(paramM)
	keys := make([]string, 0, lens)
	for k := range paramM {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b bytes.Buffer

	for i, k := range keys {
		if i == (lens - 1) {
			fmt.Fprintf(&b, "%s=%s", k, paramM[k])
		} else {
			fmt.Fprintf(&b, "%s=%s&", k, paramM[k])
		}
	}
	return b.String(), nil
}
