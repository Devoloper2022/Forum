CREATE TABLE User (
	ID numeric PRIMARY KEY AUTOINCREMENT,
	username varchar,
	password text,
	email text
);

CREATE TABLE Post (
	ID numeric PRIMARY KEY AUTOINCREMENT,
	title varchar,
	text text,
	data datetime,
	UserID numeric
);

CREATE TABLE Category (
	ID numeric PRIMARY KEY AUTOINCREMENT,
	title varchar
);

CREATE TABLE Commet (
	ID numeric PRIMARY KEY AUTOINCREMENT,
	text text,
	data datetime,
	UserID numeric,
	PostID numeric
);

