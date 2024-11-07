package repository

import (
	"CustomErrors"
	types "domain/types"
	"strconv"
	"util"
	"fmt"
)

type CsvRepository[T types.CsvSerializable[T], K types.CsvSerializable[K]] struct {
	filename string
}

func (repository CsvRepository[T, K]) FindById(id int) (*T, error) {
	dataFrame, err := util.GetDataFrame(repository.filename)

	if err != nil {
		return nil, err
	}

	for _, record := range dataFrame.Values {
		recordId, _ := strconv.Atoi(record[0])
		if recordId == id {
			entity, err := T.FromRecord(*(new(T)), record)
			if err != nil {
				return nil, err
			}
			return entity, nil
		}
	}

	return nil, CustomErrors.ErrNotFound
}

func (repository CsvRepository[T, K]) FindByColumn(column string, value string) (*T, error) {

	dataFrame, err := util.GetDataFrame(repository.filename)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	var columnId int = -1

	for i, col := range dataFrame.Columns {
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

	for _, record := range dataFrame.Values {
		if record[columnId] == value {
			entity, err := T.FromRecord(*(new(T)), record)
			if err != nil {
				return nil, err
			}
			return entity, nil
		}
	}

	return nil, CustomErrors.ErrNotFound
}

func (repository CsvRepository[T, K]) Create(newEntity *K) (*T, error) {
	dataFrame, err := util.GetDataFrame(repository.filename)

	if err != nil {
		return nil, err
	}

	maxId := -1

	for _, record := range dataFrame.Values {
		id, _ := strconv.Atoi(record[0])
		if id > maxId {
			maxId = id
		}
	}

	newId := maxId + 1

	newEntityRecord, err := K.ToRecord(*newEntity)
	entityRecord := make([]string, 1)
	entityRecord[0] = strconv.Itoa(newId)
	entityRecord = append(entityRecord, newEntityRecord...)

	if err != nil {
		return nil, CustomErrors.ErrInternalServerError
	}

	dataFrame.Values = append(dataFrame.Values, entityRecord)

	err = util.WriteDataFrameToCsv(repository.filename, dataFrame)
	if err != nil {
		return nil, err
	}

	entity, err := T.FromRecord(*new(T), entityRecord)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (repository CsvRepository[T, K]) DeleteById(id int) (*T, error) {
	dataFrame, err := util.GetDataFrame(repository.filename)

	if err != nil {
		return nil, err
	}

	deleteIndex := -1
	var deletedEntity *T = new(T);

	for i, record := range dataFrame.Values {

		recordId, err := strconv.Atoi(record[0])
		
		if err != nil {
			return nil, err
		}

		if recordId == id {
			deleteIndex = i
			deletedEntity, err = T.FromRecord(*deletedEntity, record)

			if err != nil {
				return nil, err
			}
			break
		}
	}

	if deleteIndex == -1 {
		return nil, CustomErrors.ErrNotFound
	}

	dataFrame.Values = append(dataFrame.Values[:deleteIndex], dataFrame.Values[deleteIndex+1:]...)

	err = util.WriteDataFrameToCsv(repository.filename, dataFrame)
	if err != nil {
		return nil, err
	}

	return deletedEntity, nil
}

func (repository CsvRepository[T, K]) Update(id int, entity *T) (*T, error) {
	dataFrame, err := util.GetDataFrame(repository.filename)

	if err != nil {
		return nil, err
	}

	for i, record := range dataFrame.Values {
		recordId, err := strconv.Atoi(record[0])

		if err != nil {
			return nil, err
		}

		if recordId == id {
			entityRecord, err := T.ToRecord(*entity)
			
			if err != nil {
				return nil, err
			}

			dataFrame.Values[i] = entityRecord
			break
		}
	}

	err = util.WriteDataFrameToCsv(repository.filename, dataFrame)
	if err != nil {
		return nil, err
	}

	return entity, nil
}