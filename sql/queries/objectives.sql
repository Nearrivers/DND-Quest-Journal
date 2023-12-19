-- name: GetOneObjective :one
SELECT * FROM objectives
WHERE id = ?
ORDER BY number
LIMIT 1;

-- name: GetAllQuestObjectives :many
SELECT * from objectives
WHERE quest_id = ?
ORDER BY number;

-- name: GetAllQuestActiveObjectives :many
SELECT * from objectives
WHERE quest_id = ? AND is_active = true
ORDER BY number;

-- name: GetAllQuestDoneObjectives :many
SELECT * from objectives
WHERE quest_id = ? AND is_complete = true
ORDER BY number;

-- name: CreateObjective :execresult
INSERT INTO objectives (created_at, updated_at, name, description, number, quest_id)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateObjective :execresult
UPDATE objectives SET name = ?, description = ?, updated_at = ?
WHERE id = ?;

-- name: ReorderObjective :execresult
UPDATE objectives SET number = ?
WHERE id = ?;

-- name: ActivateObjective :execresult
UPDATE objectives SET is_active = true
WHERE id = ?;

-- name: FinishObjective :execresult
UPDATE objectives SET is_complete = true
WHERE id = ?;