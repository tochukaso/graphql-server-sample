// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type EditProduct struct {
	ID     int     `json:"id"`
	Name   *string `json:"name"`
	Price  *int    `json:"price"`
	Code   *string `json:"code"`
	Detail *string `json:"detail"`
}

type NewProduct struct {
	Name   string  `json:"name"`
	Price  int     `json:"price"`
	Code   *string `json:"code"`
	Detail *string `json:"detail"`
}

type NewSku struct {
	ProductID int     `json:"productId"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Code      *string `json:"code"`
}
