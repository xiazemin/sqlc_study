// Code generated by sqlc. DO NOT EDIT.

package gen

import (
	"database/sql"
	"fmt"
)

type AuthorsType string

const (
	AuthorsTypeHuLianWangyouXiruanJian AuthorsType = "互联网/游戏/软件"
	AuthorsTypeJiaoYupeiXun            AuthorsType = "教育/培训"
	AuthorsType130ren                  AuthorsType = "1-30人"
	AuthorsTypeNULL                    AuthorsType = "null"
)

func (e *AuthorsType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AuthorsType(s)
	case string:
		*e = AuthorsType(s)

	case nil:
		*e = AuthorsTypeNULL

	default:
		return fmt.Errorf("unsupported scan type for AuthorsType: %T", src)
	}
	return nil
}

type Author struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Bio         sql.NullString `json:"bio"`
	CompanyID   int32          `json:"company_id"`
	Size        sql.NullInt64  `json:"size"`
	EmptyCol    sql.NullInt32  `json:"empty_col"`
	DefaultCol  int32          `json:"default_col"`
	Size1       sql.NullInt32  `json:"size1"`
	DefaultCol1 int32          `json:"default_col1"`
	Type        AuthorsType    `json:"type"`
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
