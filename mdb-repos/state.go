package mdbrepos

import (
	"os"
	"fmt"
	"gopkg.in/mgo.v2"
)

func Init() error {

	host := os.Getenv("MONGO_HOST")
	if (host == "") {
		panic("MONGO_HOST is not defined")
	}
	port := os.Getenv("MONGO_PORT")
	if (port == "") {
		panic("MONGO_PORT is not defined")
	}
	user := os.Getenv("MONGO_USERNAME")
	if (user == "") {
		panic("MONGO_USERNAME is not defined")
	}
	password := os.Getenv("MONGO_PASSWORD")
	if (password == "") {
		panic("MONGO_PASSWORD is not defined")
	}

	session, err := mgo.Dial(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	
	credentials := mgo.Credential {
		Username: user,
		Password: password,
	}
	if err = session.Login(&credentials); err != nil {
		return err
	}

	db.masterSession = session
	
	return nil
}

func TearDown() {
	if db.masterSession != nil {
		db.masterSession.Close()
	}
}

var db = dbState { }

type dbState struct { 
	masterSession *mgo.Session
}

func (this *dbState) getSession() *mgo.Session {
	return this.masterSession.Copy()
}

func (this *dbState) initDb() error {
	session := this.getSession()
	defer session.Close()
	session.SetMode(mgo.Strong, false)
	collection := session.DB("oddbits").C("organizations")

	index := mgo.Index {
		Key:        []string{ "code"},
		Unique:     true,
	}
	
	if err := collection.EnsureIndex(index); err != nil {
		return err
	}
	return nil;
}