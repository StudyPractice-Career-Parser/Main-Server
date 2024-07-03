package repository

type Search struct {
}

type Repository struct {
	Search
}

func NewRepository() *Repository {
	return &Repository{}
}
