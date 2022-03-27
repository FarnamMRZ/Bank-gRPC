package datastruct

type Account struct {
	Bank     string `bson:"bank"`
	Customer string `bson:"customer"`
	Number   string `bson:"number"`
	Amount   int64  `bson:"amount"`
}
