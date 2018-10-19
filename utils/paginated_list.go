package util

import (
	"net/http"
	"strconv"
)

type Cursor interface {
	Cursor() string
}

type Node struct {
	Cursor string
	Node   interface{}
}

type PageInfo struct {
	HasNextPage     bool
	HasPreviousPage bool
}

type PaginationQuery struct {
	First  int64
	After  string
	Before string
}

type PageinationResults struct {
	PageInfo PageInfo
	Edges    []Node
}

func PaginationQuery(req *http.Request) *PaginationQuery {
	args := req.URL.Query()
	first, err := strconv.Atoi(args.Get("first"))
	if err != nil {
		first = 100
	}
	return &PaginationQuery{
		First:  first,
		After:  args.Get("after"),
		Before: args.Get("after"),
	}
}

func MakePaginationResults(list []Node) *PageinationResults {
	pg := PageInfo{
		HasNextPage:     true,
		HasPreviousPage: true,
	}
	return &PageinationResults{
		PageInfo: pg,
		Edges:    list,
	}
}
