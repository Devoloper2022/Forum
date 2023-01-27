package repository

import (
	"database/sql"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	userTable = `CREATE TABLE IF NOT EXISTS users(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Username TEXT UNIQUE,
			Email TEXT UNIQUE,
			Password TEXT
		);`
	postTable = `CREATE TABLE IF NOT EXISTS posts(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Title TEXT NOT NULL,
			Text TEXT,
			Date TEXT,
			Like INTEGER,
			Dislike INTEGER,
			UserID INTEGER,
			FOREIGN KEY (UserID) REFERENCES users (ID) ON DELETE CASCADE
		);`
	categoryTable = `CREATE TABLE IF NOT EXISTS categories(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Title TEXT UNIQUE
			);`
	categoryPostTable = `CREATE TABLE IF NOT EXISTS categoriesPost(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			PostID INTEGER,
			CategoryID INTEGER,
			FOREIGN KEY (PostID) REFERENCES posts (ID) ON DELETE CASCADE,
			FOREIGN KEY (CategoryID) REFERENCES categories (ID) ON DELETE CASCADE
			);`
	commentTable = `CREATE TABLE IF NOT EXISTS comments( 
			ID INTEGER PRIMARY  KEY AUTOINCREMENT,
			Text TEXT,
			Date TEXT,
			UserID INTEGER,
			PostID INTEGER,
			FOREIGN KEY (UserID) REFERENCES users (ID) ON DELETE CASCADE,
			FOREIGN KEY (PostID) REFERENCES posts (ID) ON DELETE CASCADE
			);`
	postLikeTable = `CREATE TABLE IF NOT EXISTS postLike(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			UserID INTEGER,
			PostID INTEGER,
			Result INTEGER DEFAULT 0,
			FOREIGN KEY (UserID) REFERENCES users (ID) ON DELETE CASCADE,
			FOREIGN KEY (PostID) REFERENCES posts (ID) ON DELETE CASCADE
			);`
	commentLikeTable = `CREATE TABLE IF NOT EXISTS commentLike(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			UserID INTEGER,
			CommentID INTEGER,
			Result INTEGER DEFAULT 0,
			FOREIGN KEY (UserID) REFERENCES users (ID) ON DELETE CASCADE,
			FOREIGN KEY (CommentID) REFERENCES comments (ID) ON DELETE CASCADE
			);`

	sesionTable = `CREATE TABLE IF NOT EXISTS sessions(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Token TEXT,
			Expiry TEXT,
			UserID INTEGER,
			FOREIGN KEY (UserID) REFERENCES users (ID) ON DELETE CASCADE
			);`
	insertCategories = `INSERT OR IGNORE INTO categories(Title) VALUES
			('Go'),
			('Rust'),
			('JS'),
			('Flutter'),
			('Unity'),
			('Front End'),
			('Backend'),
			('DevOps'),
			('Cyber Security'),
			('Unreal Engine');`
)

func InitDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.Username, cfg.DBName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	allTables := []string{userTable, postTable, commentTable, categoryTable, categoryPostTable, postLikeTable, sesionTable, commentLikeTable, insertCategories}
	for _, table := range allTables {
		_, err := db.Exec(table)
		if err != nil {
			return err
		}
	}

	return nil
}
