-- name: ListItemsByID :many 
select * from items where item_id in (?);