package domain

import (
	"strconv"
)

type User struct {
	Id		 int `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CsvSerializable implementation

func (user User) ToRecord() ([]string, error) {
	record := make([]string, 4)
	record[0] = strconv.Itoa(user.Id)
	record[1] = user.Name
	record[2] = user.Email
	record[3] = user.Password

	return record, nil
}

func (user User) FromRecord(record []string) (*User, error) {
	id, err := strconv.Atoi(record[0])

	if err != nil {
		return nil, err
	}

	newUser := new(User)
	newUser.Id = id
	newUser.Name = record[1]
	newUser.Email = record[2]
	newUser.Password = record[3]

	return newUser, nil
}
