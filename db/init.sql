-- Run: "psql -U postgres -f db/init.sql"
CREATE DATABASE goweb;

\c goweb

CREATE TABLE auth_user (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
);

CREATE TABLE article (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	user_id INT NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES auth_user(id)
);