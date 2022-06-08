package main

import (
	"fmt"

	"github.com/erseros/holders"
)

func main() {
	holderGuilherme := holders.CurrenteHolder{Holder: "Guilherme", AgencyNumber: 9498, AccountNumber: 1470258, BankBalance: 125.25}
	fmt.Println(holderGuilherme)

	holderBruna := holders.CurrenteHolder{"Bruna", 9498, 1987258, 400.50}
	fmt.Println(holderBruna)

	holderCris := *&holders.CurrenteHolder{}

	holderCris.Holder = "Cris"
	holderCris.AgencyNumber = 6564
	holderCris.AccountNumber = 1245326
	holderCris.BankBalance = 600

	fmt.Println(holderCris)

	holderCris.DrawOut(500)

	fmt.Println(holderCris)

	holderCris.Deposit(240)

	fmt.Println(holderCris)

	fmt.Println("###############################")

	fmt.Println(holderGuilherme)
	fmt.Println(holderCris)

	holderGuilherme.TransferValue(100, &holderCris)

	fmt.Println(holderGuilherme)
	fmt.Println(holderCris)

}
