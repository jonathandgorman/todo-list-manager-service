CREATE DATABASE todo;
CREATE TABLE todo.lists (
                       id  VARCHAR(100) PRIMARY KEY,
                       user_id VARCHAR(100) NOT NULL,
                       name VARCHAR(100) NOT NULL,
                       created TIMESTAMP SET DEFAULT now()
);

CREATE TABLE todo.accounts (
                          account_id  UUID PRIMARY KEY,
                          user_id  VARCHAR(100),
                          username VARCHAR(100),
                          password_hash VARCHAR(100),
                          created_at TIMESTAMP DEFAULT now()
);