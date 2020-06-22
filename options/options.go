package options

import "fmt"

type Options struct {
	Filters    []Filter
	Pagination *Pagination
	Sorting    *Sorting
}

type Filter struct {
	By       string
	Value    interface{}
	Operator string // e.g : LIKE , = , > , < etc
}
type Sorting struct {
}
type Pagination struct {
}

func ParseOptionsToSQLQuery(opts *Options) (query string, args []interface{}) {
	if opts == nil {
		return "", nil
	}
	if opts.Filters != nil || len(opts.Filters) > 0 {
		query = " WHERE "
		for _, ft := range opts.Filters {
			if ft.Operator == "LIKE" {
				ft.Value = fmt.Sprint("%", ft.Value, "%")
			}
			query += fmt.Sprint(ft.By, " ", ft.Operator, " ? AND")
			args = append(args, ft.Value)
		}
		query = query[len(query)-3:]
	}
	return
}
