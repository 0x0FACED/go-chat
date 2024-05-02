package utils

const (
	ErrUsernameLength = "username len must be <= 16 and not 0"
	ErrPasswordLength = "pass len must be >= 5"

	ErrIncorrectUsernameOrPass = "incorrect username or pass"
	ErrRollbackTx              = "cant rollback tx:"
	ErrCommitTx                = "cant commit tx:"

	QueryLoginTx = "SELECT id, name from users where username = ? and password = ?"
)
