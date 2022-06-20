package repository

import "context"

type Visits interface {
	Inc(ctx context.Context) error
	Get(ctx context.Context) (int, error)
}

type Notes interface {
	//Inc(ctx context.Context) error
	GetNote(ctx context.Context, id int) (string, error)
	MakeNote(ctx context.Context, body string) (int, error)
	ChangeNote(ctx context.Context) (int, error)
	DeleteNote(ctx context.Context) (int, error)
}
