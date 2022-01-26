CREATE DATABASE keuanganku;

CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(25),
    email VARCHAR(255),
    password VARCHAR(255),
    fullname VARCHAR(50),
    token VARCHAR(255),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,

    PRIMARY KEY(id)
);

CREATE TABLE transactions(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    name VARCHAR(255),
    category VARCHAR(100),
    type TINYINT(1),
    amount BIGINT,
    transaction_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,

    PRIMARY KEY(id),
    CONSTRAINT fk_users
		FOREIGN KEY(user_id) REFERENCES users(id)
);