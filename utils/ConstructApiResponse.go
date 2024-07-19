package utils

import (
	"fakeHebe/models"
	"github.com/google/uuid"
	"reflect"
	"time"
)

func ConstructApiResponse(envelope any, statusCode int) models.ApiResponse {
	cSharpType := ""
	if envelope != nil {
		cSharpType = ConvertGoTypeToCSharpType(reflect.TypeOf(envelope).Kind())
	} else {
		cSharpType = "null"
	}
	return models.ApiResponse{
		EnvelopeType:       cSharpType,
		Envelope:           envelope,
		Status:             ConstructApiResponseStatus(statusCode),
		RequestId:          uuid.New().String(),
		Timestamp:          time.Now().UnixMilli(),
		TimestampFormatted: time.Now().Format("2006-01-02 15:04:05"),
		InResponseTo:       nil,
	}
}

func ConvertGoTypeToCSharpType(typeName reflect.Kind) string {
	switch typeName {
	case reflect.Bool:
		return "Boolean"
	case reflect.Array:
		return "IEnumerable`1"
	default:
		return typeName.String()
	}
}
