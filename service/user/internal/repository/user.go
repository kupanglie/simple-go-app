package repository

import (
	"context"
	"database/sql"

	"github.com/kupanglie/simple-go-app/service/user/helper"
	"github.com/kupanglie/simple-go-app/service/user/internal/entity"
)

type user struct {
	db *sql.DB
}

func NewUserRP(db *sql.DB) *user {
	return &user{
		db: db,
	}
}

func (u *user) Add(ctx context.Context, name string) (entity.User, error) {
	var user entity.User

	tx, err := u.db.Begin()
	if err != nil {
		return user, err
	}

	query := `INSERT INTO user (id, name, created_at) VALUES (UUID_TO_BIN(UUID()), ?, NOW())`
	_, err = tx.ExecContext(ctx, query, name)
	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return user, err
	}

	return user, nil
}

func (u *user) Edit(ctx context.Context, id, name string) (entity.User, error) {
	var user entity.User

	tx, err := u.db.Begin()
	if err != nil {
		return user, err
	}

	query := `UPDATE user SET name = ?, updated_at = NOW() WHERE id = UUID_TO_BIN(?) AND deleted_at IS NULL`
	_, err = tx.ExecContext(ctx, query, name, id)
	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return user, err
	}

	return user, nil
}

func (u *user) Find(ctx context.Context, id string) (entity.User, error) {
	var user entity.User

	query := `SELECT BIN_TO_UUID(id) as id, name, created_at, updated_at, deleted_at FROM user WHERE id = UUID_TO_BIN(?)`
	rows, err := u.db.QueryContext(ctx, query, id)
	if err != nil {
		return user, err
	}

	defer rows.Close()

	for rows.Next() {
		var createdAt, updatedAt, deletedAt sql.NullString
		err = rows.Scan(&user.ID, &user.Name, &createdAt, &updatedAt, &deletedAt)
		if err != nil {
			return user, err
		}

		user.CreatedAt = helper.CastSQLTimeToTime(createdAt.String)
		user.UpdatedAt = helper.CastSQLTimeToTime(updatedAt.String)
		user.DeletedAt = helper.CastSQLTimeToTime(deletedAt.String)
	}

	return user, nil
}

func (u *user) Delete(ctx context.Context, id string) (entity.User, error) {
	var user entity.User

	tx, err := u.db.Begin()
	if err != nil {
		return user, err
	}

	query := `UPDATE user SET deleted_at = NOW() WHERE id=UUID_TO_BIN(?)`
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return user, err
	}

	return user, nil
}
