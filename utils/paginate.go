package utils

import (
	"gorm.io/gorm"
)

type PaginateParams struct {
	PageParams int
	Limit      int
}

func Paginate(p *PaginateParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := p.PageParams
		if page <= 0 {
			page = 1
		}

		pageSize := p.Limit
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
