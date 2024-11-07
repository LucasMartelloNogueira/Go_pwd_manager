package domain

type CsvSerializable[T any] interface {
	FromRecord(record []string) (*T, error)
	ToRecord() ([]string, error)
}

