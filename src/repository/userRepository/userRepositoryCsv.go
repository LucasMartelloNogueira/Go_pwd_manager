package userRepository

import (
	"CustomErrors"
	"domain"
	"strconv"
	"util"
)

type UserRepositoryCsv struct {

}


func (repository UserRepositoryCsv) FindById(id int) (domain.User, error){
	var filename string =  "../../storage/users.csv"
	userRecords, err := util.GetCsvRecords(filename)

	if err != nil {
		return domain.User{}, err
	}

	for _, userRecord := range userRecords {
		userId, _ := strconv.Atoi(userRecord[0])
		if userId == id {
			return util.RecordToUser(userRecord), nil
		}
	}

	return domain.User{}, customErrors.ErrNotFound

}

func (repository UserRepositoryCsv) Create(user *domain.CreateUserBody) (domain.User, error) {
	var filename string =  "../../storage/users.csv"
	userRecords, err := util.GetCsvRecords(filename)

	if err != nil {
		return domain.User{}, err
	}

	var newId int = len(userRecords)
	newUser := domain.User{
		Id: newId,
		Name: user.Name,
		Email: user.Email,
		MasterPassword: user.MasterPassword,
	}

	newUserRecord := util.UserToRecord(newUser)
	userRecords = append(userRecords, newUserRecord)

	err = util.WriteCsvRecords(filename, userRecords)
	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}

func (repository UserRepositoryCsv) DeleteById(id int) domain.User{
	return domain.User{}
}

func (repository UserRepositoryCsv) Update(user domain.User) domain.User {
	return domain.User{}
}