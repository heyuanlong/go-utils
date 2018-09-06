package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/url"
)

func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

//urlencode
func Urlencode(s string) string {
	return url.QueryEscape(s)
}

func ReadFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

//字节转换成整形
func BytesToIntLittle(b []byte) (int, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	err := binary.Read(bytesBuffer, binary.LittleEndian, &x)
	if err != nil {
		return 0, err
	}
	return int(x), nil
}

//整形转换成字节
func IntToBytesLittle(n int) ([]byte, error) {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, x)
	if err != nil {
		return []byte{}, err
	}
	return bytesBuffer.Bytes(), nil
}

//整形转换成字节
func IntToBytesLittleNotError(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, x)
	if err != nil {
		return []byte{}
	}
	return bytesBuffer.Bytes()
}
