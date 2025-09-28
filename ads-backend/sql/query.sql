-- name: GetMaximumWeights :many
SELECT * FROM maximum_weight;

-- name: CreateMaximumWeight :exec
INSERT INTO maximum_weight (
    uuid, subtype, name, taxi, landing, take_off, zero_fuel
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
);

-- name: DeleteMaximumWeight :exec
DELETE FROM maximum_weight
WHERE uuid = $1;

-- name: GetCarriers :many
SELECT * FROM carrier;

-- name: GetFleets :many
SELECT * FROM fleet;

-- name: GetCarrierFleets :many
SELECT * FROM fleet
WHERE carrier = $1;

-- name: GetSubtypes :many
SELECT * FROM subtype;

-- name: GetFleetSubtypes :many
SELECT * FROM subtype
WHERE fleet = $1;

-- name: GetCarrierSubtypes :many
SELECT * from subtype
WHERE fleet IN (
    SELECT uuid FROM fleet
    WHERE carrier = $1
);
