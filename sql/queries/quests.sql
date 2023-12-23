-- name: GetOneQuest :one
SELECT * from quests
WHERE id = ?
ORDER BY number
LIMIT 1;

-- name: GetLastQuest :one
SELECT MAX(number) from quests
WHERE campaign_id = ?;


-- name: GetAllCampaignQuests :many
SELECT * FROM quests
WHERE campaign_id = ?
ORDER BY number;

-- name: GetAllCampaignActiveQuests :many
SELECT * FROM quests
WHERE campaign_id = ? AND is_active = true
ORDER BY number;

-- name: GetAllCampaignDoneQuests :many
SELECT * FROM quests
WHERE campaign_id = ? AND is_complete = true
ORDER BY number;

-- name: CreateQuest :execresult
INSERT INTO quests (created_at, updated_at, name, description, completed_description, npc, number, campaign_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

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