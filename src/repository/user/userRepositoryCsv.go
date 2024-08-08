package user

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
	usersDataFrame, err := util.GetDataFrame(filename)

	if err != nil {
		return domain.User{}, err
	}

	for _, userRecord := range usersDataFrame.Values{
		userId, _ := strconv.Atoi(userRecord[0])
		if userId == id {
			return util.RecordToUser(userRecord), nil
		}
	}

	return domain.User{}, CustomErrors.ErrNotFound

}

func (repository UserRepositoryCsv) Create(user *domain.CreateUserBody) (domain.User, error) {
	var filename string =  "../../storage/users.csv"
	usersDataFrame, err := util.GetDataFrame(filename)

	if err != nil {
		return domain.User{}, err
	}

	maxId := -1

	for _, r := range usersDataFrame.Values{
		id, _ := strconv.Atoi(r[0])
		if id > maxId {
			maxId = id
		}
	}

	newId := maxId + 1

	newUser := domain.User{
		Id: newId,
		Name: user.Name,
		Email: user.Email,
		MasterPassword: user.MasterPassword,
	}

	newUserRecord := util.UserToRecord(newUser)
	usersDataFrame.Values = append(usersDataFrame.Values, newUserRecord)

	err = util.WriteDataFrameToCsv(filename, usersDataFrame)
	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil

}

func (repository UserRepositoryCsv) DeleteById(id int) (domain.User, error){
	var filename string =  "../../storage/users.csv"
	usersDataFrame, err := util.GetDataFrame(filename)

	if err != nil {
		return domain.User{}, err
	}

	deleteIndex := -1
	var deletedUser domain.User = domain.User{}

	for i, userRecord := range usersDataFrame.Values {

		userId, _ := strconv.Atoi(userRecord[0])
		if userId == id {
			deleteIndex = i
			deletedUser = util.RecordToUser(userRecord)
			break
		}
	}

	if deleteIndex == -1 {
		return domain.User{}, CustomErrors.ErrNotFound
	}

	usersDataFrame.Values = append(usersDataFrame.Values[:deleteIndex], usersDataFrame.Values[deleteIndex+1:]...)

	err = util.WriteDataFrameToCsv(filename, usersDataFrame)
	if err != nil {
		return domain.User{}, err
	}
	
	return deletedUser, nil

}

func (repository UserRepositoryCsv) Update(user domain.User) (domain.User, error) {
	var filename string =  "../../storage/users.csv"
	usersDataFrame, err := util.GetDataFrame(filename)

	if err != nil {
		return domain.User{}, err
	}

	for i, userRecord := range usersDataFrame.Values {
		userId, _ := strconv.Atoi(userRecord[0])
		if userId == user.Id {
			usersDataFrame.Values[i] = util.UserToRecord(user)
			break
		}
	}

	err = util.WriteDataFrameToCsv(filename, usersDataFrame)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}