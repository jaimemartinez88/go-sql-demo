package sqlx

import (
	"github.com/jaimem88/go-sql-demo/pkg/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type SQLXStore struct {
	dbx *sqlx.DB
}

// New creates a SQLXStore instance
func New(connStr string) (*SQLXStore, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db connection")
	}
	return &SQLXStore{
		dbx: db,
	}, nil
}
func (s *SQLXStore) InsertUser(u *types.User) (*types.User, error) {
	rows, err := s.dbx.NamedQuery(`INSERT INTO demo.user (
		name, email, mobile, age, admin
	) VALUES(
		:name, :email, :mobile, :age, :admin
	) RETURNING id`, u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert user")
	}

	for rows.Next() {
		// posgtress does not support LastInsertedId()  https://github.com/lib/pq/issues/24
		if err := rows.Scan(&u.ID); err != nil {
			return nil, errors.Wrap(err, "failed to scan user ID")
		}
	}
	return u, nil
}
func (s *SQLXStore) InsertAddress(addr *types.Address) (*types.Address, error) {
	rows, err := s.dbx.NamedQuery(`INSERT INTO demo.address (
		user_id, street_address, suburb, postcode, state, country
	) VALUES(
		:user_id, :street_address, :suburb, :postcode, :state, :country
	) RETURNING id`, addr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert address")
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&addr.ID); err != nil {
			return nil, errors.Wrap(err, "failed to scan addr ID")
		}
	}
	return addr, nil
}

func (s *SQLXStore) GetUserByEmail(email string) (*types.User, error) {
	u := types.User{}
	err := s.dbx.Get(&u, `SELECT id, name, email, mobile, age, admin 
	FROM demo.user WHERE email = $1 `, email)

	return &u, err
}
func (s *SQLXStore) GetAddressByUserID(userID string) (*types.Address, error) {
	addr := types.Address{}
	err := s.dbx.Get(&addr, `SELECT id, user_id,street_address, suburb, postcode, state, country 
	FROM demo.address WHERE user_id = $1`, userID)
	return &addr, err
}
func (s *SQLXStore) GetAllUsersAndAddresses() ([]*types.UserAddress, error) {
	ua := []*types.UserAddress{}

	err := s.dbx.Select(&ua, `SELECT 
		u.id "user.id", 
		u.name "user.name", 
		u.email "user.email", 
		u.mobile "user.mobile", 
		u.age "user.age", 
		u.admin "user.admin",
		a.id "address.id", 
		a.user_id "address.user_id", 
		a.street_address "address.street_address", 
		a.suburb "address.suburb", 
		a.postcode "address.postcode", 
		a.state "address.state", 
		a.country "address.country"
	FROM demo.user u JOIN demo.address a ON a.user_id = u.id`)
	return ua, err
}
