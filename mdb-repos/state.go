package mdbrepos

import (
	"gopkg.in/mgo.v2"
)

func Init() error {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	
	credentials := mgo.Credential {
		Username: "oddbits",
		Password: "yellowcamelridesbike",
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