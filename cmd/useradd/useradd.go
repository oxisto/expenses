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

	if useDelegation, delegation, err := GetDelegation(); useDelegation {
		if err == nil {
			user.Delegations = append(user.Delegations, delegation)

			log.Infof("Allowing username %s (%s) to access your expenses.", delegation.AccountName, delegation.AccountID)
		} else {
			log.Errorf("Could not add delegation: %s", err)
		}
	}

	db.Upsert(user)
}

func GetDelegation() (useDelegation bool, delegation db.Delegation, err error) {
	var (
		tmp  string
		user db.User
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Do you want to add delegation? ")
	tmp, err = reader.ReadString('\n')

	if strings.TrimSpace(tmp) == "y" {
		useDelegation = true
	}

	if useDelegation {
		fmt.Print("Enter Username: ")
		tmp, err = reader.ReadString('\n')
		delegation.AccountName = strings.TrimSpace(tmp)

		err = db.Find(bson.M{"username": delegation.AccountName}, &user)
		delegation.AccountID = user.ID
	}

	return
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
