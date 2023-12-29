package main

import (
	"fmt"
	"net/url"
	"strings"
)

func buildSQLQuery(params url.Values) string {
	sqlQuery := "SELECT * FROM table"

	whereClause := make([]string, 0)
	for key, values := range params {
		if strings.HasPrefix(key, "where") {
			whereClause = append(whereClause, fmt.Sprintf("%s = '%s'", strings.TrimPrefix(key, "where_"), values[0]))
		}
	}

	if len(whereClause) > 0 {
		sqlQuery = fmt.Sprintf("%s WHERE %s", sqlQuery, strings.Join(whereClause, " AND "))
	}

	if sort, ok := params["sort"]; ok {
		sqlQuery = fmt.Sprintf("%s ORDER BY %s", sqlQuery, sort[0])
	}

	return sqlQuery
}

func main() {
	rawURL := "http://example.com/api?where_id=1&where_name=John&sort=id%20DESC"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	params := parsedURL.Query()

	sqlQuery := buildSQLQuery(params)
	fmt.Println(sqlQuery)
}
