package database

import (
	"time"

	"github.com/upper/db/v4"
	"startUp/internal/domain"
)

// Function checks if the pointer is not nil and return the Time structure.
// If pointer is nil - function will return the empty Time structure
func getTimeFromTimePtr(t *time.Time) time.Time {
	if t != nil {
		return *t
	} else {
		return time.Time{}
	}
}

// Function checks if the Time is not an empty Time structure and return its reference back.
// If pointer is a zero structure - function will the empty Time
func getTimePtrFromTime(t time.Time) *time.Time {
	empty := time.Time{}
	if t == empty {
		return nil
	} else {
		return &t
	}
}

// Return logical expression with soft deletion statement
func softDelCond(cond db.LogicalExpr, showDeleted bool) db.LogicalExpr {
	if showDeleted {
		return cond
	}
	delCond := db.Cond{"deleted_date IS": nil}
	if cond == nil {
		return delCond
	} else {
		return db.And(cond, delCond)
	}
}

type dbQueryParams struct {
	Page        uint
	PageSize    uint
	ShowDeleted bool
}

func (q *dbQueryParams) ApplyToResult(r db.Result) db.Result {
	if !q.ShowDeleted {
		r = r.And(db.Cond{"deleted_date IS": nil})
	}
	if q.PageSize != 0 {
		r = r.Paginate(q.PageSize).Page(q.Page)
	}
	return r
}

func mapDomainToDbQueryParams(p *domain.UrlQueryParams) *dbQueryParams {
	if p == nil {
		return &dbQueryParams{}
	}
	return &dbQueryParams{
		Page:        p.Page,
		PageSize:    p.PageSize,
		ShowDeleted: p.ShowDeleted,
	}
}
