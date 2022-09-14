package repository

const (
	createAccountQuery = `INSERT INTO accounts (account_id, player_id, username, email, password_hash, is_ban, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, now(), now()) RETURNING account_id, player_id, username, email, password_hash, is_ban, created_at, updated_at`

	updateAccountQuery = `UPDATE accounts a SET 
                      username=COALESCE(NULLIF($1, ''), username), 
                      email=COALESCE(NULLIF($2, ''), email), 
                      password_hash=COALESCE(NULLIF($3, ''), password_hash),
                      is_ban=COALESCE(NULLIF($4, false), is_ban),
                      updated_at = now()
                      WHERE account_id=$5
                      RETURNING account_id, player_id, username, email, password_hash, is_ban, created_at, updated_at`

	getAccountByIdQuery = `SELECT a.account_id, a.player_id, a.username, a.email, a.password_hash, a.is_ban, a.created_at, a.updated_at
	FROM accounts a WHERE a.account_id = $1`

	deleteAccountByIdQuery = `DELETE FROM accounts WHERE account_id = $1`
)
