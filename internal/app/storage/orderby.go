package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderByBuilder struct {
	expr clause.Expr
}

func (o *orderByBuilder) Clause() clause.OrderBy {
	return clause.OrderBy{Expression: o.expr}
}

func (o *orderByBuilder) Add(sql string, args ...interface{}) {
	if sql == "" {
		return
	}
	if o.expr.SQL == "" {
		o.expr.SQL = sql
	} else {
		o.expr.SQL += "," + sql
	}
	o.expr.Vars = append(o.expr.Vars, args...)
}

// ToQuery adds 'order by' clause to q.
func (o *orderByBuilder) ToQuery(q *gorm.DB) *gorm.DB {
	if o.expr.SQL == "" {
		return q
	}
	q.Clauses(o.Clause())
	return q
}
