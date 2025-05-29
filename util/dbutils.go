package util

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func SafeTime(ts pgtype.Timestamp) time.Time {
	if ts.Valid {
		return ts.Time
	}
	return time.Time{}
}
