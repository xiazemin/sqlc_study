// Code generated by sqlc. DO NOT EDIT.

package gen

import (
	"database/sql"
)

type Author struct {
	ID         int32          `json:"id"`
	Name       string         `json:"name"`
	Bio        sql.NullString `json:"bio"`
	CompanyID  int32          `json:"company_id"`
	Size       sql.NullInt64  `json:"size"`
	EmptyCol   sql.NullInt32  `json:"empty_col"`
	DefaultCol int32          `json:"default_col"`
	Size1      sql.NullInt32  `json:"size1"`
}

type Company struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	CompanyName sql.NullString `json:"companyName"`
}

type Employee struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	CompanyName sql.NullString `json:"companyName"`
}
