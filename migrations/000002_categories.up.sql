CREATE TABLE IF NOT EXISTS categories (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name_category VARCHAR(50) NOT NULL,
    description VARCHAR(100),
    icon_name VARCHAR(20) NOT NULL,
    icon_color CHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(id)
)