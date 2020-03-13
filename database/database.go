package database

import (
	"log"

	"../config"
	"github.com/globalsign/mgo"
)

// UrlsCollection - urls table at mongoDB
var UrlsCollection *mgo.Collection

// InitDB - initialises mongoDB with your configurations at config/config.go.
func InitDB() {
	session, err := mgo.Dial(config.DataBaseHost)
	if err != nil {
		log.Fatal(err)
	}

	UrlsCollection = session.DB("url-shortener").C("urls")

	if err = session.Ping(); err != nil {
		log.Fatal(err)
	}
}