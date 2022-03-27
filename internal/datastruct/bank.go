package datastruct

type Bank struct {
	Name     string    `bson:"name"`
	Accounts []Account `bson:"accounts"`
}
