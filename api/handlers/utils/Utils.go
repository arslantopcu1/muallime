package utils

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/ehsaniara/gointerlock"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func DeleteUserTag(c *fiber.Ctx) {

}

func GetType(myvar interface{}) string {
	valueOf := reflect.ValueOf(myvar)

	if valueOf.Type().Kind() == reflect.Ptr {
		return reflect.Indirect(valueOf).Type().Name()
	} else {
		return valueOf.Type().Name()
	}
}

// Simple helper function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := GetEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := GetEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}

func CronJob(name string, funcName func(), time time.Duration, redisConnector *redis.Client) {

	log.Infof("START JOB: %s", name)

	ctx := context.Background()
	var job = gointerlock.GoInterval{
		LockVendor:     gointerlock.RedisLock,
		Name:           name,
		Interval:       time,
		Arg:            funcName,
		RedisConnector: redisConnector,
	}

	job.Run(ctx)
}

func In_array(val string, array []string) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if val == v {
			index = i
			exists = true
			return
		}
	}

	return
}

func TR_TO_ENG(text string) string {
	text = strings.Replace(text, "Ö", "O", -1)
	text = strings.Replace(text, "Ş", "S", -1)
	text = strings.Replace(text, "Ğ", "G", -1)
	text = strings.Replace(text, "Ç", "C", -1)
	text = strings.Replace(text, "İ", "I", -1)
	text = strings.Replace(text, "Ü", "U", -1)
	return text
}

func ToLowerTR(text string) string {
	//lowercase := "abcçdefgğhıijklmnoöprsştuüvyz"
	//uppercase := "ABCÇDEFGĞHIİJKLMNOÖPRSŞTUÜVYZ"
	return strings.ToLowerSpecial(unicode.TurkishCase, strings.TrimSpace(text))
}

func GUnzipData(data []byte) (resData []byte, err error) {
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData = resB.Bytes()

	return
}

func GZipData(data []byte) (compressedData []byte, err error) {
	var b bytes.Buffer

	gz, _ := gzip.NewWriterLevel(&b, gzip.BestSpeed)

	_, err = gz.Write(data)
	if err != nil {
		return
	}

	if err = gz.Flush(); err != nil {
		return
	}

	if err = gz.Close(); err != nil {
		return
	}

	compressedData = b.Bytes()

	return
}
