package main

import (
	"fmt"

	"github.com/FreezeOMatic/moyskladHW/employers"
)

func main() {
	// u just should put valid credentials here to make it work
	login, password := "login", "password"

	token, err := employers.GetToken(login, password)
	if err != nil {
		fmt.Println("error while get token: ", err)
		return
	}

	// Create newone
	employee := employers.Employee{
		FirstName: "Дед",
		LastName:  "Напуган",
		//Inn:       "777490425273",
		Phone: "+7(800)623-3743",
	}
	err = employers.CreateEmployee(employee, token)
	if err != nil {
		fmt.Println("error while create employee: ", err)
		return
	}
	// Create another newone
	employee2 := employers.Employee{
		FirstName: "Гор",
		LastName:  "Сенфу",
		//Inn:       "777490425273",
		Phone: "+7(800)769-9769",
	}
	err = employers.CreateEmployee(employee2, token)
	if err != nil {
		fmt.Println("error while create employee: ", err)
		return
	}

	// Get & print list of employers
	list, err := employers.GetEmpList(token)
	if err != nil {
		fmt.Println("error while get employers list: ", err)
		return
	}

	fmt.Println("We have ", len(list.Rows), " employers.")

	if len(list.Rows) > 0 {
		for num, employ := range list.Rows {
			fmt.Println("________________________________________________________________________________")
			fmt.Println("номер сотрудника", num+1)
			fmt.Printf(
				"Имя %s, Фамилия %s, связь по номеру %s, \r\n metadata %s \r\n metadata %s \r\n metadata %s \r\n metadata %s \r\n",
				employ.FullName,
				employ.LastName,
				employ.Phone,
				employ.Meta.Href,
				employ.Meta.MetadataHref,
				employ.Meta.MediaType,
				employ.Meta.Type,
			)
		}
	}

	// Change some info in first employee (if exist)
	if len(list.Rows) > 0 {
		err = employers.ChangeEmployee(token, list.Rows[0].GetID(), "Имярек", "Инопланетный")
		if err != nil {
			fmt.Println("error while change employee: ", err)
			return
		}
	}

	// Get & print list of employers after change
	list, err = employers.GetEmpList(token)
	if err != nil {
		fmt.Println("error while get employers list: ", err)
		return
	}

	fmt.Println("We have ", len(list.Rows), " employers.")

	if len(list.Rows) > 0 {
		for num, employ := range list.Rows {
			fmt.Println("________________________________________________________________________________")
			fmt.Println("номер сотрудника", num+1)
			fmt.Printf(
				"Имя %s, Фамилия %s, связь по номеру %s, \r\n metadata %s \r\n metadata %s \r\n metadata %s \r\n metadata %s \r\n",
				employ.FullName,
				employ.LastName,
				employ.Phone,
				employ.Meta.Href,
				employ.Meta.MetadataHref,
				employ.Meta.MediaType,
				employ.Meta.Type,
			)
		}
	}
	if len(list.Rows) > 0 {
		err = employers.DeleteEmployee(token, list.Rows[0].GetID())
		if err != nil {
			fmt.Println("error while delete employee: ", err)
			return
		}
	}

	fmt.Println("Thats all...")
}
