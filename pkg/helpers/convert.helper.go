package helpers

import (
	"encoding/json"
	"fmt"
	"main/server/admin/view/shared"
	"reflect"
	"strconv"
)

func PtintToString(num *int) string {
	return strconv.Itoa(*num)
}

func Uint(num string) uint {
	return uint(Int(num))
}

func UintToPtInt(num uint) *int {
	integer := int(num)
	return &integer
}

func Print(data interface{}) error {
	jsonBytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonBytes))
	return nil
}

func Int(string string) int {
	result, _ := strconv.Atoi(string)
	return result
}

func String(integer int) string {
	result := strconv.Itoa(integer)
	return result
}

func Ustring(Uinteger uint) string {
	return String(int(Uinteger))
}

func Bstring(boolean bool) string {
	if boolean {
		return "true"
	}
	return "false"
}

func ToJson(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func GetPointedUInt(num int) *uint {
	unum := uint(num)
	return &unum
}
func GetPointedInt(num int) *int {
	return &num
}

func ToSelect[T any](data []T, keyField string, valField string) ([]shared.SelectDataType, error) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("data should be a slice")
	}

	var result []shared.SelectDataType
	for i := 0; i < val.Len(); i++ {
		item := val.Index(i).Interface()
		key := reflect.Indirect(reflect.ValueOf(item)).FieldByName(keyField)
		val := reflect.Indirect(reflect.ValueOf(item)).FieldByName(valField)

		if !key.IsValid() || !val.IsValid() {
			return nil, fmt.Errorf("invalid field names")
		}

		result = append(result, shared.SelectDataType{
			Key: fmt.Sprintf("%v", key.Interface()),
			Val: fmt.Sprintf("%v", val.Interface()),
		})
	}
	return result, nil
}
