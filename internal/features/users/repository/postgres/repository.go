package users_postgres_repository

import core_postgres_pool "github.com/fenzit/golang-todoapp/internal/core/repository/postgres/pool"

type UsersReposity struct {
	pool core_postgres_pool.Pool
}

func NewUsersRepository(
	pool core_postgres_pool.Pool,
) *UsersReposity {
	return &UsersReposity{
		pool: pool,
	}
}
