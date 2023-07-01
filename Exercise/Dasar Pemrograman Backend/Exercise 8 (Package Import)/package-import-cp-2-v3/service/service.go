package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	product, err1 := s.database.GetProductByName(productName)
	cart, err2 := s.database.GetCartItems()

	if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
	} else {
		if quantity <= 0 {
			return errors.New("invalid quantity")
		}
		cart = append(cart, entity.CartItem{ProductName: productName, Price: product.Price, Quantity: quantity})
		
		err3 := s.database.SaveCartItems(cart)
		if err3 != nil {
			return err1
		}
		return nil
	}
}

func (s *Service) RemoveCart(productName string) error {
	product, err1 := s.database.GetProductByName(productName)
	cart, err2 := s.database.GetCartItems()

	if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
	} else {
		carts := []entity.CartItem{}
		productExist := false

		for i := 0; i < len(cart); i++ {
			if cart[i].ProductName == product.Name && cart[i].Price == product.Price {
				productExist = true
				continue
			}
			carts = append(carts, cart[i])
		}

		if !productExist {
			return errors.New("product not found")
		}
		
		err3 := s.database.SaveCartItems(carts)
		if err3 != nil {
			return err1
		}
		return nil
	}
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	return s.database.SaveCartItems([]entity.CartItem{})
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	cart := s.database.GetProductData()
	return cart, nil
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	cart, err := s.ShowCart()
	if err != nil {
		return entity.PaymentInformation{}, nil
	}

	totalPrice := 0
	for _, product := range cart {
		totalPrice += product.Price * product.Quantity
	}
	change := money - totalPrice

	if change < 0 {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}
	s.ResetCart()
	return entity.PaymentInformation{ProductList: cart, TotalPrice: totalPrice, MoneyPaid: money, Change: change}, nil 
}
