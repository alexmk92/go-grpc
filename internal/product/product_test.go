package product

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"
)

func TestProductService(t *testing.T) {
    mockCtrl := gomock.NewController(t)
    defer mockCtrl.Finish()

    t.Run("tests get product by id", func (t *testing.T) {
        productStoreMock := NewMockStore(mockCtrl)
        id := "UUID-1"
        productStoreMock.
            EXPECT().
            GetProductById(id).
            Return(Product{
                ID: id,
            }, nil)

        productService := New(productStoreMock)
        product, err := productService.GetProductById(context.Background(), id)
        assert.NoError(t, err)
        assert.Equal(t, "UUID-1", product.ID)
    })

    t.Run("tests insert product", func (t *testing.T) {
        productStoreMock := NewMockStore(mockCtrl)
        id := "UUID-1"
        productStoreMock.
            EXPECT().
            InsertProduct(Product{ID:id,}).
            Return(Product{ID:id,}, nil)

        productService := New(productStoreMock)
        product, err := productService.InsertProduct(
            context.Background(),
            Product{
                ID: id,
            },
        )

        assert.NoError(t, err)
        assert.Equal(t, "UUID-1", product.ID)
    })

    t.Run("test delete product", func (t *testing.T) {
        productStoreMock := NewMockStore(mockCtrl)
        id := "UUID-1"
        productStoreMock.
            EXPECT().
            DeleteProduct(id).
            Return(nil)

        productService := New(productStoreMock)
        err := productService.DeleteProduct(id)
        assert.NoError(t, err)
    })
}
