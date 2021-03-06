CREATE TABLE user_tokens (
    id SERIAL PRIMARY KEY,
    token VARCHAR(256) NOT NULL,
    user_id INTEGER NOT NULL,
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- add constraint where user can only have 1 active token
CREATE UNIQUE INDEX user_active_constraint ON user_tokens (user_id, active) WHERE active;
