package models

const (
	allColumns = ` id, content, content, created, expires `

	insertSnippetStatement = `
		INSERT INTO snippets 
    		(title, content, created, expires) 
		VALUES ($1, $2, NOW(), NOW() + MAKE_INTERVAL(days => $3))
		RETURNING id;`

	selectSnippetStatement = `
		SELECT 
    		` + allColumns + `
		FROM snippets 
		WHERE expires > NOW() AND id = $1;`

	multipleSelect = `
		SELECT
			` + allColumns + `
		FROM snippets 
		WHERE expires > NOW() 
		ORDER BY id
		LIMIT 10;
	`
)
