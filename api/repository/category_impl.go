package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
)

type categoryRepo struct {
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepo{}
}

var structureCategory string = `user_id,name_category,description, icon_name, icon_color`                            //5
var structureCategoryStore string = `user_id,name_category,description, icon_name, icon_color,created_at,updated_at` //7
var structureCategoryUpdate string = `user_id,name_category,description, icon_name, icon_color,updated_at`           //6

func (c *categoryRepo) Store(ctx context.Context, tx *sql.Tx, category domain.Category) error {
	category.CreatedAt = time.Now().Local()
	category.UpdatedAt = category.CreatedAt
	sql := `INSERT INTO categories(
		` + structureCategoryStore + `)
		VALUES ($1,$2,$3,$4,$5,$6,$7)`

	_, err := tx.ExecContext(ctx, sql,
		category.UserId,
		category.NameCategory,
		category.Description,
		category.IconName,
		category.IconColor,
		category.CreatedAt,
		category.UpdatedAt,
	)
	helper.PanicIfError(err)

	return nil
}

func (c *categoryRepo) GetByID(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	sql := `SELECT ` + structureCategory + ` FROM categories WHERE id=$1 AND deleted_at IS NOT NULL`

	rows, err := tx.QueryContext(ctx, sql, id)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(
			&category.UserId,
			&category.NameCategory,
			&category.Description,
			&category.IconName,
			&category.IconColor,
		)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("Category Not Found")
	}
}

func (c *categoryRepo) Update(ctx context.Context, tx *sql.Tx, id int, category domain.Category) error {
	category.UpdatedAt = time.Now().Local()
	sql := `UPDATE INTO categories(
		` + structureCategoryUpdate + `)
		VALUES ($1,$2,$3,$4,$5,$6) WHERE id=$7`

	_, err := tx.ExecContext(ctx, sql,
		category.UserId,
		category.NameCategory,
		category.Description,
		category.IconName,
		category.IconColor,
		category.CreatedAt,
		category.UpdatedAt,
		id,
	)
	helper.PanicIfError(err)

	return nil
}

func (c *categoryRepo) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	deleteAt := time.Now().Local()
	sql := `UPDATE INTO categories (deleted_at)
		values ($1) WHERE id=$2`

	_, err := tx.ExecContext(
		ctx,
		sql,
		deleteAt,
		id)

	helper.PanicIfError(err)

	return nil
}

func (c *categoryRepo) ListByUser(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Category, error) {
	sql := `SELECT ` + structureCategory + ` FROM categories WHERE user_id=$1 AND deleted_at IS NOT NULL`

	rows, err := tx.QueryContext(ctx, sql, userId)
	helper.PanicIfError(err)

	if rows.Next() == false {
		return nil, errors.New("User Id Not Found")
	}

	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(
			&category.UserId,
			&category.NameCategory,
			&category.Description,
			&category.IconName,
			&category.IconColor,
		)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories, nil
}
