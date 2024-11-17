-- +goose Up
CREATE TABLE posts (
	id uuid DEFAULT gen_random_uuid(),
	created_at TIMESTAMP NOT NULL, 
	updated_at TIMESTAMP NOT NULL,
	title TEXT NOT NULL, 
	url VARCHAR UNIQUE NOT NULL,
	description TEXT,
	published_at TIMESTAMP,
	feed_id uuid NOT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY (feed_id) REFERENCES feeds(id)
	ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;