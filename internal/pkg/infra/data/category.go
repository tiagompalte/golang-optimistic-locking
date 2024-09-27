package data

import (
	"context"
	"fmt"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/repository"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	pkgRepo "github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type CategoryRepository struct {
	conn         pkgRepo.Connector
	mainTable    string
	selectFields string
}

func NewCategoryRepository(conn pkgRepo.Connector) repository.CategoryRepository {
	return CategoryRepository{
		conn:      conn,
		mainTable: "tb_category",
		selectFields: `
			id
			, created_at
			, updated_at
			, slug
			, name
			, user_id
			
			FROM tb_category
		`,
	}
}

func (r CategoryRepository) parseEntity(s pkgRepo.Scanner) (entity.Category, error) {
	var slug string
	var category entity.Category
	err := s.Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
		&slug,
		&category.Name,
		&category.UserID,
	)
	if err != nil {
		return entity.Category{}, errors.Repo(err, r.mainTable)
	}
	category.SetSlug(slug)

	return category, nil
}

func (r CategoryRepository) Insert(ctx context.Context, category entity.Category) (uint32, error) {
	res, err := r.conn.ExecContext(ctx,
		"INSERT INTO tb_category (slug, name, user_id) VALUES (?,?,?)",
		category.GetSlug(),
		category.Name,
		category.UserID,
	)
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.Repo(err, r.mainTable)
	}

	return uint32(id), nil
}

func (r CategoryRepository) FindBySlugAndUserID(ctx context.Context, slug string, userID uint32) (entity.Category, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at AND slug = ? AND user_id = ?
	`

	q := fmt.Sprintf(query, r.selectFields)

	category, err := r.parseEntity(
		r.conn.QueryRowContext(
			ctx,
			q,
			slug,
			userID,
		))
	if err != nil {
		return entity.Category{}, errors.Repo(err, r.mainTable)
	}

	return category, nil
}

func (r CategoryRepository) FindByUserID(ctx context.Context, userID uint32) ([]entity.Category, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at AND user_id = ?
	`

	q := fmt.Sprintf(query, r.selectFields)

	result, err := r.conn.QueryContext(
		ctx,
		q,
		userID,
	)

	list, err := pkgRepo.ParseEntities[entity.Category](
		r.parseEntity,
		result,
		err,
	)
	if err != nil {
		return []entity.Category{}, errors.Repo(err, r.mainTable)
	}

	return list, nil
}

func (r CategoryRepository) FindByID(ctx context.Context, id uint32) (entity.Category, error) {
	query := `
		SELECT %s
			WHERE NOT deleted_at AND id = ?
	`

	q := fmt.Sprintf(query, r.selectFields)

	category, err := r.parseEntity(
		r.conn.QueryRowContext(
			ctx,
			q,
			id,
		))
	if err != nil {
		return entity.Category{}, errors.Repo(err, r.mainTable)
	}

	return category, nil
}
