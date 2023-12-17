-- name: GetOneQuest :one
SELECT * from quests
WHERE id = ? LIMIT 1;

-- name: GetAllCampaignQuests :many
SELECT * FROM quests
WHERE campaign_id = ?;

-- name: GetAllCampaignActiveQuests :many
SELECT * FROM quests
WHERE campaign_id = ? AND is_active = true;

-- name: GetAllCampaignDoneQuests :many
SELECT * FROM quests
WHERE campaign_id = ? AND is_complete = true;

-- name: CreateQuest :execresult
INSERT INTO quests (created_at, updated_at, name, description, npc, number, campaign_id)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateQuest :execresult
UPDATE quests SET name = ?, description = ?, npc = ?, updated_at = ?
WHERE id = ?;

-- name: ReorderQuest :execresult
UPDATE quests SET number = ?
WHERE id = ?;

-- name: ActivateQuest :execresult
UPDATE quests SET is_active = true
WHERE id = ?;

-- name: FinishQuest :execresult
UPDATE quests SET is_complete = true
WHERE id = ?;