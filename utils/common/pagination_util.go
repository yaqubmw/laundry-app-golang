package common

import (
	"enigma-laundry-apps/model/dto"
	"math"
	"os"
	"strconv"
)

func GetPaginationParams(params dto.PaginationParam) dto.PaginationQuery {
	// err := LoadEnv()
	// exceptions.CheckErr(err)

	var (
		page, take, skip int
	)

	if params.Page > 0 {
		page = params.Page
	} else {
		page = 1
	}

	if params.Limit == 0 {
		n, _ := strconv.Atoi(os.Getenv("DEFAULT_ROWS_PER_PAGE"))
		take = n
	} else {
		take = params.Limit
	}

	if page > 0 {
		skip = (page - 1) * take
	} else {
		skip = 0
	}

	return dto.PaginationQuery{
		Page: page,
		Take: take,
		Skip: skip,
	}
}

func Paginate(page, limit, totalRows int) dto.Paging {
	return dto.Paging{
		Page:        page,
		RowsPerPage: limit,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(limit))),
	}
}

// rumus offset / rumus pagination
// product => 10 | page 1 s.d 5
// product => 10 | page 6 s.d 10
// offset = (page - 1) * limit
// offset = (1 - 1) * 5 === 0
// offset = (2 - 1) * 5 === 5
// SELECT * FROM product LIMIT 5 OFFSET 5
