CREATE TABLE transactions
(
    id             NUMERIC(20, 0) PRIMARY KEY CHECK (id >= 0),
    user_id        BIGINT REFERENCES users(id) CHECK (user_id >= 0),
    transaction_id VARCHAR(255) NOT NULL UNIQUE,
    state VARCHAR(50) NOT NULL,
    amount BIGINT NOT NULL CHECK (amount >= 0),

    created_at     TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);