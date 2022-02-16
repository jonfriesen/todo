package todo

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/gofrs/uuid"
)

func init() {
	fmt.Println("starting todo package init")
}

var (
	ErrNotFound = errors.New("todo item not found")
	ErrExists   = errors.New("todo item already exists")
)

var todos = map[string]string{
	"a": "Pickup dry cleaning",
	"b": "Groceries",
	"c": "Fight a bear",
}

func Get(id string) (string, error) {
	item, e := todos[id]
	if !e {
		return "", ErrNotFound
	}

	return item, nil
}

func Set(value string) (map[string]string, error) {
	for _, v := range todos {
		if v == value {
			return nil, ErrExists
		}
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "generating item uuid")
	}

	todos[id.String()] = value

	return todos, nil
}

func List() map[string]string {
	return todos
}

func Complete(id string) (map[string]string, error) {
	_, e := todos[id]
	if !e {
		return nil, ErrNotFound
	}

	delete(todos, id)

	return todos, nil
}
