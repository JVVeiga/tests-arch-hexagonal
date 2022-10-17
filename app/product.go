package app

import "errors"

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float32
}

const (
	ENABLED  = 1
	DISABLED = 0
)

type Product struct {
	ID     string
	Name   string
	Price  float32
	Status int8
}

func (p *Product) IsValid() (bool, error) {
	return false, nil
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
