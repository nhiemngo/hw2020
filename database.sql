USE mysql;
CREATE DATABASE IF NOT EXISTS sellerdb;
USE sellerdb;

CREATE USER IF NOT EXISTS 'user'@'localhost';
GRANT SELECT, INSERT, UPDATE, CREATE, ALTER ON sellerdb.* TO 'user'@'localhost';

ALTER USER 'user'@'localhost' IDENTIFIED BY 'password';

DROP TABLE IF EXISTS schedule;
DROP TABLE IF EXISTS seller;
CREATE TABLE IF NOT EXISTS  seller(
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
name VARCHAR(30) NOT NULL,
logo TEXT,
image TEXT,
phone VARCHAR(20),
location VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS schedule (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
monday BOOLEAN NOT NULL,
tuesday BOOLEAN NOT NULL,
wednesday BOOLEAN NOT NULL,
thursday BOOLEAN NOT NULL,
friday BOOLEAN NOT NULL,
saturday BOOLEAN NOT NULL,
sunday BOOLEAN NOT NULL,
monday_start VARCHAR(5),
tuesday_start VARCHAR(5),
wednesday_start VARCHAR(5),
thursday_start VARCHAR(5),
friday_start VARCHAR(5),
saturday_start VARCHAR(5),
sunday_start VARCHAR(5),
monday_end VARCHAR(5),
tuesday_end VARCHAR(5),
wednesday_end VARCHAR(5),
thursday_end VARCHAR(5),
friday_end VARCHAR(5),
saturday_end VARCHAR(5),
sunday_end VARCHAR(5),
monday_location VARCHAR(50),
tuesday_location VARCHAR(50),
wednesday_location VARCHAR(50),
thursday_location VARCHAR(50),
friday_location VARCHAR(50),
saturday_location VARCHAR(50),
sunday_location VARCHAR(50),
FOREIGN KEY (id) REFERENCES seller(id)
);
