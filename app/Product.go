package app

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
	Status string
}

func (p *Product) IsValid() (bool, error) {
	return false, nil
}

func (p *Product) Enable() error {
	return nil
}

func (p *Product) Disable() error {
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

func (p *Product) GetStatus() string {
	return p.Status
}
