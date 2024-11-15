CREATE TABLE transactions
(
    id             SERIAL PRIMARY KEY CHECK (id >= 0),
    user_id        BIGINT REFERENCES users(id) CHECK (user_id >= 0),
    transaction_id VARCHAR(255) NOT NULL UNIQUE, --unique creates a unique index(?) for faster search.
    --if not, maybe a cache system would do better
    state VARCHAR(15) NOT NULL,
    amount BIGINT NOT NULL CHECK (amount >= 0),

    created_at     TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);