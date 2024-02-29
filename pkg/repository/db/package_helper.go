package db

import (
	"fmt"

	"github.com/atadzan/simple-crud/pkg/models"
)

func addStorageDomain(imgPath *string, domain string) {
	*imgPath = fmt.Sprintf("%s/image/%s", domain, *imgPath)
	return
}

func getFiltersConditionQuery(filters models.BookFilter) (result string) {
	switch {
	case filters.AuthorId != 0 && filters.GenreId != 0:
		result = fmt.Sprintf("AND b.genre_id=%[1]d AND b.author_id=%[2]d", filters.GenreId, filters.AuthorId)
	case filters.GenreId != 0:
		result = fmt.Sprintf("AND b.genre_id=%d", filters.GenreId)
	case filters.AuthorId != 0:
		result = fmt.Sprintf("AND b.author_id=%d", filters.AuthorId)
	}
	return
}
