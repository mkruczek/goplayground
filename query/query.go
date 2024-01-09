package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func buildSQLQuery(params url.Values, operator string) (string, []interface{}, error) {
	sqlQuery := "SELECT * FROM table"
	var args []interface{}

	whereClause := make([]string, 0)
	for key, values := range params {
		if strings.HasPrefix(key, "where") {
			if values[0] == "true" || values[0] == "false" {
				whereClause = append(whereClause, fmt.Sprintf("%s IS ?", strings.TrimPrefix(key, "where_")))
			} else if values[0] != "true" && values[0] != "false" {
				return "", nil, errors.New("Invalid boolean value")
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

	return sqlQuery, args, nil
}

func main() {
	rawURL := "http://example.com/api?where_id=1&where_name=John&where_active=true&sort=id%20DESC"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return
	}

	params := parsedURL.Query()

	log.Println("Building SQL query...")
	sqlQuery, args, err := buildSQLQuery(params, "OR")
	if err != nil {
		log.Println("Error building SQL query:", err)
		return
	}
	log.Println("SQL query built successfully")

	fmt.Println(sqlQuery)
	fmt.Println(args)

	// case with error
	rawURL = "http://example.com/api?where_id=1&where_name=John&where_active=not_a_bool&sort=id%20DESC"
	parsedURL, err = url.Parse(rawURL)
	if err != nil {
		log.Println("Error parsing URL:", err)
		return
	}

	params = parsedURL.Query()

	log.Println("Building SQL query...")
	sqlQuery, args, err = buildSQLQuery(params, "OR")
	if err != nil {
		log.Println("Error building SQL query:", err)
		return
	}
	log.Println("SQL query built successfully")

	fmt.Println(sqlQuery)
	fmt.Println(args)
}
