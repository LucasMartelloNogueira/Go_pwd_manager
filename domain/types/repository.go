package domain

type Repository[T any, K any] interface {
	FindById(id int) (*T, error)
	FindByColumn(column string, value string) (*T, error)
	Create(entity *K) (*T, error)
	DeleteById(id int) (*T, error)
	Update(id int, entity *T) (*T, error)
}