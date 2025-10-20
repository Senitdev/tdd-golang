package transaction

var balance = 0

func Deposite(montant int) int {
	balance += montant
	return balance
}
func Retrait(montant int) int {
	if balance >= montant {
		balance -= montant
	}
	return balance
}
