package db

import (
	"math/rand"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/oxisto/track-expenses/model"
	"github.com/sirupsen/logrus"
)

var log *logrus.Entry
var mongo *mgo.Database
var rnd *rand.Rand

func init() {
	// create a new seed
	rnd = rand.New(rand.NewSource(time.Now().Unix()))

	mongoAddr := "localhost:27017"
	log = logrus.WithField("component", "db")

	session, err := mgo.Dial(mongoAddr)
	if err != nil {
		log.Fatal(err)
	}

	mongo = session.DB("track-expenses")
}

// FindExpenses returns an array of all expenses
// TODO: should only return expenses of a particular account, or all accounts someone has access to
func FindExpenses() (expenses []model.Expense, err error) {
	err = mongo.C("expenses").Find(bson.M{}).All(&expenses)

	return expenses, err
}

// InsertExpense inserts one expense into the database
func InsertExpense(expense model.Expense) (err error) {
	return mongo.C("expenses").Insert(&expense)
}

// NextID generates a new random 4-byte ID
func NextID() (ID string) {
	return bson.NewObjectId().Hex()
}
