package storage

import "fmt"

type ReadRequest struct {
	Page         int      `json:"page"`
	PerPage      int      `json:"per_page"`
	Search       string   `json:"search"`
	SearchField  string   `json:"search_field"`
	SearchFields []string `json:"search_fields"`
}

func (r *ReadRequest) Offset() int {
	return (r.Page - 1) * r.PerPage
}

func (r *ReadRequest) Validate(searchFields []string) error {
	if r.Page < 1 {
		return fmt.Errorf("page should be >= 1, found: %d", r.Page)
	}
	if r.PerPage > 100 {
		return fmt.Errorf("per_page should be <= 100, found: %d", r.PerPage)
	}
	if r.Search == "" {
		return nil
	}
	if len(r.Search) > 256 {
		return fmt.Errorf("search length should be less than 256 symbols, found: %d", len(r.Search))
	}
	if r.SearchField != "" {
		if !in(searchFields, r.SearchField) {
			return fmt.Errorf("search_field should be in %v", searchFields)
		}
		return nil
	}
	if !stringSlice(r.SearchFields).in(searchFields) {
		return fmt.Errorf("search_fields should be in %v", searchFields)
	}
	return nil
}

type stringSlice []string

func (ss stringSlice) in(items []string) bool {
	for _, item := range items {
		if !in(ss, item) {
			return false
		}
	}
	return true
}

func in(items []string, s string) bool {
	for _, item := range items {
		if s == item {
			return true
		}
	}
	return false
}
