/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUser(username string) (User, error)
	SetUser(Username string, identifier int) error
	SetUsername(Username string, newUsername string) (int, error)
	SetPhoto(Username string, identifier uint64, file string) error
	SetBan(Username string, token int, banIdentifier int) error
	RemoveBan(banIdentifier int) error
	GetBans(Token int) ([]Ban, error)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		usersDatabase := `CREATE TABLE users (
			id INTEGER PRIMARY KEY, 
			username TEXT);`
		photosDatabase := `CREATE TABLE photos (
				id INTEGER PRIMARY KEY, 
				photos TEXT);`
		likesDatabase := `CREATE TABLE likes (
			id INTEGER PRIMARY KEY, 
			likes TEXT);`
		commentsDatabase := `CREATE TABLE comments(
			id INTEGER PRIMARY KEY, 
			comments TEXT);`
		bansDatabase := `CREATE TABLE bans (
			id INTEGER PRIMARY KEY, 
			username TEXT,
			token INTEGER);`
		followersDatabase := `CREATE TABLE followers (
				id INTEGER PRIMARY KEY, 
				followers TEXT);`
		_, err = db.Exec(usersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(photosDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		_, err = db.Exec(likesDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		_, err = db.Exec(commentsDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		_, err = db.Exec(bansDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		_, err = db.Exec(followersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
