package utils

const (
	ErrOpenDB     = "cant open db:"
	ErrPingDB     = "cant ping db:"
	ErrCloseDB    = "cant close db:"
	ErrDriverName = "unknown db driver"
)

const (
	ErrUsernameLength = "username len must be <= 16 and not 0"
	ErrPasswordLength = "pass len must be >= 5"

	ErrUserExists = "user exists"

	ErrIncorrectUsernameOrPass = "incorrect username or pass"
)

const (
	ErrBeginTx     = "cant begin tx:"
	ErrRollbackTx  = "cant rollback tx:"
	ErrCommitTx    = "cant commit tx:"
	ErrPrepareTx   = "cant prepare tx:"
	ErrExecQueryTx = "cant exec or query to db tx:"
)

const (
	QueryLoginTx       = `SELECT id, name, username, registration_date, description FROM users WHERE username = $1 AND password = $2`
	QueryRegisterTx    = `INSERT INTO users (name, username, password, description) VALUES ($1, $2, $3, $4)`
	QueryGetPasswordTx = `SELECT password FROM users WHERE username = $1`
	QuerySaveMessageTx = `INSERT INTO messages (sender_id, receiver_id, text, chat_id) VALUES ($1, $2, $3, $4) RETURNING id, timestamp`
	QueryCreateChatTx  = `INSERT INTO chats (first_user_id, second_user_id) VALUES ($1, $2) RETURNING id`
	QueryFindChatTx    = `SELECT id FROM chats WHERE (first_user_id = $1 AND second_user_id = $2) OR (first_user_id = $2 AND second_user_id = $1)`
)
