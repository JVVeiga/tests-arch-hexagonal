package db

import (
	"database/sql"
	"github.com/jvveiga/tests-arch-hexagonal/app"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(ID string) (app.ProductInterface, error) {
	var product app.Product
	stmt, err := p.db.Prepare(`SELECT p.id, p.name, p.price, p.status FROM products AS p WHERE p.id = ?`)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(ID).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product app.ProductInterface) (app.ProductInterface, error) {
	var rows int
	var err error

	p.db.QueryRow(`SELECT p.id FROM products AS p WHERE p.id = ?`, product.GetID()).Scan(&rows)
	if rows == 0 {
		_, err = p.create(product)
	} else {
		_, err = p.update(product)
	}
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) create(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare(`INSERT INTO products (id, name, price, status) values (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product app.ProductInterface) (app.ProductInterface, error) {
	_, err := p.db.Exec(`UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?`,
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	return product, nil

}
