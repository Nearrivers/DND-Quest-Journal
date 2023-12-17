-- +goose Up

CREATE TABLE quests (
  id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  description LONGTEXT NOT NULL,
  npc TEXT NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT FALSE,
  is_complete BOOLEAN NOT NULL DEFAULT FALSE,
  number INT NOT NULL,
  campaign_id INT NOT NULL
);

ALTER TABLE quests
ADD FOREIGN KEY (campaign_id) REFERENCES campaigns(id);

-- +goose Down

DROP TABLE quests;