package main

import (
	"bank/accounts"
	"bank/holders"
	"fmt"
)

type checkAccount interface {
	Withdraw(withdrawalAmount float64) (string, float64)
}

func Payment(account checkAccount, value float64) {
	account.Withdraw(value)
}

func main() {
	fooHolder := holders.Holder{
		Name:     "Foobar",
		Document: "123.456.789.10",
		Job:      "Software Engineer",
	}

	fooSavings := accounts.SavingsAccount{
		AccountHolder: fooHolder,
		AccountCode:   123,
		AccountNumber: 12345,
	}

	fooChecking := accounts.CheckingAccount{
		AccountHolder: fooHolder,
		AccountCode:   fooSavings.AccountCode,
		AccountNumber: fooSavings.AccountNumber,
	}

	fooSavings.Deposit(500)
	fooChecking.Deposit(250)

	Payment(&fooSavings, 50)
	Payment(&fooChecking, 50)

	fmt.Println(fooSavings.GetBalance())
	fmt.Println(fooChecking.GetBalance())
}
