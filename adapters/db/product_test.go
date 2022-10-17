package db_test

import (
	"database/sql"
	"github.com/jvveiga/tests-arch-hexagonal/adapters/db"
	"github.com/jvveiga/tests-arch-hexagonal/app"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
    	"id" string,
    	"name" string,
    	"price" float,
    	"status" int
	)`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES("klp", "Product 1", 0, 0)`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("klp")
	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, float32(0.0), product.GetPrice())
	require.Equal(t, int8(0), product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := app.NewProduct()
	product.Name = "Product 1"
	product.Price = 10

	productCreate, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productCreate.GetName())
	require.Equal(t, product.Price, productCreate.GetPrice())
	require.Equal(t, product.Status, productCreate.GetStatus())

	product.Status = app.ENABLED

	productUpdate, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productUpdate.GetName())
	require.Equal(t, product.Price, productUpdate.GetPrice())
	require.Equal(t, product.Status, productUpdate.GetStatus())
}
