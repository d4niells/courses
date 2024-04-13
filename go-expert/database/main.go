package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float32
}

func NewProduct(name string, price float32) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	p := NewProduct("Lamp", 40.0)

	err = insertProduct(db, p)
	if err != nil {
		panic(err)
	}

	// p.Price = 123.43
	//
	// err = updateProduct(db, p)
	// if err != nil {
	// 	panic(err)
	// }

	// seachedProduct, err := getProduct(db, "5eecccca-9952-4c91-9f78-b01dc798d149")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("The product %s with price %.2f was found in our database!\n", seachedProduct.Name, seachedProduct.Price)

	// products, err := getAllProducts(db)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for _, v := range products {
	// 	fmt.Printf("Name: %s and Price: %.2f\n", v.Name, v.Price)
	// }

	deleteProduct(db, "8f1344e4-a003-4567-a39d-1f76faeda3e3")
	fmt.Println("Product was deleted permanently!")
}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}

	return nil
}

func getProduct(db *sql.DB, productID string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(productID).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func getAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product

		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
