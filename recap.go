package main


var balance = 2000

func main() {
	println(checkBalance())
	depositMoney(1000)
	println(checkBalance())
	withdrawMoney(100)
}

func checkBalance() int{
	return balance
}

func depositMoney(amount int){
	balance = balance + amount
}

func withdrawMoney(amount int){
	if amount <= balance{
		balance = balance - amount
	}else{
		print("You have insufficient funds")
	}
}
