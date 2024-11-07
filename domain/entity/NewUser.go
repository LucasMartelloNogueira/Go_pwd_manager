package domain

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


func (newUser NewUser) ToRecord() ([]string, error) {
	record := make([]string, 3)
	record[0] = newUser.Name
	record[1] = newUser.Email
	record[2] = newUser.Password

	return record, nil
}

func (user NewUser) FromRecord(record []string) (*NewUser, error) {
	newUser := new(NewUser)
	newUser.Name = record[0]
	newUser.Email = record[1]
	newUser.Password = record[2]

	return newUser, nil
}