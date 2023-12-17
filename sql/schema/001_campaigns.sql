-- +goose Up

CREATE TABLE campaigns (
  id INT PRIMARY KEY  NOT NULL AUTO_INCREMENT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL
);

-- +goose Down

DROP TABLE campaigns;