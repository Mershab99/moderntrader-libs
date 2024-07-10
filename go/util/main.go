package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

func UnmarshalArrayIntoStruct(data []byte, v interface{}) error {
	// Unmarshal into a slice of slices of interfaces
	var rawData [][]interface{}
	err := json.Unmarshal(data, &rawData)
	if err != nil {
		return err
	}

	// Get the reflect.Value of the slice to be populated
	sliceVal := reflect.ValueOf(v).Elem()

	// Get the type of the struct
	structType := sliceVal.Type().Elem()

	// Iterate over rawData and map to struct
	for _, item := range rawData {
		if len(item) < structType.NumField() {
			return fmt.Errorf("unexpected item length")
		}

		// Create a new instance of the struct
		structVal := reflect.New(structType).Elem()

		// Iterate over the fields and set the values
		for i := 0; i < structType.NumField(); i++ {
			field := structVal.Field(i)
			switch field.Kind() {
			case reflect.Int32:
				field.SetInt(int64(item[i].(float64)))
			case reflect.String:
				field.SetString(item[i].(string))
			case reflect.Bool:
				field.SetBool(item[i].(bool))
			case reflect.Struct:
				if field.Type() == reflect.TypeOf(sql.NullString{}) {
					field.Set(reflect.ValueOf(sql.NullString{String: item[i].(string), Valid: item[i].(string) != ""}))
				} else if field.Type() == reflect.TypeOf(time.Time{}) {
					t, err := time.Parse(time.RFC3339, item[i].(string))
					if err != nil {
						return err
					}
					field.Set(reflect.ValueOf(t))
				}
			case reflect.Slice:
				if field.Type().Elem().Kind() == reflect.String {
					field.Set(reflect.ValueOf(ConvertToStringSlice(item[i].([]interface{}))))
				}
			}
		}

		// Append the struct to the slice
		sliceVal.Set(reflect.Append(sliceVal, structVal))
	}

	return nil
}

// Helper function to convert interface slice to string slice
func ConvertToStringSlice(arr []interface{}) []string {
	var result []string
	for _, item := range arr {
		if str, ok := item.(string); ok {
			result = append(result, str)
		}
	}
	return result
}
