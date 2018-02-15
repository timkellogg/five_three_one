CREATE TABLE user_secrets (
    id              SERIAL PRIMARY KEY,
    user_id         INTEGER NOT NULL,
    client_id       VARCHAR(128) NOT NULL,
    client_secret   VARCHAR(128) NOT NULL,
    active          BOOLEAN DEFAULT TRUE,
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP
);

-- add constraint where user can only have 1 active token
CREATE UNIQUE INDEX user_secret_active_constraint ON user_secrets (user_id, active) WHERE active;
