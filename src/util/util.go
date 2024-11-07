package util

import (
	"domain/entity"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func GetRequestBody(r *http.Request, body any) error {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("error reading request body")
	}

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return errors.New(fmt.Sprintf("error converting request body to entity %v\n", body))
	}
	return nil
}

func RecordToUser(record []string) (*domain.User, error) {

	id, err := strconv.Atoi(record[0])

	if err != nil {
		return nil, errors.New("error converting userId (string) to int")
	}

	var name string = record[1]
	var email string = record[2]
	var password = record[3]

	var user *domain.User = new(domain.User)
	user.Id = id
	user.Name = name
	user.Email = email
	user.Password = password

	return user, nil
}

func UserToRecord(user *domain.User) []string {
	return []string{
		strconv.FormatInt(int64(user.Id), 10),
		user.Name,
		user.Email,
		user.Password,
	}
}

func GetDataFrame(filename string) (*DataFrame, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var dataFrame *DataFrame = new(DataFrame)
	dataFrame.Columns = records[0]
	dataFrame.Values = records[1:]

	file.Close()
	return dataFrame, nil
}

func WriteDataFrameToCsv(filename string, dataFrame *DataFrame) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	var records [][]string = [][]string{dataFrame.Columns}
	records = append(records, dataFrame.Values...)

	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)

	file.Close()
	return err
}
