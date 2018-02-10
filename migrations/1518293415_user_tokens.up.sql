CREATE TABLE user_tokens (
    id SERIAL PRIMARY KEY,
    token VARCHAR(128) NOT NULL,
    user_id INTEGER NOT NULL,
    active BOOLEAN DEFAULT TRUE
);

-- add constraint where user must not have the same user_id and token
ALTER TABLE user_tokens ADD UNIQUE (user_id, token);

-- add constraint where user can only have 1 active token
CREATE UNIQUE INDEX user_active_constraint ON user_tokens (user_id, active) WHERE active;
