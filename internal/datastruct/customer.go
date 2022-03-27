package datastruct

type Customer struct {
	Name     string `bson:"name"`
	Safe     int64  `bson:"safe"`
	Accounts int64  `bson:"accounts"`
}
