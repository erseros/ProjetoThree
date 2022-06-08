package main

import (
	"fmt"
)

func main() {
	holderGuilherme := holders.currenteHolder{holder: "Guilherme", agencyNumber: 9498, accountNumber: 1470258, bankBalance: 125.25}
	fmt.Println(holderGuilherme)

	holderBruna := holders.currenteHolder{"Bruna", 9498, 1987258, 400.50}
	fmt.Println(holderBruna)

	holderCris := *&holders.currenteHolder{}

	holderCris.holder = "Cris"
	holderCris.agencyNumber = 6564
	holderCris.accountNumber = 1245326
	holderCris.bankBalance = 600

	fmt.Println(holderCris)

	holderCris.drawOut(500)

	fmt.Println(holderCris)

	holderCris.deposit(240)

	fmt.Println(holderCris)

	fmt.Println("###############################")

	fmt.Println(holderGuilherme)
	fmt.Println(holderCris)

	holderGuilherme.transferValue(100, &holderCris)

	fmt.Println(holderGuilherme)
	fmt.Println(holderCris)

}
