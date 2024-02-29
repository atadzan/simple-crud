package models

import "io"

type PaginationParams struct {
	Offset uint64
	Limit  uint64
}

type FileResponse struct {
	Reader io.Reader
	Size   int64
}
