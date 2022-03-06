CREATE DATABASE keuanganku;

CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT COMMENT "The user ID",
    username VARCHAR(25) COMMENT "The user username",
    email VARCHAR(255) COMMENT "The user email",
    password VARCHAR(255) COMMENT "The user password, crypted using Bcrypt",
    fullname VARCHAR(50) COMMENT "The user fullname",
    role INT NOT NULL COMMENT "The user role, either 1 for admin, either for regular user",
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "User created time",
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "User updated time",
    deleted_at DATETIME NULL COMMENT "User deleted time",

    PRIMARY KEY(id)
);

CREATE TABLE auths(
    id INT NOT NULL AUTO_INCREMENT COMMENT "Auth ID",
    user_id INT NOT NULL COMMENT "User ID that owned the token",
    refresh_token VARCHAR(255) COMMENT "The refresh token",

    PRIMARY KEY(id),
    CONSTRAINT fk_auth_users
		FOREIGN KEY(user_id) REFERENCES users(id)
)

CREATE TABLE transactions(
    id INT NOT NULL AUTO_INCREMENT COMMENT "Transaction ID",
    user_id INT NOT NULL COMMENT "User ID that inserted the transaction",
    transaction_name VARCHAR(255) NOT NULL COMMENT "The name of the transaction",
    category VARCHAR(100) NOT NULL COMMENT "Category of transaction",
    transaction_type TINYINT(1) NOT NULL COMMENT "Transaction Type, either 1 or 0, 1 = income, 0 = outcome",
    amount BIGINT NOT NULL COMMENT "The amount of transaction spent",
    transaction_at DATETIME NULL COMMENT "Transaction timestamp",
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "Transaction created time",
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "Transaction updated time",
    deleted_at DATETIME NULL COMMENT "Transaction deleted time",

    PRIMARY KEY(id),
    CONSTRAINT fk_users
		FOREIGN KEY(user_id) REFERENCES users(id)
);