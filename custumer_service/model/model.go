package model

import "fmt"

type Customer struct {
	Id int
	Name string
	Age int
	Gender string
	Phone  string
	Email  string
}

func NewCustomer(id int,name string,age int,gender string, phone string, email string) Customer {
	return Customer{
		Id : id,
		Name : name,
		Age : age,
		Gender : gender,
		Phone : phone,
		Email : email,
	}
}

func NewCustomer2(name string,age int,gender string, phone string, email string) Customer {
	return Customer{
		Name : name,
		Age : age,
		Gender : gender,
		Phone : phone,
		Email : email,
	}
}

func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t %v\t %v\t %v\t %v\t %v\t",this.Id,
			this.Name,this.Age,this.Gender,this.Phone,this.Email)

	return info
}