package repository

import (
	"go-fiber-starter/app/models"
	"go-fiber-starter/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClientRepository struct{}

func (r *ClientRepository) filterClauses(keyword string) (clauses []clause.Expression) {
	if keyword != "" {
		vars := setKeywordVarsByTotalExpr(keyword, 3)
		query := lowerLikeQuery("first_name") + " OR " + lowerLikeQuery("last_name") + " OR " + lowerLikeQuery("mobile_number")
		clauses = append(clauses, clause.Expr{SQL: query, Vars: vars, WithoutParentheses: true})
	}

	return clauses
}

func (r *ClientRepository) GetAll(pagination utils.Pagination) (*utils.Pagination, error) {
	var items []models.Client
	clauses := r.filterClauses(pagination.Keyword)
	filter := filterPaginate(items, &pagination, clauses)
	if err := DB.Scopes(filter).Find(&items).Error; err != nil {
		return nil, err
	}

	pagination.Rows = items
	return &pagination, nil
}

func (r *ClientRepository) InsertMany(tx *gorm.DB, clients []models.Client, batchSize int) error {
	if err := DB.CreateInBatches(&clients, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func (r *ClientRepository) GetListIDs(tx *gorm.DB) []int {
	var clients []models.Client
	ids := make([]int, 0)
	if err := tx.Select("id").Find(&clients).Error; err != nil {
		return ids
	}

	for _, c := range clients {
		ids = append(ids, c.ID)
	}

	return ids
}

func (r *ClientRepository) FindByGUID(guid string) (models.Client, error) {
	var m models.Client
	if err := DB.First(&m, "uuid = ?", guid).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *ClientRepository) UpdateByGUID(tx *gorm.DB, guid string, storeData models.Client) (models.Client, error) {
	var m models.Client
	if err := tx.Model(&m).Where("uuid = ?", guid).Updates(&storeData).Error; err != nil {
		return m, err
	}

	return m, nil
}

func (r *ClientRepository) DeleteByGUID(guid string) error {
	var m models.Client
	if err := DB.Where("uuid = ?", guid).Delete(&m).Error; err != nil {
		return err
	}

	return nil
}
