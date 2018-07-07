/*
Copyright 2018 Christian Banse

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package db

const UsersCollectionName = "users"

type User struct {
	ID           string       `json:"id" bson:"_id"`
	Username     string       `json:"username" bson:"username"`
	PasswordHash string       `json:"passwordHash,omitempty" bson:"passwordHash"`
	Delegations  []Delegation `json:"delegations,omitempty" bson:"delegations"`
}

type Delegation struct {
	AccountID   string `json:"accountId" bson:"accountId"`
	AccountName string `json:"accountName" bson:"accountName"`
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

func (u User) CanAccess(expense Expense) bool {
	userIDs := FindDelegatedAccounts(u)

	_, ok := userIDs[expense.AccountID]

	return ok
}
