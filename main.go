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
	size, err := queries.GetTotalSize(context.Background(), []int32{1234})
	fmt.Println(size, err)
	sizeNull, err := queries.GetTotalSizeNull(context.Background(), []int32{1234})
	fmt.Println(sizeNull, err)
	sizeNull, err = queries.GetTotalSizeNull(context.Background(), []int32{0, 1})
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
