package dbhelp

import (
	"errors"
	"reflect"
	"universityServer/internal/models"
)

func GetRecordId() (string, error) {
	recordName := reflect.VisibleFields(reflect.TypeOf(models.Students_records{}))

	for _, value := range recordName {
		if value.Tag == "gorm:\"primaryKey\"" {
			return value.Name, nil
		}
	}
	return "", errors.New("struct error")

}
