package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/oxisto/expenses/db"

	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

var log *logrus.Entry

func init() {
	// Set log level to debug
	logrus.SetLevel(logrus.DebugLevel)

	log = logrus.WithField("component", "main")
}

func main() {
	var (
		username string
		password string
		hash     []byte
		err      error
	)

	if username, password, err = GetCredentials(); err != nil {
		fmt.Errorf("An error occured %s", err)
		return
	}

	if hash, err = bcrypt.GenerateFromPassword([]byte(password), 10); err != nil {
		fmt.Errorf("An error occured %s", err)
		return
	}

	// check, if user already exists, we don't want to change the IDs
	var user db.User
	err = db.Find(bson.M{"username": username}, &user)
	if err == db.ErrNotFound {
		user = db.NewUser(username)
		log.Infof("Creating new user %s.", user.Username)
	} else {
		log.Infof("User %s already exists, using ID %s to modify entry.", user.Username, user.ID)
	}
	user.PasswordHash = string(hash)

	db.Upsert(user)
}

func GetCredentials() (username string, password string, err error) {
	var passwordBytes []byte

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err = reader.ReadString('\n')

	fmt.Print("Enter Password: ")
	passwordBytes, err = terminal.ReadPassword(int(syscall.Stdin))

	password = string(passwordBytes)

	fmt.Println()

	return strings.TrimSpace(username), strings.TrimSpace(password), err
}
