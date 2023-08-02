package dto

// untuk paging di parameter
type PaginationParam struct {
	Page   int
	Offset int
	Limit  int
}

// untuk paging di return
type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

// untuk ditaruh di response
type Paging struct {
	Page        int `json:"page"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalRows   int `json:"totalRows"`
	TotalPages  int `json:"totalPages"`
}
