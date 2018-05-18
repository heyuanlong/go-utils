package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"io/ioutil"
)



func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

//urlencode
func Urlencode(s string) string{
	return url.QueryEscape(s)
}

func ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
