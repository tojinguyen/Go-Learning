package main

func main() {
	a := &Account{Balance: 10}
	b := Account{Balance: 20}

	a.Deposit(100)
	b.Deposit(200)

	println("Account A Balance:", a.GetBalance())
	println("Account B Balance:", b.GetBalance())

	Test()
}
