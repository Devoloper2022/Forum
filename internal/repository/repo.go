package repository

type CRUD interface {
	Insert(interface{}) (int64, error)

	Delete(interface{}) error

	Update(interface{}) error

	Get(interface{}) error
}
