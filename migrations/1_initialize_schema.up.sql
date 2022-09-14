CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create Table Account
CREATE TABLE IF NOT EXISTS accounts (
      account_id UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
      player_id VARCHAR(11) NOT NULL UNIQUE CHECK ( player_id <> '' ),
      username VARCHAR(255) NOT NULL UNIQUE CHECK ( username <> '' ),
      email VARCHAR(320) NOT NULL UNIQUE CHECK ( email ~ '^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$' ),
      password_hash varchar(255) NOT NULL,
      is_ban BOOLEAN DEFAULT FALSE,
      created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);