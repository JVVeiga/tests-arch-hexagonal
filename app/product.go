package app

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	valid.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() int8
	GetPrice() float32
}

type ProductServiceInterface interface {
	Get(ID string) (ProductInterface, error)
	Create(name string, price float32) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReader interface {
	Get(ID string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

const (
	ENABLED  = 1
	DISABLED = 0
)

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float32 `valid:"float,optional"`
	Status int8    `valid:"int,optional"`
}

func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}
	return &product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status < 0 {
		p.Status = DISABLED
	}
	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal zero")
	}

	_, err := valid.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price <= 0 {
		return errors.New("the price must be greater than zero to enable the product")
	}
	p.Status = ENABLED
	return nil
}

func (p *Product) Disable() error {
	if p.Price != 0 {
		return errors.New("the price must be zero in order to have the product disable")
	}
	p.Status = DISABLED
	return nil
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) GetStatus() int8 {
	return p.Status
}
