package utils

import "fakeHebe/models"

func ConstructApiResponseStatus(statusCode int) models.ApiResponseStatus {
	message := "OK"

	switch statusCode {
	case 100:
		message = "Użytkownik nie jest uprawniony do przeglądania żądanych danych"
	case 101:
		message = "Żądanie nie zostało prawidłowo związane z oczekiwanym typem"
	case 102:
		message = "Brakuje nagłówka"
	}

	return models.ApiResponseStatus{
		Code:    statusCode,
		Message: message,
	}
}
