package users

const (
	allColumnsUsers     = ` id, name, mail, pass, created_at `
	insertUserStatement = `
		INSERT INTO users (
		    name, mail, pass, created_at 
		) VALUES (
		    $1, $2, $3, NOW()
		)
	`
	selectByEmail = `
		SELECT 
	` + allColumnsUsers + ` 
		FROM users 
		WHERE mail = $1
	`
	selectById = `SELECT ` +
		allColumnsUsers + ` FROM users WHERE id = $1`
)
