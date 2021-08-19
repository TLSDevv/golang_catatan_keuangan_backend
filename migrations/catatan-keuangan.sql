CREATE DATABASE catatan_keuangan;
CREATE TABLE transactions (
    id BIGINT PRIMARY KEY,
    user_id BIGINT,
    name_transaction VARCHAR(50),
    type_transaction CHAR(10),
    category_id BIGINT,
    nominal INTEGER,
    description VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_categories FOREIGN KEY(category_id) REFERENCES categories(id)
);

CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    username VARCHAR (100) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    gender CHAR(2) NOT NULL,
    age INT NOT NULL,
    job VARCHAR(100)
);

CREATE TABLE categories (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name_category VARCHAR(50) NOT NULL,
    description VARCHAR(100),
    icon_name VARCHAR(20) NOT NULL,
    icon_color CHAR(10) NOT NULL,
    transaction_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(id)
);