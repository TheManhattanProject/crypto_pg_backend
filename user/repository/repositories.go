package repository

import userdomain "github.com/TheManhattanProject/crypt_pg_backend/user/domain/user"

type Repositories struct {
	User userdomain.Repository
}

type Caches struct {
	User userdomain.Cache
}
