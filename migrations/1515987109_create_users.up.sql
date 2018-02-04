CREATE TABLE users (
    id                  serial PRIMARY KEY,
    obfuscated_id       VARCHAR(128) NOT NULL UNIQUE,
    email               VARCHAR(128) NOT NULL UNIQUE,
    encrypted_password  VARCHAR(128) NOT NULL,
    created_at          TIMESTAMP,
    updated_at          TIMESTAMP,
    active              BOOLEAN DEFAULT TRUE
);
