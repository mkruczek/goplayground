package main

import (
	"fmt"
	"net/url"
	"strings"
)

func buildSQLQuery(params url.Values, operator string) (string, []any) {
	sqlQuery := "SELECT * FROM table"
	var args []any

	whereClause := make([]string, 0)
	for key, values := range params {
		if strings.HasPrefix(key, "where") {
			if values[0] == "true" || values[0] == "false" {
				whereClause = append(whereClause, fmt.Sprintf("%s IS ?", strings.TrimPrefix(key, "where_")))
			} else {
				whereClause = append(whereClause, fmt.Sprintf("%s = ?", strings.TrimPrefix(key, "where_")))
			}
			args = append(args, values[0])
		}
	}

	if len(whereClause) > 0 {
		sqlQuery = fmt.Sprintf("%s WHERE %s", sqlQuery, strings.Join(whereClause, " "+operator+" "))
	}

	if sort, ok := params["sort"]; ok {
		sqlQuery = fmt.Sprintf("%s ORDER BY %s", sqlQuery, sort[0])
	}

	return sqlQuery, args
}

func main() {
	rawURL := "http://example.com/api?where_id=1&where_name=John&where_active=true&sort=id%20DESC"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	params := parsedURL.Query()

	sqlQuery, args := buildSQLQuery(params, "OR")
	fmt.Println(sqlQuery)
	fmt.Println(args)
}
