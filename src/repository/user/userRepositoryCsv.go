package user

import (
	"CustomErrors"
	"domain"
	"fmt"
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

func (repository UserRepositoryCsv) DeleteById(id int) (domain.User, error){
	fmt.Printf("id: %d\n", id)
	var filename string =  "../../storage/users.csv"
	userRecords, err := util.GetCsvRecords(filename)

	if err != nil {
		return domain.User{}, err
	}

	deleteIndex := -1
	var deletedUser domain.User = domain.User{}

	for i, userRecord := range userRecords {
		userId, _ := strconv.Atoi(userRecord[0])
		if userId == id {
			fmt.Printf("achou user de id %d\n", id)
			deleteIndex = i
			deletedUser = util.RecordToUser(userRecord)
			break
		}
	}

	if deleteIndex == -1 {
		return domain.User{}, customErrors.ErrNotFound
	}

	userRecords = append(userRecords[:deleteIndex], userRecords[deleteIndex+1:]...)
	fmt.Printf("userRecords: %v\n", userRecords)

	err = util.WriteCsvRecords(filename, userRecords)
	if err != nil {
		return domain.User{}, err
	}
	
	return deletedUser, nil

}

func (repository UserRepositoryCsv) Update(user domain.User) (domain.User, error) {
	var filename string =  "../../storage/users.csv"
	userRecords, err := util.GetCsvRecords(filename)

	if err != nil {
		return domain.User{}, err
	}

	for i, userRecord := range userRecords {
		userId, _ := strconv.Atoi(userRecord[0])
		if userId == user.Id {
			userRecords[i] = util.UserToRecord(user)
			break
		}
	}

	err = util.WriteCsvRecords(filename, userRecords)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil


}