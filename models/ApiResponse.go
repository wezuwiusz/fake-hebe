package models

type ApiResponse struct {
	EnvelopeType       string
	Envelope           any
	Status             ApiResponseStatus
	RequestId          string
	Timestamp          int64
	TimestampFormatted string
	InResponseTo       *string
}

type ApiResponseStatus struct {
	Code    int
	Message string
}
