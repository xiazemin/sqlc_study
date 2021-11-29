-- name: ListMeasurements :many
select * from measurements where `value` in (?);