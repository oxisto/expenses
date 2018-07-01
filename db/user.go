package db

const UsersCollectionName = "users"

type User struct {
	ID           string `json:"id" bson:"_id"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"passwordHash,omitempty" bson: "passwordHash"`
}

func (u User) Collection() string {
	return UsersCollectionName
}

func (u User) Identifier() string {
	return u.ID
}

// NewUser creates a new user with default values
func NewUser(username string) User {
	return User{
		ID:       NextID(),
		Username: username,
	}
}
