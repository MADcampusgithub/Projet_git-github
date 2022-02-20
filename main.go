package main

import (
	m "Projet_git-github/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//router.Handle("POST", "/contact", PostContact)

	err := http.ListenAndServe("0.0.0.0:8080", router)

	if err != nil {
		log.Fatal(err)
	}
}

func PostContact(c *gin.Context) {
	content, err := ioutil.ReadFile("contacts.json")

	if err != nil {
		c.String(http.StatusInternalServerError, "Could not read contacts file")
		log.Fatal(err)
	}

	var contacts []m.Contact
	err = json.Unmarshal(content, &contacts)

	if err != nil {
		c.String(http.StatusInternalServerError, "Could not parse contacts file")
		log.Fatal(err)
	}

	var message m.Contact
	err = c.ShouldBindJSON(&message)

	if err != nil {
		c.String(http.StatusBadRequest, "Data could not be parsed ")
		return
	}

	contacts = append(contacts, message)
	content, err = json.Marshal(contacts)

	if err != nil {
		c.String(http.StatusInternalServerError, "Could not parse contacts")
		log.Fatal(err)
	}

	err = ioutil.WriteFile("contacts.json", content, 0644)

	if err != nil {
		c.String(http.StatusInternalServerError, "Could not write contacts in file")
		log.Fatal(err)
	}

	c.String(http.StatusOK, "Ok")
}
