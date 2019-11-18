package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	credentials := map[string]string{
		"0a041b9462caa4a31bac3567e0b6e6fd9100787db2ab433d96f6d178cabfce90": "0b14d501a594442a01c6859541bcb3e8164d183d32937b851835442f69d5c94e",
		"user2": "password2",
	}

	fmt.Print("Username: ")
	usertoken, _ := terminal.ReadPassword(0)
	hasher := sha256.New()
	io.WriteString(hasher, string(usertoken))
	userhash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(userhash)

	fmt.Print("Password: ")
	passtoken, _ := terminal.ReadPassword(0)
	hasher = sha256.New()
	io.WriteString(hasher, string(passtoken))
	passhash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(passhash)

	fmt.Println()
	if credentials[userhash] == passhash {
		fmt.Println("Login succesful.")
	} else {
		fmt.Println("Invalid username or password.")
	}
}

//func encoder([]byte token) string {
//
//	token, err := terminal.ReadPassword(0)
//	hasher := sha256.New()
//	io.WriteString(hasher, string(token))
//	hash := hex.EncodeToString(hasher.Sum(nil))
//	return hash
//
//}
