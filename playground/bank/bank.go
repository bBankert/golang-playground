package bank

import (
	"fmt"
	"os"
	"strconv"

	"example.com/playground/lib"
)

const balanceFileName string = "Balance.txt"

type Transaction int64

const (
	Exit         Transaction = 0
	CheckBalance Transaction = 1
	Depoit       Transaction = 2
	Withdraw     Transaction = 3
)

func RunBank() {
	var bankInput int
	var selectionResult Transaction
	accountBalance := ReadSavedBalance()

	for {
		DisplayBankOptions()
		lib.GetUserInput("Please select an option:", &bankInput)

		selectionResult = ChooseBankOption(bankInput)

		if selectionResult == Exit {
			PersistBalance(accountBalance)
			fmt.Println("Exiting...")
			break
		}

		switch selectionResult {
		case Depoit:
			accountBalance = HandleDeposit(accountBalance)
		case Withdraw:
			accountBalance = HandleWithdraw(accountBalance)
		case CheckBalance:
			HandleViewBalance(accountBalance)
		}

	}

}

func DisplayBankOptions() {
	fmt.Println("Welcome to the Go Bank!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func ChooseBankOption(choice int) (transaction Transaction) {
	switch choice {
	case 1:
		return CheckBalance
	case 2:
		return Depoit
	case 3:
		return Withdraw
	default:
		return Exit
	}
}

func HandleDeposit(initialBalance float64) (balance float64) {
	var deposit float64
	lib.GetUserInput("Enter a deposit amount:", &deposit)

	if deposit <= 0.0 {
		fmt.Println("Cannot deposit values less than 0")
		balance = initialBalance

		return balance
	}

	balance = initialBalance + deposit

	fmt.Printf("New Balance: %.2f\n", balance)
	return balance
}

func HandleWithdraw(initialBalance float64) (balance float64) {

	var withdrawalAmount float64
	lib.GetUserInput("Enter a withdrawal amount:", &withdrawalAmount)

	if initialBalance <= 0.0 {
		fmt.Println("Cannot withdraw from an empty account")
		balance = initialBalance

		return balance
	} else if withdrawalAmount > initialBalance {
		fmt.Println("Cannot withdraw more than the current balance")
		balance = initialBalance

		return balance
	}

	balance = initialBalance - withdrawalAmount

	fmt.Printf("Remaining Balance: %.2f\n", balance)
	return balance
}

func HandleViewBalance(balance float64) {
	fmt.Printf("Current Balance: %.2f\n", balance)
}

func PersistBalance(currentBalance float64) {
	balanceText := fmt.Sprint(currentBalance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func ReadSavedBalance() (balance float64) {
	if lib.FileExists(balanceFileName) {
		balanceString := lib.ReadFileData(balanceFileName)

		balance, _ = strconv.ParseFloat(balanceString, 64)
	} else {
		balance = 5000
	}

	return balance
}
