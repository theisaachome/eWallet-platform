

CREATE DATABASE eWallet_db;
-- Enable UUID extension (only needs to be done once per database)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Define ENUM type for status (PostgreSQL doesn't use inline ENUMs)
DO $$
    BEGIN
        IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
            CREATE TYPE user_status AS ENUM ('active', 'suspended', 'deleted');
        END IF;
    END$$;

-- Create users table
CREATE TABLE users (
                       id BIGSERIAL PRIMARY KEY,
                       public_id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
                       phone_number VARCHAR(20) UNIQUE NOT NULL,
                       hash_password VARCHAR(255) NOT NULL ,
                       status user_status DEFAULT 'active',
                       created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE wallet_status AS ENUM ('active', 'frozen');
CREATE TABLE wallets (
                         id BIGSERIAL PRIMARY KEY,
                         public_id UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
                         user_id BIGINT NOT NULL UNIQUE,
                         balance DECIMAL(18,2) DEFAULT 0.00,
                         currency CHAR(3) DEFAULT 'MYR',
                         status wallet_status DEFAULT 'active',
                         created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                         FOREIGN KEY (user_id) REFERENCES users(id)
);
