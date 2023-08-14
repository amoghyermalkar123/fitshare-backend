package types

type UserDetails struct {
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Height   string `json:"height" bson:"height"`
	Weight   string `json:"weight" bson:"weight"`
	Age      string `json:"age" bson:"age"`
	UserType string `json:"user_type" bson:"user_type"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetailsResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Height   string `json:"height"`
	Weight   string `json:"weight"`
	Age      string `json:"age"`
	UserType string `json:"user_type"`
}
