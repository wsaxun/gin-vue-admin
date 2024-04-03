package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}

func Pointer[T any](in T) (out *T) {
	return &in
}

func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// MaheHump 将字符串转换为驼峰命名
func MaheHump(s string) string {
	words := strings.Split(s, "-")

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}

	return strings.Join(words, "")
}

// 随机字符串
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[RandomInt(0, len(letters))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// FilterFields 根据struct 定义的json tag过滤结构体字段
func FilterFields(data any, jsonTags ...string) []map[string]any {
	var results []map[string]any

	// 获取数据的反射值
	dataValue := reflect.ValueOf(data)

	// 如果数据不是切片类型，则创建一个包含单个元素的切片
	if dataValue.Kind() != reflect.Slice {
		sliceType := reflect.SliceOf(reflect.TypeOf(data))
		slice := reflect.MakeSlice(sliceType, 1, 1)
		slice.Index(0).Set(dataValue)
		dataValue = slice
	}

	for i := 0; i < dataValue.Len(); i++ {
		// 获取当前元素的值
		elemValue := dataValue.Index(i)

		// 获取当前元素的类型
		elemType := elemValue.Type()

		result := make(map[string]interface{})

		for j := 0; j < elemValue.NumField(); j++ {
			// 获取字段的值
			fieldValue := elemValue.Field(j)

			// 获取字段的类型
			fieldType := elemType.Field(j)

			for _, jsonTag := range jsonTags {
				if tagValue := fieldType.Tag.Get("json"); tagValue == jsonTag {
					result[jsonTag] = fieldValue.Interface()
					break
				}
			}
		}

		results = append(results, result)
	}

	return results
}
