package dao

const (
	URLTable = `
CREATE TABLE IF NOT EXISTS url
(
    url_id  serial PRIMARY KEY,
    original text NOT NULL UNIQUE, 
	created_by text,
	is_deleted boolean NOT NULL DEFAULT false
);
`
)
