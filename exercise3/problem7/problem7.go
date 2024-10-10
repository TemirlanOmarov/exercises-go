package problem7

import (
	"fmt"
)

type BankAccount struct {
	name    string
	balance float64
}

type FedexAccount struct {
	name     string
	packages []string
}

type KazPostAccount struct {
	name     string
	balance  float64
	packages []string
}

func withdrawMoney(amount float64, accounts ...interface{}) {
	for _, account := range accounts {
		switch a := account.(type) {
		case *BankAccount:
			if a.balance >= amount {
				a.balance -= amount
				fmt.Printf("%s withdrew %.2f. New balance: %.2f\n", a.name, amount, a.balance)
			} else {
				fmt.Printf("%s has insufficient funds to withdraw %.2f.\n", a.name, amount)
			}
		case *KazPostAccount:
			if a.balance >= amount {
				a.balance -= amount
				fmt.Printf("%s withdrew %.2f. New balance: %.2f\n", a.name, amount, a.balance)
			} else {
				fmt.Printf("%s has insufficient funds to withdraw %.2f.\n", a.name, amount)
			}
		default:
			fmt.Printf("Account type not supported for withdrawal.\n")
		}
	}
}

func sendPackagesTo(recipient string, accounts ...interface{}) {
	for _, account := range accounts {
		switch a := account.(type) {
		case *FedexAccount:
			if len(a.packages) > 0 {
				for _, pkg := range a.packages {
					fmt.Printf("%s sent package '%s' to %s via FedEx.\n", a.name, pkg, recipient)
				}
				a.packages = nil
			} else {
				fmt.Printf("%s has no packages to send.\n", a.name)
			}
		case *KazPostAccount:
			if len(a.packages) > 0 {
				for _, pkg := range a.packages {
					fmt.Printf("%s sent package '%s' to %s via KazPost.\n", a.name, pkg, recipient)
				}
				a.packages = nil
			} else {
				fmt.Printf("%s has no packages to send.\n", a.name)
			}
		default:
			fmt.Printf("Account type not supported for sending packages.\n")
		}
	}
}

func main() {
	normanOsborne := &BankAccount{
		name:    "Norman Osborne",
		balance: 1_000_000,
	}
	peterParker := &FedexAccount{
		name:     "Peter Parker",
		packages: []string{"Web Shooter", "Spider Suit"},
	}
	auntMay := &KazPostAccount{
		name:     "Aunt May",
		balance:  200,
		packages: []string{"Cookies", "A Letter"},
	}

	withdrawMoney(10, normanOsborne, auntMay)
	sendPackagesTo("Mary Jane", peterParker, auntMay)
}
