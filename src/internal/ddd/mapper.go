package ddd

import "errors"

type Mapper interface {
	ToDomain()
}

func InvalidEntity() error {
	return errors.New("invalid entity")
}
