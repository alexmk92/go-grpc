//go:generate mockgen -destination=product_mocks_test.go -package=product github.com/alexmk92/go-grpc/internal/product Store
package product


import "context"

type Product struct {
    ID string
    Name string
    Type string
    Orders int
}

// Defines the interface we expect our db to conform to
type Store interface {
    GetProductByID(id string) (Product, error)
    InsertProduct(product Product) (Product, error)
    DeleteProduct(id string) error
}

// Product service responsible for updating
// the product db
type Service struct {
    Store Store
}

// Returns a new product service
func New(store Store) Service {
    return Service{
        Store: store,
    }
}

// Retrieves a product based on the ID
func (s Service) GetProductById(ctx context.Context, id string) (Product, error) {
    product, err := s.Store.GetProductByID(id)
    if err != nil {
        return Product{}, err
    }

    return product, nil
}

// Inserts a new product
func (s Service) InsertProduct(ctx context.Context, product Product) (Product, error) {
    product, err := s.Store.InsertProduct(product)
    if err != nil {
        return Product{}, err
    }
    return product, nil
}

// Delete a product
func (s Service) DeleteProduct(id string) error {
    err := s.Store.DeleteProduct(id)
    if err != nil {
        return err
    }
    return nil
}
