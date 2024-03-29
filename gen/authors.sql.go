// Code generated by sqlc. DO NOT EDIT.
// source: authors.sql

package gen

import (
	"context"
	"database/sql"
	"fmt"
)

const batchCreateAuthor = `-- name: BatchCreateAuthor :execresult
INSERT INTO authors (
  id,name,bio,company_id,default_col,default_col1
) VALUES %s ON DUPLICATE KEY UPDATE id=VALUES(id)
`

type BatchCreateAuthorParams struct {
	ID int32 `json:"id"`

	Name string `json:"name"`

	Bio sql.NullString `json:"bio"`

	DefaultCol int32 `json:"default_col"`
}

func (q *Queries) BatchCreateAuthor(ctx context.Context, arg []BatchCreateAuthorParams) (sql.Result, error) {

	batchCreateAuthor := repeatN(batchCreateAuthor, "(?,?,?,1,?,4)", len(arg))
	var args []interface{}
	for i := 0; i < len(arg); i++ {
		args = append(args, arg[i].ID)
		args = append(args, arg[i].Name)
		args = append(args, arg[i].Bio)
		args = append(args, arg[i].DefaultCol)
	}
	return q.db.ExecContext(ctx, batchCreateAuthor, args...)
}

const createAuthor = `-- name: CreateAuthor :execresult
INSERT INTO authors (
  id,name,bio,company_id
) VALUES (
  ?,?, ?,1 
)
`

