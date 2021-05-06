package main

import "errors"

const (
	goodScore      = 450
	lowScoreRatio  = 10
	goodScoreRatio = 20
)

var (
	ErrCreditScore = errors.New("invalid credit score")
	ErrIncome      = errors.New("income invalid")
	ErrLoanAmount  = errors.New("loan amount invalid")
	ErrLoanTerm    = errors.New("loan term not a multiple of 12")
)

func main() {

}

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	interest := 20.0

	if creditScore >= goodScore {
		interest = 15.0
	}

	if creditScore < 1 {
		return ErrCreditScore
	}

	if income < 1 {
		return ErrIncome
	}

	if loanAmount < 1 {
		return ErrLoanAmount
	}

	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}

	rate := interest / 100

	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)

	print(payment)
	// totalInterest := (payment * loanTerm) - loanAmount

	return nil
}
