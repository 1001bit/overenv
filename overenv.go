package overenv

import (
	"fmt"
	"os"
	"reflect"
)

const tagName = "env"

func Get(key string) string {
	return os.Getenv(key)
}

func LoadStruct(str any) error {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		if tag := field.Tag.Get(tagName); tag != "" {
			envVarValue := Get(tag)

			switch field.Type.Kind() {
			case reflect.String:
				strVal := reflect.ValueOf(str).Elem().FieldByName(field.Name)
				strVal.SetString(envVarValue)
			default:
				return fmt.Errorf("unsupported type %s for field %s", field.Type, field.Name)
			}
		}
	}

	return nil
}
