package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadData() {

	var medicineName string
	var companyName string
	var quantity int
	var price float64
	var expiryDate string
	//var prescription string
	//take input from user

	println("Enter Medicine Details:")
	_, err := fmt.Scan(&medicineName, &companyName, &quantity, &price, &expiryDate)
	if err != nil {
		fmt.Println("Error reading medicine name:", err)
		return
	}

	println("Enter Prescription Details:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	//print the input data
	fmt.Printf("Medicine Name: %s\n", medicineName)
	fmt.Printf("Company Name: %s\n", companyName)
	fmt.Printf("Quantity: %d\n", quantity)
	fmt.Printf("Price: %.2f\n", price)
	fmt.Printf("Expiry Date: %s\n", expiryDate)
	fmt.Printf("Prescription: %s\n", text)
}
