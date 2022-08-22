package utils

import (
	"strconv"
	"strings"
)

type QueryParams struct {
	Filters  map[string]string
	Page     int64
	Limit    int64
	Day      int64
	Ordering []string
	Search   string
	Id       string
	From     string
	To       string
	Token    string
}

func ParseQueryParams(queryParams map[string][]string) (*QueryParams, []string) {
	params := QueryParams{
		Filters:  make(map[string]string),
		Day:      1,
		Page:     1,
		Limit:    10,
		Ordering: []string{},
		Search:   " ",
		Id:       " ",
		From:     " ",
		To:       " ",
		Token:    " ",
	}
	var errStr []string
	var err error

	for key, value := range queryParams {
		if key == "page" {
			params.Page, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `page` param")
			}
			continue
		}

		if key == "author" {
			params.Filters[key] = value[0]
			continue
		}

		if key == "category" {
			params.Filters[key] = value[0]
			continue
		}

		if key == "limit" {
			params.Limit, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `limit` param")
			}
			continue
		}

		if key == "search" {
			params.Search = value[0]
			continue
		}

		if key == "id" {
			params.Id = value[0]
			continue
		}

		if key == "from" {
			params.From = value[0]
			continue
		}

		if key == "to" {
			params.To = value[0]
			continue
		}

		if key == "token" {
			params.Token = value[0]
			continue
		}

		if key == "ordering" {
			params.Ordering = strings.Split(value[0], ",")
			continue
		}
		params.Filters[key] = value[0]

		if key == "day" {
			params.Day, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "Invalid `limit` param")
			}
			continue
		}
		params.Filters[key] = value[0]
	}

	return &params, errStr
}
