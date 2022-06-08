package holders

import "fmt"

type currenteHolder struct {
	holder        string
	agencyNumber  int
	accountNumber int
	bankBalance   float64
}

func (holder *currenteHolder) drawOut(value float64) bool {
	var success bool = true

	if value > 0 {
		if holder.bankBalance < value {
			success = false
			fmt.Println("Insufficient funds")
		} else {
			holder.bankBalance -= value
			fmt.Println("Sucessful drawout")
		}
	} else {
		fmt.Println("Enter a value greater than zero")
	}

	return success
}

func (holder *currenteHolder) deposit(value float64) {
	holder.bankBalance += value

	fmt.Println("Successful deposit")
}

func (holderOwner *currenteHolder) transferValue(value float64, holder *currenteHolder) {
	if holderOwner.bankBalance < value {
		fmt.Println("Insufficient funds")
	} else {
		holderOwner.bankBalance -= value
		holder.deposit(value)
	}
}
