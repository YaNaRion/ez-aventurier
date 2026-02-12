package infra

import (
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"log"
	"main/infra/models"
	"math/big"
	"os"
)

type Person struct {
	Name  string
	Ordre string
	Unity string
}

// Custom character sets
var (
	LettersOnly  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NumbersOnly  = "0123456789"
	AlphaNumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	URLSafe      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-"
	Readable     = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz" // No confusing chars
)

func CustomID(length int, charset string) (string, error) {
	id := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		id[i] = charset[num.Int64()]
	}

	return string(id), nil
}
func readCSVToStruct(filename string) ([]Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var people []Person

	// Skip header row if exists
	startIndex := 0
	if len(records) > 0 {
		// Optional: check if first row is header
		startIndex = 1
	}

	for i := startIndex; i < len(records); i++ {
		if len(records[i]) >= 3 {
			person := Person{
				Name:  records[i][0],
				Unity: records[i][1],
				Ordre: records[i][2],
			}
			people = append(people, person)
		}
	}

	return people, nil
}

func parseInt(s string) int {
	// Simple conversion - add error handling in production
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}

const path = "infra/Liste-Jeune-Aventurier-avec-equipe.xlsx - Sheet1.csv"

func Add_jeune_to_DB(db *DB) {
	persons, err := readCSVToStruct(path)
	if err != nil {
		log.Println("error while reading csv")
		return
	}

	var users []models.User
	for _, person := range persons {
		userID, err := CustomID(6, AlphaNumeric)
		if err != nil {
			log.Println(err)
			return
		}
		user := models.NewUser(person.Name, userID, person.Unity, person.Ordre)
		users = append(users, *user)
	}

	// log.Println(len(users))

	err = db.AddUsers(users)
	if err != nil {
		log.Println(err)
	}
	log.Println("DATA ADDED SUCESSFULL")
}
