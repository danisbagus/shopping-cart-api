package user

import (
	"database/sql"
	"fmt"

	"github.com/danisbagus/shopping-cart-api/core/domain"
	"github.com/danisbagus/shopping-cart-api/utils/constant"

	"github.com/jmoiron/sqlx"
)

type (
	Repo struct {
		db *sqlx.DB
	}

	User struct {
		ID        uint64 `db:"id"`
		Name      string `db:"name"`
		Email     string `db:"email"`
		Password  string `db:"password"`
		RoleID    uint64 `db:"role_id"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}
)

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r Repo) FindOneByEmail(email string) (*domain.User, error) {
	var user User
	query := "select id, name, email, password, role_id from users where email = $1"
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID)

	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("[user repo] failed find user:%v", err)
	}

	output := user.mapOut()
	return output, nil
}

func (r Repo) Insert(inData *domain.User) (*domain.User, error) {
	user := toUser(inData)

	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("[user repo] error while beginning trx:%v", err)
	}

	query := `INSERT INTO users(name, email, password, role_id, created_at, updated_at)
				  VALUES($1, $2, $3, $4, $5, $6)
				  RETURNING id`

	var userID uint64
	err = tx.QueryRow(query, user.Name, user.Email, user.Password, user.RoleID, user.CreatedAt, user.UpdatedAt).Scan(&userID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[user repo] error while insert user:%v", err)

	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[user repo] error while commiting trx:%v", err)

	}

	user.ID = userID
	output := user.mapOut()
	return output, nil
}

func toUser(inData *domain.User) *User {
	user := new(User)

	user.ID = inData.ID
	user.Name = inData.Name
	user.Email = inData.Email
	user.Password = inData.Password
	user.RoleID = inData.RoleID
	user.CreatedAt = inData.CreatedAt.Format(constant.DATE_TIME_FORMAT)
	user.UpdatedAt = inData.UpdatedAt.Format(constant.DATE_TIME_FORMAT)

	return user
}

func (u User) mapOut() *domain.User {
	output := new(domain.User)

	output.ID = u.ID
	output.Name = u.Name
	output.Email = u.Email
	output.Password = u.Password
	output.RoleID = u.RoleID

	return output
}
