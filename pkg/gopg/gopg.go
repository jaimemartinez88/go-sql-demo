package gopg

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/jaimemartinez88/go-sql-benchmark/pkg/types"
)

type GoPGStore struct {
	db *pg.DB
}

// New creates a SQLXStore instance
func New(user, pass, dbName, host, port string) (*GoPGStore, error) {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: pass,
		Database: dbName,
		Addr:     fmt.Sprintf("%s:%s", host, port),
	})

	return &GoPGStore{
		db: db,
	}, nil
}

func (s *GoPGStore) InsertUser(u *types.User) error {
	return nil
}
func (s *GoPGStore) InsertAddress(*types.Address) error {
	return nil
}
func (s *GoPGStore) GetUser(id string) (*types.User, error) {
	return nil, nil
}
func (s *GoPGStore) GetAddress(id string) (*types.Address, error) {
	return nil, nil
}
func (s *GoPGStore) GetAllUsersAndAddresses() ([]*types.UserAddress, error) {
	return nil, nil
}