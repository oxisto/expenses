package db

import (
	"math/rand"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/oxisto/expenses/model"
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

// FindExpenses returns an array of all expenses
// TODO: should only return expenses of a particular account, or all accounts someone has access to
func FindExpenses(collection string) (expenses []model.Expense, err error) {
	expenses = []model.Expense{}

	err = mongo.C(collection).Find(bson.M{}).All(&expenses)

	return
}

// Find returns an object given an ID
func Find(ID string, object model.DBObject) (err error) {
	err = mongo.C(object.Collection()).FindId(ID).One(object)

	return
}

// Insert inserts one object into the database
func Insert(object model.DBObject) (err error) {
	return mongo.C(object.Collection()).Insert(&object)
}

// Upsert inserts or updates the given expense
func Upsert(object model.DBObject) (err error) {
	_, err = mongo.C(object.Collection()).Upsert(bson.M{"_id": object.Identifer()}, object)

	return err
}

// NextID generates a new random 4-byte ID
func NextID() (ID string) {
	return bson.NewObjectId().Hex()
}
