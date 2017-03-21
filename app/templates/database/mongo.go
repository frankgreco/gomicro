package database

import (
    "fmt"
    "time"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "<%= vcs %>/<%= repo %>/<%= project %>/models"
)

type Database struct {
    Host string
    Port string
    User string
    Pass string
    Name string
}

func (db *Database) Create(<%= nounSingularLower %> *models.<%= nounSingularUpper %>) (string, *models.<%= nounSingularUpper %>, error) {
    mongoSession, err := createMongoSession(db)
	if err != nil {
		return nil, err
	}
    mongoSession.SetMode(mgo.Monotonic, true)
    id := bson.NewObjectId()
    <%= nounSingularLower %>.Id = id
    sessionCopy := mongoSession.Copy()
    defer sessionCopy.Close()
    collection := sessionCopy.DB(db.Name).C("people")
    _ = collection.Insert(<%= nounSingularLower %>)
    return id.String()[13:len(id.String())-2], <%= nounSingularLower %>, nil
}

func (db *Database) RetrieveOne(id string) (*models.<%= nounSingularUpper %>, error) {
    return nil, nil
}

func (db *Database) RetrieveAll() (*models.<%= nounPluralUpper %>, error) {
    mongoSession, err := createMongoSession(db)
	if err != nil {
		return nil, err
	}
    mongoSession.SetMode(mgo.Monotonic, true)
    sessionCopy := mongoSession.Copy()
    defer sessionCopy.Close()
    collection := sessionCopy.DB(db.Name).C("people")
    var <%= nounPluralLower %> models.<%= nounPluralUpper %>
	err = collection.Find(nil).All(&<%= nounPluralLower %>)
    if err != nil {
        return nil, err
    }
    return &<%= nounPluralLower %>, nil
}

func (db *Database) UpdateOne(<%= nounSingularLower %> models.<%= nounSingularUpper %>, id string) (*int64, error) {
    return nil, nil
}

func (db *Database) DeleteOne(id string) (*int64, error) {
    return nil, nil
}

func (db *Database) DeleteAll() error {
    return nil
}

func createMongoSession(db *Database) (*mgo.Session, error) {
    return mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%s", db.Host, db.Port)},
		Timeout:  60 * time.Second,
		//Database: "admin",
		//Username: db.User,
		//Password: db.Pass,
	})
}
