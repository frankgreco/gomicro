package database

import (
    "fmt"
    "strconv"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
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
    // open db
    database, err := sql.Open("mysql", createDBConnString(db.Host, db.Port, db.Name, db.User, db.Pass))
    if err != nil{
        return nil, err
    }
    // make query
    result, err := database.Exec(fmt.Sprintf("INSERT INTO %s (paramOne, paramTwo) VALUES (%s, %s)",
        db.Name + ".<%= nounPluralLower %>",
        getValueOrNull(<%= nounSingularLower %>.ParamOne),
        getValueOrNull(<%= nounSingularLower %>.ParamTwo),
    ))

    if err != nil{
        return nil, err
    }
    // close db
    if err := database.Close(); err != nil {
        return nil, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }
    <%= nounSingularLower %>.Id = id
    return strconv.FormatInt(id, 10), <%= nounSingularLower %>, nil
}

func (db *Database) RetrieveOne(id string) (*models.<%= nounSingularUpper %>, error) {
    database, err := sql.Open("mysql", createDBConnString(db.Host, db.Port, db.Name, db.User, db.Pass))
    if err != nil{
        return nil, err
    }
    var <%= nounSingularLower %> models.<%= nounSingularUpper %>
    var paramOne   sql.NullString
    var paramTwo   sql.NullString
    err = database.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE id='%s'", db.Name + ".<%= nounPluralLower %>", id)).Scan(
        &<%= nounSingularLower %>.Id,
        &paramOne,
        &paramTwo,
    )
    // close db
    if err := database.Close(); err != nil {
        return nil, err
    }
    switch {
    case err == sql.ErrNoRows:
        return nil, nil
    case err != nil:
        return nil, err
    default:
        <%= nounSingularLower %>.ParamOne = paramOne.String
        <%= nounSingularLower %>.ParamTwo = paramTwo.String
        return &<%= nounSingularLower %>, nil
    }
}

func (db *Database) RetrieveAll() (*models.<%= nounPluralUpper %>, error) {
    // open db
    database, err := sql.Open("mysql", createDBConnString(db.Host, db.Port, db.Name, db.User, db.Pass))
    if err != nil{
        return nil, err
    }
    // make query
    rows, err := database.Query(fmt.Sprintf("SELECT * FROM %s", db.Name + ".<%= nounPluralLower %>"))
    if err != nil {
        return nil, err
    }
    <%= nounPluralLower %> := make(models.<%= nounPluralUpper %>, 0)
    for rows.Next() {
        var <%= nounSingularLower %> models.<%= nounSingularUpper %>
        var paramOne   sql.NullString
        var paramTwo   sql.NullString
        err = rows.Scan(
            &<%= nounSingularLower %>.Id,
            &paramOne,
            &paramTwo,
        )
        if err != nil {
            return nil, err
        }
        <%= nounSingularLower %>.ParamOne = paramOne.String
        <%= nounSingularLower %>.ParamTwo = paramTwo.String
        <%= nounPluralLower %> = append(<%= nounPluralLower %>, <%= nounSingularLower %>)
    }
    // close db
    if err := database.Close(); err != nil {
        return nil, err
    }
    return &<%= nounPluralLower %>, nil
}

func (db *Database) UpdateOne(<%= nounSingularLower %> models.<%= nounSingularUpper %>, id string) (*int64, error) {
    database, err := sql.Open("mysql", createDBConnString(db.Host, db.Port, db.Name, db.User, db.Pass))
    if err != nil{
        return nil, err
    }
    // make query
    result, err := database.Exec(fmt.Sprintf("UPDATE %s SET paramOne=%s, paramTwo=%s WHERE id=%s",
        db.Name + ".<%= nounPluralLower %>",
        getValueOrNull(<%= nounSingularLower %>.ParamOne),
        getValueOrNull(<%= nounSingularLower %>.ParamTwo),
        id,
    ));
    if err != nil {
        return nil, err
    }
    rows, err := result.RowsAffected()
    if err != nil {
        return nil, err
    }
    // close db
    if err := database.Close(); err != nil {
        return nil, err
    }
    return &rows, nil
    return nil, nil
}

func (db *Database) DeleteOne(id string) (*int64, error) {
    // open db
    database, err := sql.Open("mysql", createDBConnString(db.Host, db.Port, db.Name, db.User, db.Pass))
    if err != nil{
        return nil, err
    }
    // make query
    result, err := database.Exec(fmt.Sprintf("DELETE FROM %s WHERE id='%s'", db.Name + ".<%= nounPluralLower %>", id));
    if err != nil {
        return nil, err
    }
    rows, err := result.RowsAffected()
    if err != nil {
        return nil, err
    }
    // close db
    if err := database.Close(); err != nil {
        return nil, err
    }
    return &rows, nil
}

func (db *Database) DeleteAll() error {
    // open db
    database, err := sql.Open("mysql", createDBConnString(db.Host, db.Port, db.Name, db.User, db.Pass))
    if err != nil{
        return err
    }
    // make query
    _, err = database.Exec(fmt.Sprintf("TRUNCATE TABLE %s", db.Name + ".<%= nounPluralLower %>"))
    if err != nil {
        return err
    }
    // close db
    if err := database.Close(); err != nil {
        return err
    }
    return nil
}

// [user[:password]@][net[(addr)]]/dbname[?param1=value1&paramN=valueN]
func createDBConnString(server string, port string, db string, user string, pass string) string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, server, port, db)
}

func getValueOrNull(s string) string {
    if s == ""{
        return "NULL"
    }
    return fmt.Sprintf("'%s'", s)
}