type CreateAuthorParams struct {
	ID int32 `json:"id"`

	Name string `json:"name"`

	Bio sql.NullString `json:"bio"`
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (sql.Result, error) {

	createAuthor := createAuthor

	return q.db.ExecContext(ctx, createAuthor, arg.ID, arg.Name, arg.Bio)
}

const deleteAuthor = `-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?
`

func (q *Queries) DeleteAuthor(ctx context.Context, id int32) error {

	_, err := q.db.ExecContext(ctx, deleteAuthor, id)
	return err
}

const deleteAuthorIn = `-- name: DeleteAuthorIn :exec
DELETE FROM authors
WHERE id in (?)
`

func (q *Queries) DeleteAuthorIn(ctx context.Context, id []int32) error {

	if len(id) <= 0 {
		return fmt.Errorf("id length is invalid")
	}
	param := "?"
	for i := 0; i < len(id)-1; i++ {
		param += ",?"
	}
	deleteAuthorIn := replaceNth(deleteAuthorIn, "(?)", "( "+param+" )", 1)

	_, err := q.db.ExecContext(ctx, deleteAuthorIn, int32Slice2interface(id)...)
	return err
}

const getAuthorsInCompany = `-- name: GetAuthorsInCompany :many
SELECT id, name, bio, company_id, size, empty_col, default_col, size1, default_col1, type, type1 FROM authors where company_id in ( select id from company where id in (?) and name in (?) )
`

type GetAuthorsInCompanyParams struct {
	ID []int32 `json:"id"`

	Name []string `json:"name"`
}

func (q *Queries) GetAuthorsInCompany(ctx context.Context, arg GetAuthorsInCompanyParams) ([]Author, error) {

	getAuthorsInCompany := getAuthorsInCompany

	if len(arg.ID) <= 0 {
		return nil, fmt.Errorf("arg.ID length is invalid")
	}
	{
		param := "?"
		for i := 0; i < len(arg.ID)-1; i++ {
			param += ",?"
		}
		getAuthorsInCompany = replaceNth(getAuthorsInCompany, "(?)", "( "+param+" )", 1)
	}

	if len(arg.Name) <= 0 {
		return nil, fmt.Errorf("arg.Name length is invalid")
	}
	{
		param := "?"
		for i := 0; i < len(arg.Name)-1; i++ {
			param += ",?"
		}
		getAuthorsInCompany = replaceNth(getAuthorsInCompany, "(?)", "( "+param+" )", 1)
	}

	rows, err := q.db.QueryContext(ctx, getAuthorsInCompany, append(int32Slice2interface(arg.ID), stringSlice2interface(arg.Name)...)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CompanyID,
			&i.Size,
			&i.EmptyCol,
			&i.DefaultCol,
			&i.Size1,
			&i.DefaultCol1,
			&i.Type,
			&i.typeOverrides,
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

const getAuthorsInCompanyById = `-- name: GetAuthorsInCompanyById :many
SELECT id, name, bio, company_id, size, empty_col, default_col, size1, default_col1, type, type1 FROM authors where company_id in ( select id from company where id in (?) )
`

func (q *Queries) GetAuthorsInCompanyById(ctx context.Context, id []int32) ([]Author, error) {

	if len(id) <= 0 {
		return nil, fmt.Errorf("id length is invalid")
	}
	param := "?"
	for i := 0; i < len(id)-1; i++ {
		param += ",?"
	}
	getAuthorsInCompanyById := replaceNth(getAuthorsInCompanyById, "(?)", "( "+param+" )", 1)

	rows, err := q.db.QueryContext(ctx, getAuthorsInCompanyById, int32Slice2interface(id)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CompanyID,
			&i.Size,
			&i.EmptyCol,
			&i.DefaultCol,
			&i.Size1,
			&i.DefaultCol1,
			&i.Type,
			&i.typeOverrides,
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

const getAuthorsInOneCompany = `-- name: GetAuthorsInOneCompany :many
SELECT id, name, bio, company_id, size, empty_col, default_col, size1, default_col1, type, type1 FROM authors where company_id in ( select id from company where id = ? )
`

func (q *Queries) GetAuthorsInOneCompany(ctx context.Context, id int32) ([]Author, error) {

	rows, err := q.db.QueryContext(ctx, getAuthorsInOneCompany, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CompanyID,
			&i.Size,
			&i.EmptyCol,
			&i.DefaultCol,
			&i.Size1,
			&i.DefaultCol1,
			&i.Type,
			&i.typeOverrides,
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

const getLargestID = `-- name: GetLargestID :one
SELECT id FROM authors order by id desc limit 1
`

func (q *Queries) GetLargestID(ctx context.Context) (int32, error) {

	row := q.db.QueryRowContext(ctx, getLargestID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getMaxID = `-- name: GetMaxID :one
SELECT MAX(id) FROM authors
`

func (q *Queries) GetMaxID(ctx context.Context) (sql.NullInt32, error) {

	row := q.db.QueryRowContext(ctx, getMaxID)
	var max sql.NullInt32
	err := row.Scan(&max)
	return max, err
}

const getMax_default_col = `-- name: GetMax_default_col :one
select max(default_col) from authors where id =?
`

func (q *Queries) GetMax_default_col(ctx context.Context, id int32) (sql.NullInt32, error) {

	row := q.db.QueryRowContext(ctx, getMax_default_col, id)
	var max sql.NullInt32
	err := row.Scan(&max)
	return max, err
}

const getMax_default_col1 = `-- name: GetMax_default_col1 :one
select max(default_col1) from authors where id =?
`

func (q *Queries) GetMax_default_col1(ctx context.Context, id int32) (sql.NullInt32, error) {

	row := q.db.QueryRowContext(ctx, getMax_default_col1, id)
	var max sql.NullInt32
	err := row.Scan(&max)
	return max, err
}

const getMax_empty_col = `-- name: GetMax_empty_col :one
select max(empty_col) from authors where id =?
`

func (q *Queries) GetMax_empty_col(ctx context.Context, id int32) (sql.NullInt32, error) {

	row := q.db.QueryRowContext(ctx, getMax_empty_col, id)
	var max sql.NullInt32
	err := row.Scan(&max)
	return max, err
}

const getOneAuthor = `-- name: GetOneAuthor :one
SELECT id, name, bio, company_id, size, empty_col, default_col, size1, default_col1, type, type1 FROM authors where  id in (?)  and bio=? and  name in (?) and company_id in (?) limit 1
`

type GetOneAuthorParams struct {
	ID []int32 `json:"id"`

	Bio sql.NullString `json:"bio"`

	Name []string `json:"name"`

	CompanyID []int32 `json:"company_id"`
}

func (q *Queries) GetOneAuthor(ctx context.Context, arg GetOneAuthorParams) (Author, error) {

	getOneAuthor := getOneAuthor

	if len(arg.ID) <= 0 {
		return Author{}, fmt.Errorf("arg.ID length is invalid")
	}
	{
		param := "?"
		for i := 0; i < len(arg.ID)-1; i++ {
			param += ",?"
		}
		getOneAuthor = replaceNth(getOneAuthor, "(?)", "( "+param+" )", 1)
	}

	if len(arg.Name) <= 0 {
		return Author{}, fmt.Errorf("arg.Name length is invalid")
	}
	{
		param := "?"
		for i := 0; i < len(arg.Name)-1; i++ {
			param += ",?"
		}
		getOneAuthor = replaceNth(getOneAuthor, "(?)", "( "+param+" )", 1)
	}

	if len(arg.CompanyID) <= 0 {
		return Author{}, fmt.Errorf("arg.CompanyID length is invalid")
	}
	{
		param := "?"
		for i := 0; i < len(arg.CompanyID)-1; i++ {
			param += ",?"
		}
		getOneAuthor = replaceNth(getOneAuthor, "(?)", "( "+param+" )", 1)
	}

	row := q.db.QueryRowContext(ctx, getOneAuthor, append(append(append(int32Slice2interface(arg.ID), arg.Bio), stringSlice2interface(arg.Name)...), int32Slice2interface(arg.CompanyID)...)...)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Bio,
		&i.CompanyID,
		&i.Size,
		&i.EmptyCol,
		&i.DefaultCol,
		&i.Size1,
		&i.DefaultCol1,
		&i.Type,
		&i.typeOverrides,
	)
	return i, err
}

const getTotalSize = `-- name: GetTotalSize :one
SELECT ifnull(sum(size),0) from authors WHERE id in (?)
`

func (q *Queries) GetTotalSize(ctx context.Context, id []int32) (int64, error) {

	if len(id) <= 0 {
		return 0, fmt.Errorf("id length is invalid")
	}
	param := "?"
	for i := 0; i < len(id)-1; i++ {
		param += ",?"
	}
	getTotalSize := replaceNth(getTotalSize, "(?)", "( "+param+" )", 1)

	row := q.db.QueryRowContext(ctx, getTotalSize, int32Slice2interface(id)...)
	var ifnull int64
	err := row.Scan(&ifnull)
	return ifnull, err
}

const getTotalSize1Null = `-- name: GetTotalSize1Null :one
SELECT max(size1) from authors WHERE id = ?
`

func (q *Queries) GetTotalSize1Null(ctx context.Context, id int32) (sql.NullInt32, error) {

	row := q.db.QueryRowContext(ctx, getTotalSize1Null, id)
	var max sql.NullInt32
	err := row.Scan(&max)
	return max, err
}

const getTotalSize1NullIn = `-- name: GetTotalSize1NullIn :one
SELECT sum(size1) from authors WHERE id in (?)
`

func (q *Queries) GetTotalSize1NullIn(ctx context.Context, id []int32) (sql.NullInt32, error) {

	if len(id) <= 0 {
		return sql.NullInt32{}, fmt.Errorf("id length is invalid")
	}
	param := "?"
	for i := 0; i < len(id)-1; i++ {
		param += ",?"
	}
	getTotalSize1NullIn := replaceNth(getTotalSize1NullIn, "(?)", "( "+param+" )", 1)

	row := q.db.QueryRowContext(ctx, getTotalSize1NullIn, int32Slice2interface(id)...)
	var sum sql.NullInt32
	err := row.Scan(&sum)
	return sum, err
}

const getTotalSizeNull = `-- name: GetTotalSizeNull :one
SELECT sum(size) from authors WHERE id =?
`

func (q *Queries) GetTotalSizeNull(ctx context.Context, id int32) (sql.NullInt64, error) {

	row := q.db.QueryRowContext(ctx, getTotalSizeNull, id)
	var sum sql.NullInt64
	err := row.Scan(&sum)
	return sum, err
}

const getTotalSizeNullIn = `-- name: GetTotalSizeNullIn :one
SELECT sum(size) from authors WHERE id in (?)
`

func (q *Queries) GetTotalSizeNullIn(ctx context.Context, id []int32) (sql.NullInt64, error) {

	if len(id) <= 0 {
		return sql.NullInt64{}, fmt.Errorf("id length is invalid")
	}
	param := "?"
	for i := 0; i < len(id)-1; i++ {
		param += ",?"
	}
	getTotalSizeNullIn := replaceNth(getTotalSizeNullIn, "(?)", "( "+param+" )", 1)

	row := q.db.QueryRowContext(ctx, getTotalSizeNullIn, int32Slice2interface(id)...)
	var sum sql.NullInt64
	err := row.Scan(&sum)
	return sum, err
}

const listAllAuthors = `-- name: ListAllAuthors :many
SELECT id, name, bio, company_id, size, empty_col, default_col, size1, default_col1, type, type1 FROM authors
ORDER BY name
`

func (q *Queries) ListAllAuthors(ctx context.Context) ([]Author, error) {

	rows, err := q.db.QueryContext(ctx, listAllAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CompanyID,
			&i.Size,
			&i.EmptyCol,
			&i.DefaultCol,
			&i.Size1,
			&i.DefaultCol1,
			&i.Type,
			&i.typeOverrides,
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

const listAuthors = `-- name: ListAuthors :many
SELECT id, name, bio, company_id, size, empty_col, default_col, size1, default_col1, type, type1 FROM authors where  bio=? and id in (?)  and name in (?)  ORDER BY name
`

type ListAuthorsParams struct {
	Bio sql.NullString `json:"bio"`

	ID []int32 `json:"id"`

	Name []string `json:"name"`
}

func (q *Queries) ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error) {

	listAuthors := listAuthors

	if len(arg.ID) <= 0 {
		return nil, fmt.Errorf("arg.ID length is invalid")
	}
	{
		param := "?"
		for i := 0; i < len(arg.ID)-1; i++ {
			param += ",?"
		}
		listAuthors = replaceNth(listAuthors, "(?)", "( "+param+" )", 1)
	}

	if len(arg.Name) <= 0 {
		return nil, fmt.Errorf("arg.Name length is invalid")
	}
	{
		param := "?"
		for i := 0; i < len(arg.Name)-1; i++ {
			param += ",?"
		}
		listAuthors = replaceNth(listAuthors, "(?)", "( "+param+" )", 1)
	}

	rows, err := q.db.QueryContext(ctx, listAuthors, append(append([]interface{}{arg.Bio}, int32Slice2interface(arg.ID)...), stringSlice2interface(arg.Name)...)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CompanyID,
			&i.Size,
			&i.EmptyCol,
			&i.DefaultCol,
			&i.Size1,
			&i.DefaultCol1,
			&i.Type,
			&i.typeOverrides,
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

const listAuthorsOmit = `-- name: ListAuthorsOmit :many
select id, name, bio, company_id, size, empty_col, default_col, size1, default_col1, type, type1 from authors where id=? and name=? or bio=?
`

type ListAuthorsOmitParams struct {
	ID int32 `json:"id"`

	Name string `json:"name"`

	Bio sql.NullString `json:"bio"`
}

func (q *Queries) ListAuthorsOmit(ctx context.Context, arg ListAuthorsOmitParams) ([]Author, error) {

	listAuthorsOmit := listAuthorsOmit

	rows, err := q.db.QueryContext(ctx, listAuthorsOmit, arg.ID, arg.Name, arg.Bio)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Bio,
			&i.CompanyID,
			&i.Size,
			&i.EmptyCol,
			&i.DefaultCol,
			&i.Size1,
			&i.DefaultCol1,
			&i.Type,
			&i.typeOverrides,
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
