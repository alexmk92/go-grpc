package db

import (
	"fmt"
	"os"

	"github.com/alexmk92/go-grpc/internal/product"
	"github.com/jmoiron/sqlx"
)

type Store struct {
    db *sqlx.DB
}

// New - returns a new store or error
func New() (Store, error) {
    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbTable := os.Getenv("DB_TABLE")
    dbPort := os.Getenv("DB_PORT")
    dbSSLMode := os.Getenv("DB_SSL_MODE")

    connectionString := fmt.Sprintf(
        "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
        dbHost, dbPort, dbUsername, dbTable, dbPassword, dbSSLMode,
    )

    db, err := sqlx.Connect("postgres", connectionString)
    if err != nil {
        return Store{}, err
    }

    return Store{
        db: db,
    }, nil
}

func (s Store) GetProductByID(id string) (product.Product, error) {
    return product.Product{}, nil
}

func (s Store) InsertProduct (prod product.Product) (product.Product, error) {
    return product.Product{}, nil
}

func (s Store) DeleteProduct (id string) (error) {
    return nil
}
