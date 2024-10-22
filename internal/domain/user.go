package domain

type User struct {
	ID       string `dynamo:"id" json:"id"`
	UserName string `dynamo:"username" json:"username"`
	Name     string `dynamo:"name" json:"name"`
	Password string `dynamo:"password" json:"-"`
}
