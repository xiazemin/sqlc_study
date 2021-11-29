// Code generated by sqlc. DO NOT EDIT.
// source: items.sql

package gorse

import (
	"context"
	"fmt"
)

const listItemsByID = `-- name: ListItemsByID :many
select item_id, time_stamp, labels, comment from items where item_id in (?)
`

func (q *Queries) ListItemsByID(ctx context.Context, itemID []string) ([]Item, error) {

	if len(itemID) <= 0 {
		return nil, fmt.Errorf("itemID length is invalid")
	}
	param := "?"
	for i := 0; i < len(itemID)-1; i++ {
		param += ",?"
	}
	listItemsByID := replaceNth(listItemsByID, "(?)", "( "+param+" )", 1)

	rows, err := q.db.QueryContext(ctx, listItemsByID, stringSlice2interface(itemID)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ItemID,
			&i.TimeStamp,
			&i.Labels,
			&i.Comment,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
