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

import (
	"math/rand"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
)

var (
	log   *logrus.Entry
	mongo *mgo.Database
	rnd   *rand.Rand
)

// ErrNotFound indicates that an object was not found in the database
var ErrNotFound = mgo.ErrNotFound

func init() {
	// create a new seed
	rnd = rand.New(rand.NewSource(time.Now().Unix()))

	mongoAddr := "localhost:27017"
	log = logrus.WithField("component", "db")

	log.Infof("Connecting to MongoDB @ %s...", mongoAddr)

	session, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatal(err)
	}

	mongo = session.DB("expenses")
}

func GetUserIDsWithDelegation(user User) (userIDs []string) {
	// add the own user ID
	userIDs = append(userIDs, user.ID)

	// get the list of IDs this user has access to
	delegatedIDs, err := FindDistinctUsers(bson.M{"delegations.accountId": user.ID})

	if err == nil || err == ErrNotFound {
		// add the delegated IDs
		userIDs = append(userIDs, delegatedIDs...)
	}

	return userIDs
}

// FindExpenses returns an array of all expenses
func FindExpenses(user User, collection string) (expenses []Expense, err error) {
	expenses = []Expense{}

	userIDs := GetUserIDsWithDelegation(user)

	// query all expenses that belong to those user IDs
	err = mongo.C(collection).Find(bson.M{"accountID": bson.M{"$in": userIDs}}).Sort("-timestamp").All(&expenses)

	return
}

func FindDistinctUsers(filter bson.M) (userIDs []string, err error) {
	err = mongo.C("users").Find(filter).Distinct("_id", &userIDs)

	return
}

func FindDelegatedAccounts(user User) (users map[string]User) {
	var delegate User

	users = make(map[string]User)

	// add the user itself
	users[user.ID] = user

	query := mongo.C("users").Find(bson.M{"delegations.accountId": user.ID}).Select(bson.M{"_id": 1, "username": 1}).Iter()

	for !query.Done() {
		query.Next(&delegate)

		if query.Err() == nil {
			users[delegate.ID] = delegate
		}
	}

	return
}

func FindExpense(user User, ID string) (expense Expense, err error) {
	userIDs := GetUserIDsWithDelegation(user)

	err = mongo.C(expense.Collection()).Find(bson.M{"_id": ID, "accountID": bson.M{"$in": userIDs}}).One(&expense)

	return
}

// Find returns an object given an ID
func FindID(ID string, object DBObject) (err error) {
	err = mongo.C(object.Collection()).FindId(ID).One(object)

	return
}

// Find returns an object given a query
func Find(query bson.M, object DBObject) (err error) {
	err = mongo.C(object.Collection()).Find(query).One(object)

	return
}

// Insert inserts one object into the database
func Insert(object DBObject) (err error) {
	return mongo.C(object.Collection()).Insert(&object)
}

// Upsert inserts or updates the given expense
func Upsert(object DBObject) (err error) {
	_, err = mongo.C(object.Collection()).Upsert(bson.M{"_id": object.Identifier()}, object)

	return err
}

// NextID generates a new random 4-byte ID
func NextID() (ID string) {
	return bson.NewObjectId().Hex()
}
