package tools

import (
	"math"
)

type Paginate struct {
}

/**
获取分页参数
*/
func (p *Paginate) Init(pageSize int, pageIndex int, totalCount int) map[string]int {
	offset := (pageIndex - 1) * pageSize
	totalPage := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	return map[string]int{"total_page": totalPage, "total_count": totalCount, "offset": offset}
}
