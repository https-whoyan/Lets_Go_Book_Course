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
)
