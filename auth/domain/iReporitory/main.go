package iRepository

type Repository struct {
	Account AccountInterface
}

type RepositoryInterface interface {
	GetAccount() AccountInterface
}

func (repo *Repository) GetAccount() AccountInterface {
	return repo.Account
}
