CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id BIGINT,
    name_transaction VARCHAR(50),
    type_transaction CHAR(10),
    category_id BIGINT,
    amount INTEGER,
    description VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_categories FOREIGN KEY(category_id) REFERENCES categories(id)
)