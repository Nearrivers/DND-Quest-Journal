-- name: GetOneCampaign :one
SELECT * FROM campaigns
WHERE id = ? LIMIT 1;

-- name: GetAllCampaigns :many
SELECT * FROM campaigns;

-- name: CreateCampaign :execresult
INSERT INTO campaigns (created_at, updated_at, name)
VALUES (?, ?, ?);

-- name: UpdateCampaign :execresult
UPDATE campaigns SET name = ?, updated_at = ? WHERE id = ?;

-- name: DeleteCampaign :exec
DELETE FROM campaigns WHERE id = ?;