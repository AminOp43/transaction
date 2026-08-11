CREATE TABLE transactions (
                              id BIGSERIAL PRIMARY KEY,
                              user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                              type VARCHAR(20) NOT NULL CHECK (type IN ('income', 'expense')),
                              amount BIGINT NOT NULL CHECK (amount > 0),
                              category VARCHAR(100) NOT NULL,
                              description TEXT,
                              occurred_at TIMESTAMPTZ NOT NULL,
                              created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                              updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);