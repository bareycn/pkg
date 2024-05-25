package paginator

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Pagination struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	HasNext  bool        `json:"has_next"`
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
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

// Paginator 分页返回数据
func Paginator(c *gin.Context, tx *gorm.DB, data interface{}) *Pagination {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	// 获取总数
	var total int64
	tx.Count(&total)

	return &Pagination{
		Page:     page,
		PageSize: pageSize,
		Data:     data,
		Total:    total,
		HasNext:  total > int64(page*pageSize),
	}
}
