package chatbotbase

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

// JSON - make json to field
func JSON(key string, obj interface{}) zap.Field {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	s, err := json.Marshal(obj)
	if err != nil {
		return zap.Error(err)
	}

	return zap.String(key, string(s))
}

// BuildLogSubFilename -
func BuildLogSubFilename(appName string, version string) string {
	tm := time.Now()
	str := tm.Format("2006-01-02_15:04:05")
	return fmt.Sprintf("%v.%v.%v", appName, version, str)
}

// BuildLogFilename -
func BuildLogFilename(logtype string, subname string) string {
	return fmt.Sprintf("%v.%v.log", subname, logtype)
}

// AppendString - append string
func AppendString(strs ...string) string {
	var buffer bytes.Buffer

	for _, str := range strs {
		if len(str) > 0 {
			buffer.WriteString(str)
		}
	}

	return buffer.String()
}

// GetCurTime - append string
func GetCurTime() int64 {
	return time.Now().Unix()
}

// IndexOfArrayString - find a string in []string
func IndexOfArrayString(arr []string, str string) int {
	for i, v := range arr {
		if v == str {
			return i
		}
	}

	return -1
}

// Timestamp2Str - unix timestamp to string
func Timestamp2Str(ts int64) string {
	tm := time.Unix(ts, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// Int64ToBytes - int64 -> bytes
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// BytesToInt64 - bytes -> int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

// MD5Buffer - hash buffer
func MD5Buffer(buf []byte) string {
	h := md5.New()
	h.Write(buf)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// JSONFormat - json format string
func JSONFormat(obj interface{}) (string, error) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	s, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}

	return string(s), nil
}

// FormatFileSize - format size to string
func FormatFileSize(s int64) string {
	if s <= 0 {
		return "invalid size"
	}

	if s < 1024 {
		return fmt.Sprintf("%v b", s)
	} else if s < 1024*1024 {
		return fmt.Sprintf("%.2f k", float32(s)/1024)
	} else if s < 1024*1024*1024 {
		return fmt.Sprintf("%.2f m", float32(s)/1024/1024)
	} else if s < 1024*1024*1024*1024 {
		return fmt.Sprintf("%.2f g", float32(s)/1024/1024/1024)
	} else if s < 1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.2f t", float32(s)/1024/1024/1024/1024)
	} else if s < 1024*1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.2f p", float32(s)/1024/1024/1024/1024/1024)
	}

	return fmt.Sprintf("%v", s)
}

// FormatCommand - format command, /start => start
func FormatCommand(str string) string {
	str = strings.TrimSpace(str)

	if len(str) <= 1 {
		return str
	}

	if str[0] == '/' &&
		strings.IndexByte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			str[1]) >= 0 {

		return str[1:]
	}

	return str
}

// ID2Str - id -> string
func ID2Str(id int) string {
	return strconv.Itoa(id)
}

// Str2ID - string -> id
func Str2ID(str string) (int, error) {
	return strconv.Atoi(str)
}

// ID642Str - id64 -> string
func ID642Str(id int64) string {
	return strconv.FormatInt(id, 10)
}

// Str2ID64 - string -> id
func Str2ID64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
