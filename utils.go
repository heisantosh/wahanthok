package wahanthok

import (
	"os"
	"log"
	"io/ioutil"
	"encoding/gob"
)

// Load data into a Was struct. Data is stored on disk in gob format.
// Return a pointer to a Was struct.
func LoadDB() (*Was, error) {
	db := make(Was)

	f, err := os.Open("db/db.gob")
	if err != nil {
		log.Println("Failed to open db/db.gob")
		return nil, err
	}

	dec := gob.NewDecoder(f)
	err = dec.Decode(&db)
	if err != nil {
		log.Println("Failed to decode gob read from db/db.gob")
		return nil, err
	}

	return &db, nil
}

// Save a Was struct on disk in gob format.
func SaveDB(db *Was) error {
	return nil
}

// Return access token from the config file.
func GetAccessToken(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Println("Failed to open", file)
		return "", err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("Failed to read data")
		return "", err
	}
	return string(data), nil
}

// Return a random image URL from a collection of images.
func GetAnImageURL(tag string) string {
	return ""
}