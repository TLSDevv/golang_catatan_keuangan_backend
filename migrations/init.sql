CREATE DATABASE keuanganku;

CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT,
    username VARCHAR(25),
    email VARCHAR(255),
    password VARCHAR(255),
    fullname VARCHAR(50),
    role INT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME,

    PRIMARY KEY(id)
);

CREATE TABLE auths(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    refresh_token VARCHAR(255),

    PRIMARY KEY(id),
    CONSTRAINT fk_auth_users
		FOREIGN KEY(user_id) REFERENCES users(id)
)

CREATE TABLE transactions(
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    transaction_name VARCHAR(255) NULL,
    category VARCHAR(100) NULL,
    transaction_type TINYINT(1) NULL,
    amount BIGINT NULL,
    transaction_at DATETIME NULL,
    created_at DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,

    PRIMARY KEY(id),
    CONSTRAINT fk_users
		FOREIGN KEY(user_id) REFERENCES users(id)
);