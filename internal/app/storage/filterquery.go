package storage

import (
	"strings"

	"gorm.io/gorm"
)

type filterQuery struct {
	q *gorm.DB
}

func newFilterQuery(q *gorm.DB, r *ReadRequest) *filterQuery {
	q = q.Limit(r.PerPage).Offset(r.Offset())
	f := &filterQuery{q}
	if r.Search == "" {
		return f
	}
	var order orderByBuilder
	if r.SearchField != "" {
		return f.searchOne(&order, r.Search, r.SearchField)
	}
	return f.searchMany(&order, r.Search, r.SearchFields)
}

func (f *filterQuery) searchOne(order *orderByBuilder, search, field string) *filterQuery {
	f.q.Where(`word_similarity(?,`+field+`) > 0.3`, search)
	order.Add(`word_similarity(?,`+field+`) DESC`, search)
	order.ToQuery(f.q)
	return f
}

func (f *filterQuery) searchMany(order *orderByBuilder, search string, fields []string) *filterQuery {
	var whereCond, orderCond strings.Builder
	args := make([]interface{}, 0, len(fields))

	whereCond.WriteString("1=1")
	orderCond.WriteString("0")
	for _, field := range fields {
		whereCond.WriteString(` OR word_similarity(?,` + field + `) > 0.3`)
		orderCond.WriteString(`+` + `word_similarity(?,` + field + `)`)
		args = append(args, search)
	}
	f.q.Where(whereCond.String(), args...)
	order.Add(orderCond.String()+` DESC`, args...)

	order.ToQuery(f.q)
	return f
}
