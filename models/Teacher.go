package models

// Teacher The school headmaster has a null name and surname.
type Teacher struct {
	Id          int
	Name        *string
	Surname     *string
	DisplayName string
	Position    int
	Description string
}
