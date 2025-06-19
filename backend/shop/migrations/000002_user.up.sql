CREATE TABLE IF NOT EXISTS users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       first_name VARCHAR(100) NOT NULL,
                       last_name VARCHAR(100) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       role VARCHAR(50) NOT NULL DEFAULT 'user',
                       password varchar(255) NOT NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
