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
about TEXT,
logo TEXT,
image TEXT,
second_image TEXT,
third_image TEXT,
fourth_image TEXT,
phone VARCHAR(20),
location VARCHAR(100),
email VARCHAR(40),
twitter VARCHAR(20),
facebook VARCHAR(50),
instagram VARCHAR(30),
pinterest VARCHAR(30)
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
monday_start VARCHAR(10),
tuesday_start VARCHAR(10),
wednesday_start VARCHAR(10),
thursday_start VARCHAR(10),
friday_start VARCHAR(10),
saturday_start VARCHAR(10),
sunday_start VARCHAR(10),
monday_end VARCHAR(10),
tuesday_end VARCHAR(10),
wednesday_end VARCHAR(10),
thursday_end VARCHAR(10),
friday_end VARCHAR(10),
saturday_end VARCHAR(10),
sunday_end VARCHAR(10),
monday_location VARCHAR(100),
tuesday_location VARCHAR(100),
wednesday_location VARCHAR(100),
thursday_location VARCHAR(100),
friday_location VARCHAR(100),
saturday_location VARCHAR(100),
sunday_location VARCHAR(100),
monday_address VARCHAR(100),
tuesday_address VARCHAR(100),
wednesday_address VARCHAR(100),
thursday_address VARCHAR(100),
friday_address VARCHAR(100),
saturday_address VARCHAR(100),
sunday_address VARCHAR(100),

FOREIGN KEY (id) REFERENCES seller(id)
);
