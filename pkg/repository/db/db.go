package db

import (
	"context"
	"fmt"

	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/go-errors/errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	authorsTable = "authors"
	genresTable  = "genres"
	booksTable   = "books"
)

type Book interface {
	GetGenres(ctx context.Context) ([]models.Genre, error)
	GetAll(ctx context.Context, paginationParams models.BooksParams) ([]models.BookList, error)
	GetById(ctx context.Context, id int) (models.BookDetailed, error)
	Create(ctx context.Context, params models.CreateBookParams) error
	Update(ctx context.Context, params models.UpdateBookParams) error
	Delete(ctx context.Context, id, authorId uint32) error
	Search(ctx context.Context, params models.SearchParams) ([]models.BookList, error)
}

type Auth interface {
	Register(ctx context.Context, params models.AuthParams) error
	GetAuthorId(ctx context.Context, params models.AuthParams) (uint32, error)
}

type DB interface {
	Book
	Auth
}
type db struct {
	dbClient      *pgxpool.Pool
	storageDomain string
}

func New(dbClient *pgxpool.Pool, domain string) *db {
	return &db{
		dbClient:      dbClient,
		storageDomain: domain,
	}
}

func (d *db) Register(ctx context.Context, params models.AuthParams) error {
	query := fmt.Sprintf(`INSERT INTO %s(username, password_hash) VALUES($1, $2)`, authorsTable)
	row, err := d.dbClient.Exec(ctx, query, params.Username, params.Password)
	if err != nil {
		return errors.New(err)
	}
	if row.RowsAffected() == 0 {
		return errors.New("operation failed")
	}
	return nil
}

func (d *db) GetAuthorId(ctx context.Context, params models.AuthParams) (authorID uint32, err error) {
	query := fmt.Sprintf(`SELECT id FROM %s WHERE username=$1 AND password_hash=$2`, authorsTable)
	if err = d.dbClient.QueryRow(ctx, query, params.Username, params.Password).Scan(&authorID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = errors.New(ErrNotFound)
		} else {
			err = errors.New(err)
		}
		return
	}
	return
}

func (d *db) GetGenres(ctx context.Context) (genres []models.Genre, err error) {
	query := fmt.Sprintf(`SELECT id, title FROM %s ORDER BY id`, genresTable)
	rows, err := d.dbClient.Query(ctx, query)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = errors.New(ErrNotFound)

		} else {
			err = errors.New(err)
		}
		return
	}
	defer rows.Close()

	for rows.Next() {
		var genre models.Genre

		if err = rows.Scan(&genre.Id, &genre.Title); err != nil {
			err = errors.New(err)
			return
		}

		genres = append(genres, genre)
	}
	return
}

func (d *db) GetAll(ctx context.Context, params models.BooksParams) (books []models.BookList, err error) {
	books = make([]models.BookList, 0)
	query := fmt.Sprintf(`
							SELECT 
								b.id,
								b.author_id,
								a.username,
								b.genre_id,
								g.title,
								b.title,
								b.image_path,
								b.created_at
							FROM %s b
							JOIN %[2]s a ON a.id = b.author_id
							JOIN %[3]s g ON b.genre_id = g.id
							WHERE %[4]s
							ORDER BY b.created_at %[5]s
							OFFSET $1 LIMIT $2`, booksTable, authorsTable, genresTable,
		getFiltersConditionQuery(params.Filter), params.CreatedAtSort)
	rows, err := d.dbClient.Query(ctx, query, params.Pagination.Offset, params.Pagination.Limit)
	if err != nil {
		err = errors.New(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var book models.BookList
		if err = rows.Scan(
			&book.ID,
			&book.Author.Id,
			&book.Author.Name,
			&book.Genre.Id,
			&book.Genre.Title,
			&book.Title,
			&book.ImgURL,
			&book.CreatedAt,
		); err != nil {
			err = errors.New(err)
			return
		}
		addStorageDomain(&book.ImgURL, d.storageDomain)
		books = append(books, book)
	}
	return
}

func (d *db) GetById(ctx context.Context, id int) (book models.BookDetailed, err error) {
	query := fmt.Sprintf(`
						SELECT 
								b.id,
								a.username,
								b.genre_id,
								g.title,
								b.title,
								b.description,
								b.image_path,
								b.created_at
							FROM %[1]s b
							JOIN %[2]s a ON a.id = b.author_id
							JOIN %[3]s g ON b.genre_id = g.id 
							WHERE b.id=$1`, booksTable, authorsTable, genresTable)
	if err = d.dbClient.QueryRow(ctx, query, id).Scan(
		&book.ID,
		&book.Author,
		&book.Genre.Id,
		&book.Genre.Title,
		&book.Title,
		&book.Description,
		&book.ImgURL,
		&book.CreatedAt,
	); err != nil {
		err = errors.New(err)
		return
	}
	addStorageDomain(&book.ImgURL, d.storageDomain)
	return
}

func (d *db) Create(ctx context.Context, params models.CreateBookParams) error {
	query := fmt.Sprintf("INSERT INTO %s (genre_id, title, description, author_id) VALUES($1, $2, $3, $4)", booksTable)
	row, err := d.dbClient.Exec(ctx, query, params.GenreID, params.Title, params.Description, params.AuthorId)
	if err != nil {
		return errors.New(err)
	}
	if row.RowsAffected() == 0 {
		return errors.New("operation failed")
	}
	return nil
}

func (d *db) Update(ctx context.Context, params models.UpdateBookParams) error {
	query := fmt.Sprintf(`UPDATE %s SET genre_id=$1, title=$2, description=$3 WHERE id=$4 AND author_id=$5`, booksTable)
	row, err := d.dbClient.Exec(ctx, query, params.GenreID, params.Title, params.Description, params.BookId, params.AuthorId)
	if err != nil {
		return errors.New(err)
	}
	if row.RowsAffected() == 0 {
		return errors.New("operation failed")
	}
	return nil
}

func (d *db) Delete(ctx context.Context, id, authorId uint32) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 AND author_id=$2`, booksTable)
	row, err := d.dbClient.Exec(ctx, query, id, authorId)
	if err != nil {
		return errors.New(err)
	}
	if row.RowsAffected() == 0 {
		return errors.New("operation failed")
	}
	return nil
}

func (d *db) Search(ctx context.Context, params models.SearchParams) (books []models.BookList, err error) {
	books = make([]models.BookList, 0)
	query := fmt.Sprintf(`
							SELECT 
								b.id,
								b.author_id,
								a.username,
								b.genre_id,
								g.title,
								b.title,
								b.image_path,
								b.created_at
							FROM %s b
							JOIN %[2]s a ON a.id = b.author_id
							JOIN %[3]s g ON b.genre_id = g.id
							WHERE b.title like '%%%[4]s%%' %[5]s
							ORDER BY b.created_at %[6]s
							OFFSET $1 LIMIT $2`, booksTable, authorsTable, genresTable, params.SearchWord,
		getFiltersConditionQuery(params.Filter), params.CreatedAtSort)
	rows, err := d.dbClient.Query(ctx, query, params.Pagination.Offset, params.Pagination.Limit)
	if err != nil {
		err = errors.New(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var book models.BookList
		if err = rows.Scan(
			&book.ID,
			&book.Author.Id,
			&book.Author.Name,
			&book.Genre.Id,
			&book.Genre.Title,
			&book.Title,
			&book.ImgURL,
			&book.CreatedAt,
		); err != nil {
			err = errors.New(err)
			return
		}
		addStorageDomain(&book.ImgURL, d.storageDomain)
		books = append(books, book)
	}
	return
	return
}
