-- +goose Up
CREATE TABLE feeds (
	id uuid DEFAULT gen_random_uuid(),
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	name VARCHAR NOT NULL,
	url VARCHAR UNIQUE NOT NULL, 
	user_id uuid NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY(user_id) REFERENCES users(id)
	ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;