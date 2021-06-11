package main

import (
	"database/sql"
	"fmt"
	"log"
	"sqlc/gen"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := gen.New(db)

	authors, err := queries.GetAuthorsInCompany(context.Background(), gen.GetAuthorsInCompanyParams{
		ID:   []int32{0, 1},
		Name: []string{"a"},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(authors)
	size, err := queries.GetTotalSize(context.Background(), []int32{1234789699})
	fmt.Println("id not exist if null ", size, err)
	sizeNull64In, err := queries.GetTotalSizeNullIn(context.Background(), []int32{1234789699})
	fmt.Println("id not exist sizeNull64In ", sizeNull64In, err)
	sizeNull32In, err := queries.GetTotalSize1NullIn(context.Background(), []int32{1234789699})
	fmt.Println("id not exist sizeNull32In", sizeNull32In, err)
	sizeNull64, err := queries.GetTotalSizeNull(context.Background(), 1234789699)
	fmt.Println("id not exist sizeNull64 ", sizeNull64, err)
	sizeNull32, err := queries.GetTotalSize1Null(context.Background(), 1234789699)
	fmt.Println("id not exist sizeNull32", sizeNull32, err)
	ret1, err := queries.GetMax_default_col(context.Background(), 1234789699)
	//0 sql: Scan error on column index 0, name "max(default_col)": converting NULL to int32 is unsupported
	fmt.Println("GetMax_default_col", ret1, err)
	ret2, err := queries.GetMax_default_col1(context.Background(), 1234789699)
	//GetMax_default_col1 id not exist ,exist value 0 sql: Scan error on column index 0, name "max(default_col1)": converting NULL to int32 is unsupported
	fmt.Println("GetMax_default_col1 id not exist ,exist value", ret2, err)
	ret1, err = queries.GetMax_default_col(context.Background(), 1)
	fmt.Println("GetMax_default_col ID exist", ret1, err)
	ret, err := queries.GetMax_empty_col(context.Background(), 1234789699)
	fmt.Println("GetMax_empty_col", ret, err)
	sizeNull, err := queries.GetTotalSizeNullIn(context.Background(), []int32{0, 1})
	fmt.Println(sizeNull, err)
	size, err = queries.GetTotalSize(context.Background(), []int32{0, 1})
	fmt.Println(size, err)

	authors, err = queries.ListAuthors(context.Background(), gen.ListAuthorsParams{
		ID:   []int32{4, 12, 10},
		Name: []string{"Brian Kernighan"},
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		log.Fatal("ListAuthors error:", err)
	}
	fmt.Println(authors)

	maxID, err := queries.GetMaxID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	if _, err := queries.CreateAuthor(context.Background(), gen.CreateAuthorParams{
		ID:   1 + int32(maxID),
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	}); err != nil {
		log.Fatal("CreateAuthor error:", err)
	}

	fetchedAuthor, err := queries.GetOneAuthor(context.Background(), gen.GetOneAuthorParams{
		ID:        []int32{4, 12, 10},
		Name:      []string{"Brian Kernighan"},
		Bio:       sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
		CompanyID: []int32{1, 2},
	})
	if err != nil {
		log.Fatal("GetOneAuthor error:", err)
	}
	fmt.Println(fetchedAuthor)

	err = queries.DeleteAuthor(context.Background(), fetchedAuthor.ID)
	if err != nil {
		log.Fatal("DeleteAuthor error:", err)
	}
}
