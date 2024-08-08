package util

import (
	"domain"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func CheckErr(err error, msg string){
	if err != nil {
		log.Panic(msg)
	}
}

func GetRequestBody(r *http.Request, body any) {
	bodyBytes, err := io.ReadAll(r.Body)
	CheckErr(err, "[util] io error reading request body")

	err = json.Unmarshal(bodyBytes, &body)
	CheckErr(err, "[util] error reading body from request")
}

func RecordToUser(record []string) domain.User {

	id, err := strconv.Atoi(record[0])
	var name string = record[1]
	var email string = record[2]
	var masterPassword = record[3]

	CheckErr(err, "[util] error converting csv record string to int")

	return domain.User{
		Id: id,
		Name: name,
		Email: email,
		MasterPassword: masterPassword,
	}
}

func UserToRecord(user domain.User) []string{
	return []string{
		strconv.FormatInt(int64(user.Id), 10),
		user.Name,
		user.Email,
		user.MasterPassword,
	}
}

func GetDataFrame(filename string) (*DataFrame, error){
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