CREATE TABLE accounts (
    account_id SERIAL NOT NULL,
    player_id VARCHAR(6) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(320) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL,
    is_ban BOOLEAN DEFAULT FALSE,
    create_at timestamp DEFAULT now(),
    CONSTRAINT accounts_pk PRIMARY KEY (account_id),
    CONSTRAINT accounts_unique UNIQUE (player_id, username, email)
)

