package services

import "math/rand"

func init() {
	plansData := make([]int, 50)
	for i := 0; i < 50; i++ {
		plansData[i] = i + 1
	}
	for _, plan := range plansData {
		println("Plan:", plan)
	}
}

func GeneratePlans() map[int]string {

	//map
	plans := make(map[int]string)
	min, max := 5000, 1000000
	for i := 1; i <= 50; i++ {

		income := rand.Intn(max-min+1) + min // Random income between 0 and 150000
		if income < 10000 {
			plans[i] = "Basic"
		} else if income >= 10000 && income < 50000 {
			plans[i] = "Standard"
		} else if income >= 50000 && income < 100000 {
			plans[i] = "Premium"
		} else {
			plans[i] = "Enterprise"
		}
	}
	return plans
}

func AggregatePayments(payments ...int) int {
	total := 0
	for _, amount := range payments {
		total += amount
	}
	return total
}
