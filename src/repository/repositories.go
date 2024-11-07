package repository

import (
	types "domain/types"
	entities "domain/entity"
	"util"
)

var (
	CsvUserReposiory types.Repository[entities.User, entities.NewUser] = CsvRepository[entities.User, entities.NewUser]{filename: util.Constants["USER_CSV_FILENAME"]}
)
