package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Repository struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

type Note struct {
	body string `json:"body"`
}

func (r *Repository) GetNote(ctx context.Context, id int) (string, error) {

	rows, err := r.db.Query(ctx, "SELECT body FROM notes WHERE id = $1", id)
	if err != nil {
		return "", err
	}

	var note string
	rows.Next()
	err = rows.Scan(note)
	if err != nil {
		return "", err
	}

	/*var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return err
		}
	}*/

	return note, rows.Err()
}

func (r *Repository) MakeNote(ctx context.Context, body string) (int, error) {

	_, err := r.db.Exec(context.Background(),
		"insert into notes(body) values($1)", body)

	return 0, err
	/*note := Note{}
	err := ctx.BodyParser(&note)
	if err != nil {
		return err
	}
	err := r.db.Create(&book).Error
	if err != nil {
		return err
	}
	return nil*/
}

func (r *Repository) ChangeNote(ctx context.Context, description string, id int32) error {
	_, err := conn.Exec( "update notes set body=$1 where id=$2", description, id)
	return err
}

func (r *Repository) DeleteNote(ctx context.Context, id int32) error {
	_, err = conn.Exec("delete from notes where id=$1", id)
	
	return err
}

func (r *Repository) Get(ctx context.Context) (int, error) {
	rows, err := r.db.Query(ctx, "SELECT counter FROM visits")
	if err != nil {
		return 0, err
	}
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, rows.Err()
}

func (r *Repository) Inc(ctx context.Context) error {
	current, err := r.Get(ctx)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, "UPDATE visits SET counter = $1", current+1)
	return err
}
