package user

import (
	"CustomErrors"
	"domain/entity"
	"strconv"
	"util"
)

type UserRepositoryCsv struct {
}

func (repository UserRepositoryCsv) FindById(id int) (*domain.UserWithId, error) {
	var filename string = "../../storage/users.csv"
	usersDataFrame, err := util.GetDataFrame(filename)

	if err != nil {
		return nil, err
	}

	for _, userRecord := range usersDataFrame.Values {
		userId, _ := strconv.Atoi(userRecord[0])
		if userId == id {
			user, err := util.RecordToUser(userRecord)
			if err != nil {
				return nil, err
			}
			return user, nil
		}
	}

	return nil, CustomErrors.ErrNotFound

}

func (repository UserRepositoryCsv) FindByColumn(column string, value string) (*domain.UserWithId, error) {
	var filename string = "../../storage/users.csv"
	usersDataFrame, err := util.GetDataFrame(filename)

	var columnId int = -1

	for i, col := range usersDataFrame.Columns {
		if col == column {
			columnId = i
		}
	}

	if columnId == -1 {
		return nil, CustomErrors.ErrInternalServerError
	}

	if err != nil {
		return nil, err
	}

	for _, userRecord := range usersDataFrame.Values {
		if userRecord[columnId] == value {
			user, err := util.RecordToUser(userRecord)
			if err != nil {
				return nil, err
			}
			return user, nil
		}
	}

	return nil, CustomErrors.ErrNotFound

}

func (repository UserRepositoryCsv) Create(user *domain.User) (*domain.UserWithId, error) {
	var filename string = "../../storage/users.csv"
	usersDataFrame, err := util.GetDataFrame(filename)

	if err != nil {
		return nil, err
	}

	maxId := -1

	for _, r := range usersDataFrame.Values {
		id, _ := strconv.Atoi(r[0])
		if id > maxId {
			maxId = id
		}
	}

	newId := maxId + 1

	var newUser *domain.UserWithId = new(domain.UserWithId)
	newUser.Id = newId
	newUser.Name = user.Name
	newUser.Password = user.Password

	newUserRecord := util.UserToRecord(newUser)
	usersDataFrame.Values = append(usersDataFrame.Values, newUserRecord)

	err = util.WriteDataFrameToCsv(filename, usersDataFrame)
	if err != nil {
		return nil, err
	}

	return newUser, nil

}

func (repository UserRepositoryCsv) DeleteById(id int) (*domain.UserWithId, error) {
	var filename string = "../../storage/users.csv"
	usersDataFrame, err := util.GetDataFrame(filename)

	if err != nil {
		return nil, err
	}

	deleteIndex := -1
	var deletedUser *domain.UserWithId = new(domain.UserWithId)

	for i, userRecord := range usersDataFrame.Values {

		userId, _ := strconv.Atoi(userRecord[0])
		if userId == id {
			deleteIndex = i
			deletedUser, err = util.RecordToUser(userRecord)

			if err != nil {
				return nil, err
			}
			break
		}
	}

	if deleteIndex == -1 {
		return nil, CustomErrors.ErrNotFound
	}

	usersDataFrame.Values = append(usersDataFrame.Values[:deleteIndex], usersDataFrame.Values[deleteIndex+1:]...)

	err = util.WriteDataFrameToCsv(filename, usersDataFrame)
	if err != nil {
		return nil, err
	}

	return deletedUser, nil

}

func (repository UserRepositoryCsv) Update(user *domain.UserWithId) (*domain.UserWithId, error) {
	var filename string = "../../storage/users.csv"
	usersDataFrame, err := util.GetDataFrame(filename)

	if err != nil {
		return nil, err
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
		return nil, err
	}

	return user, nil

}
