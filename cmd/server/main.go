package main

import (
	"log"

	"github.com/alexmk92/go-grpc/internal/db"
	"github.com/alexmk92/go-grpc/internal/product"
)

func Run() error {
    // responsible for init and starting gRPC server
    productStore, err := db.New()
    if err != nil {
        return err
    }
    productStore.Migrate()
    if err != nil {
        log.Println("Failed to run migrations")
        return err
    }

    _ = product.New(productStore)

    return nil
}

func main () {
    if err := Run(); err != nil {
        log.Fatal(err)
    }
}
