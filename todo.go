package todo

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultDBUrl = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

type TodoItem struct {
	gorm.Model
	Value     string `gorm:"uniqueIndex"`
	Completed bool
}

var db *gorm.DB

func init() {
	fmt.Println("starting todo package init")
	dbURL := defaultDBUrl
	if d := os.Getenv("DATABASE_URL"); d != "" {
		dbURL = d
	}

	var err error
	db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&TodoItem{})
	if err != nil {
		panic(err)
	}

}

var (
	ErrNotFound = errors.New("todo item not found")
	ErrExists   = errors.New("todo item already exists")
)

func Get(id string) (*TodoItem, error) {
	var i *TodoItem
	db.First(&i, id)

	if i == nil {
		return nil, ErrNotFound
	}

	return i, nil
}

func Set(value string) ([]TodoItem, error) {
	db.Create(&TodoItem{
		Value: value,
	})

	return List(false), nil
}

func List(includeCompleted bool) []TodoItem {
	var items []TodoItem

	if includeCompleted {
		db.Find(&items)
	} else {
		db.Where("completed = ?", false).Find(&items)
	}

	return items
}

func Complete(id string) ([]TodoItem, error) {
	db.Model(&TodoItem{}).Where("id = ?", id).Update("completed", true)

	return List(false), nil
}
