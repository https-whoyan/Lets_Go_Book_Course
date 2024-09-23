package errors

import (
	"errors"
	"github.com/lib/pq"
)

func CheckIsSQLUniqueConstrainsError(err error) bool {
	var sqlErr *pq.Error
	if errors.As(err, &sqlErr) {
		if sqlErr.Code == "23505" {
			return true
		}
	}
	return false
}
