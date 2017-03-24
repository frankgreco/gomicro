package database

import (
    "fmt"
    "errors"
    "strconv"

    "github.com/jinzhu/gorm"
    "<%= vcs %>/<%= repo %>/<%= project %>/models"
    _ "github.com/jinzhu/gorm/dialects/<%= db %>"
)

type Database struct {
    <%if (db != "sqlite") { %>
    Host string
    Port string
    User string
    Pass string
    <% } %>
    Name string
    <%if (db == "sqlite") { %>
    Location string
    <% } %>
}

// Attemp to open a new database connection
func (db *Database) Open() (*gorm.DB, error){
    conn, err := gorm.Open("<%= db == "sqlite" ? "sqlite3" : db %>", db.createConnString())
    if err != nil {
        return nil, err
    }
    err = conn.DB().Ping()
    if err != nil {
        return nil, err
    }
    conn.LogMode(false)
    conn.AutoMigrate(&models.<%= nounSingularUpper %>{})
    return conn, nil
}

func (db *Database) Ping() bool {
    conn, err := db.Open()
    if err != nil {
        return false
    }
    defer conn.Close()
    return true
}

func (db *Database) createConnString() string {
    <%if (db == "mysql") { %>
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        db.User,
        db.Pass,
        db.Host,
        db.Port,
        db.Name,
    )
    <% } %>
    <%if (db == "postgres") { %>
    return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
        db.Host,
        db.Port,
        db.User,
        db.Name,
        db.Pass,
    )
    <% } %>
    <%if (db == "sqlite") { %>
    return fmt.Sprintf("%s",
        db.Location,
    )
    <% } %>
}

func (db *Database) Create(<%= nounSingularLower %> *models.<%= nounSingularUpper %>) (*models.<%= nounSingularUpper %>, error) {
    conn, err := db.Open()
    if err != nil {
        return nil, err
    }
    defer conn.Close()
    result := conn.Create(&<%= nounSingularLower %>)
    if result.Error != nil {
        return nil, result.Error
    } else if result.RowsAffected < 1 {
        return nil, errors.New("record not inserted")
    }
    return <%= nounSingularLower %>, nil
}

func (db *Database) RetrieveOne(id string) (*models.<%= nounSingularUpper %>, error) {
    conn, err := db.Open()
    if err != nil {
        return nil, err
    }
    defer conn.Close()
    ID, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return nil, err
    }
    <%= nounSingularLower %> := models.<%= nounSingularUpper %>{}
    result := conn.First(&<%= nounSingularLower %>, ID)
    if result.Error != nil {
        return nil, result.Error
    } else if result.RowsAffected < 1 {
        return nil, nil
    }
    return &<%= nounSingularLower %>, nil
}

func (db *Database) RetrieveAll() (*models.<%= nounPluralUpper %>, error) {
    conn, err := db.Open()
    if err != nil {
        return nil, err
    }
    defer conn.Close()
    <%= nounPluralLower %> := models.<%= nounPluralUpper %>{}
    result := conn.Find(&<%= nounPluralLower %>)
    if result.Error != nil {
        return nil, result.Error
    } else if result.RowsAffected < 1 {
        return nil, nil
    }
    return &<%= nounPluralLower %>, nil
}

func (db *Database) UpdateOne(update models.<%= nounSingularUpper %>, id string) (*models.<%= nounSingularUpper %>, error) {
    conn, err := db.Open()
    if err != nil {
        return nil, err
    }
    defer conn.Close()
    ID, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        return nil, err
    }
    <%= nounSingularLower %> := models.<%= nounSingularUpper %>{}
    result := conn.First(&<%= nounSingularLower %>, ID)
    if result.Error != nil {
        return nil, result.Error
    }
    if result.RowsAffected < 1 {
        return nil, nil
    }
    update.ID = ID
    result = conn.Save(&update)
    if result.Error != nil {
        return nil, result.Error
    }
    return &update, nil
}

func (db *Database) DeleteOne(id string) (*models.<%= nounSingularUpper %>, error) {
    conn, err := db.Open()
    if err != nil {
        return nil, err
    }
    defer conn.Close()
    ID, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        return nil, err
    }
    <%= nounSingularLower %> := models.<%= nounSingularUpper %>{ID: ID}
    result := conn.Delete(&<%= nounSingularLower %>)
    if result.Error != nil {
        return nil, result.Error
    } else if result.RowsAffected < 1 {
        return nil, nil
    }
    return &<%= nounSingularLower %>, nil
}

func (db *Database) DeleteAll() (*models.<%= nounPluralUpper %>, error) {
    conn, err := db.Open()
    if err != nil {
        return nil, err
    }
    defer conn.Close()
    <%= nounPluralLower %> := models.<%= nounPluralUpper %>{}
    result := conn.Delete(&<%= nounPluralLower %>)
    if result.Error != nil {
        return nil, result.Error
    } else if result.RowsAffected < 1 {
        return nil, nil
    }
    return &<%= nounPluralLower %>, nil
}
