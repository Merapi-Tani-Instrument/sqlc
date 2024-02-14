// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package booktest

import (
	"context"
	"database/sql"
	"strings"
	"time"
)

const booksByTags = `-- name: BooksByTags :many
SELECT
  book_id,
  title,
  name,
  isbn,
  tag
FROM books
LEFT JOIN authors ON books.author_id = authors.author_id
WHERE tag IN (/*SLICE:tags*/?)
`

type BooksByTagsRow struct {
	BookID int64
	Title  string
	Name   sql.NullString
	Isbn   string
	Tag    string
}

func (q *Queries) BooksByTags(ctx context.Context, tags []string) ([]BooksByTagsRow, error) {
	query := booksByTags
	var queryParams []interface{}
	if len(tags) > 0 {
		for _, v := range tags {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:tags*/?", strings.Repeat(",?", len(tags))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:tags*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BooksByTagsRow
	for rows.Next() {
		var i BooksByTagsRow
		if err := rows.Scan(
			&i.BookID,
			&i.Title,
			&i.Name,
			&i.Isbn,
			&i.Tag,
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

const booksByTitleYear = `-- name: BooksByTitleYear :many
SELECT book_id, author_id, isbn, book_type, title, yr, available, tag FROM books
WHERE title = ? AND yr = ?
`

type BooksByTitleYearParams struct {
	Title string
	Yr    int64
}

func (q *Queries) BooksByTitleYear(ctx context.Context, arg BooksByTitleYearParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, booksByTitleYear, arg.Title, arg.Yr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.AuthorID,
			&i.Isbn,
			&i.BookType,
			&i.Title,
			&i.Yr,
			&i.Available,
			&i.Tag,
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

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO authors (name) VALUES (?)
RETURNING author_id, name
`

func (q *Queries) CreateAuthor(ctx context.Context, name string) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, name)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name)
	return i, err
}

const createBook = `-- name: CreateBook :one
INSERT INTO books (
    author_id,
    isbn,
    book_type,
    title,
    yr,
    available,
    tag
) VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING book_id, author_id, isbn, book_type, title, yr, available, tag
`

type CreateBookParams struct {
	AuthorID  int64
	Isbn      string
	BookType  string
	Title     string
	Yr        int64
	Available time.Time
	Tag       string
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.AuthorID,
		arg.Isbn,
		arg.BookType,
		arg.Title,
		arg.Yr,
		arg.Available,
		arg.Tag,
	)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Yr,
		&i.Available,
		&i.Tag,
	)
	return i, err
}

const deleteAuthorBeforeYear = `-- name: DeleteAuthorBeforeYear :exec
DELETE FROM books
WHERE yr < ? AND author_id = ?
`

type DeleteAuthorBeforeYearParams struct {
	Yr       int64
	AuthorID int64
}

func (q *Queries) DeleteAuthorBeforeYear(ctx context.Context, arg DeleteAuthorBeforeYearParams) error {
	_, err := q.db.ExecContext(ctx, deleteAuthorBeforeYear, arg.Yr, arg.AuthorID)
	return err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE book_id = ?
`

func (q *Queries) DeleteBook(ctx context.Context, bookID int64) error {
	_, err := q.db.ExecContext(ctx, deleteBook, bookID)
	return err
}

const getAuthor = `-- name: GetAuthor :one
SELECT author_id, name FROM authors
WHERE author_id = ?
`

func (q *Queries) GetAuthor(ctx context.Context, authorID int64) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthor, authorID)
	var i Author
	err := row.Scan(&i.AuthorID, &i.Name)
	return i, err
}

const getBook = `-- name: GetBook :one
SELECT book_id, author_id, isbn, book_type, title, yr, available, tag FROM books
WHERE book_id = ?
`

func (q *Queries) GetBook(ctx context.Context, bookID int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, bookID)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.AuthorID,
		&i.Isbn,
		&i.BookType,
		&i.Title,
		&i.Yr,
		&i.Available,
		&i.Tag,
	)
	return i, err
}

const updateBook = `-- name: UpdateBook :exec
UPDATE books
SET title = ?1, tag = ?2
WHERE book_id = ?3
`

type UpdateBookParams struct {
	Title  string
	Tag    string
	BookID int64
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.ExecContext(ctx, updateBook, arg.Title, arg.Tag, arg.BookID)
	return err
}

const updateBookISBN = `-- name: UpdateBookISBN :exec
UPDATE books
SET title = ?1, tag = ?2, isbn = ?4
WHERE book_id = ?3
`

type UpdateBookISBNParams struct {
	Title  string
	Tag    string
	BookID int64
	Isbn   string
}

func (q *Queries) UpdateBookISBN(ctx context.Context, arg UpdateBookISBNParams) error {
	_, err := q.db.ExecContext(ctx, updateBookISBN,
		arg.Title,
		arg.Tag,
		arg.BookID,
		arg.Isbn,
	)
	return err
}