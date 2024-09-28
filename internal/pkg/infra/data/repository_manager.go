package data

import (
	"github.com/tiagompalte/golang-clean-optimistic-locking/internal/app/repository"
	pkgRepo "github.com/tiagompalte/golang-clean-optimistic-locking/pkg/repository"
)

type RepositoryManager interface {
	User() repository.UserRepository
}

type repo struct {
	user repository.UserRepository
}

func NewRepositoryManager(conn pkgRepo.Connector) RepositoryManager {
	return repo{
		user: NewUserRepository(conn),
	}
}

func (r repo) User() repository.UserRepository {
	return r.user
}
