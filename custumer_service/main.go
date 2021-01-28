package main

import (
	"fmt"
	"go-templ-example-master/custumer_service/model"
	"go-templ-example-master/custumer_service/service"
)

type customerView struct {
	key string
	loop bool

	customerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.customerService.List()

	fmt.Println("-------------Customers list-----------------")
	fmt.Println("id\tName\tGender\tAge\tPhone\tEmail\t")
	for i := 0; i < len(customers); i ++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("----------Customers list Finished-----------")
	fmt.Println()
}


func (this *customerView) add() {
	fmt.Println("-------------Add Customer--------------------")
	fmt.Print("Name: ")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("Gender: ")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("Age: ")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("Phone: ")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("Email: ")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name,age,gender,phone,email)

	if this.customerService.Add(customer) {
		fmt.Println("------------Finished adding----------------")
		fmt.Println()
	} else {
		fmt.Print("--------------Failed to add --------------------")
		fmt.Println()
	}

}

func (this *customerView) delete() {
	fmt.Println("------------delete Customer-------------")
	fmt.Print("Please select the Id you need to delete (tap -1 to exit): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return 
	}
	fmt.Print("Do you want to Delete (Y/N)?: ")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("-------------------Successful delete-------------------")
		} else {
			fmt.Println("---------Failed to Delete. The Id does not exit--------")
		}
	}
}

func (this *customerView) exit() {
	fmt.Println("Do you want to Exit (Y/N)?: ")
	for {
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key == "Y" || this.key == "n" || this.key == "N" {
			break
		}

		fmt.Print("There are error, Do you really want to Exit (Y/N)?: ")
	}

	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}


func (this *customerView) mainMenu() {
	for {
		fmt.Println("----------------Customer information systems-----------------")
		fmt.Println("		1 add customer                               ")
		fmt.Println("		2 change customer                            ")
		fmt.Println("		3 delete customer                            ")
		fmt.Println("		4 customer list                              ")
		fmt.Println("		5 Exit                                       ")
		fmt.Print("Please select (1-5): ")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("change customer")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default :
			fmt.Println("Input not correct, Please try again...")
		}

		if !this.loop {
			break
		}
	}

	fmt.Println("Already Exit system...")
}


func main()  {
	customerView := customerView {
		key : "",
		loop : true,
	}

	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()
}