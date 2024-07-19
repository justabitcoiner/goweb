-- Run: "psql -U postgres -f db/init.sql"
CREATE DATABASE goweb;

\c goweb

CREATE TABLE auth_user (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
	salt TEXT NOT NULL
);