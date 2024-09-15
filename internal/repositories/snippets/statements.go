package snippets

const (
	allColumnsSnippets = ` id, content, content, created, expires `

	insertSnippetStatement = `
		INSERT INTO snippets 
    		(title, content, created, expires) 
		VALUES ($1, $2, NOW(), NOW() + MAKE_INTERVAL(days => $3))
		RETURNING id;`

	selectSnippetStatement = `
		SELECT 
    		` + allColumnsSnippets + `
		FROM snippets 
		WHERE expires > NOW() AND id = $1;`

	multipleSelectSnippet = `
		SELECT
			` + allColumnsSnippets + `
		FROM snippets 
		WHERE expires > NOW() 
		ORDER BY id
		LIMIT 10;
	`
)
