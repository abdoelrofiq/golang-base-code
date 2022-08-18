package core

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Query struct {
	key   string
	value string
}

func getURL(c echo.Context) string {
	url := c.Request().URL.String()

	return url
}

func getBracketIndex(value string) (int, int) {
	firstBracketIndex := strings.Index(value, "[")
	secondBracketIndex := strings.Index(value, "]")

	return firstBracketIndex, secondBracketIndex
}

func valueType(value string) string {

	types := map[string]interface{}{
		"INT":    "int",
		"STRING": "string",
		"DATE":   "date",
		"ARRAY":  "array",
	}

	return types[value].(string)

}

func getOperator(value string) (string, error) {
	const emptyOperator = ""

	translateMethods := map[string]interface{}{
		"EQ":     "=",
		"NE":     "!=",
		"IN":     "IN",
		"NOT IN": "NOT IN",
	}

	operator := translateMethods[value]

	if operator != nil {
		return operator.(string), nil
	}

	return emptyOperator, errors.New("operator only supports = and != for now")
}

func filterQueries(urlQuery url.Values) []*Query {
	queries := make([]*Query, 0)

	for key, element := range urlQuery {
		firstBracketIndex, secondBracketIndex := getBracketIndex(key)

		if firstBracketIndex != -1 && secondBracketIndex != -1 {
			query := &Query{
				key:   key,
				value: element[0],
			}

			queries = append(queries, query)
		}

	}

	return queries
}

func filterBuilder(columnName string, operator string, connector string, replacementName string) string {
	return columnName + " " + operator + " " + "@" + replacementName + connector
}

func connectorBuilder(queries []*Query, index int) string {
	var connector string

	if len(queries) == index {
		connector = ""
	} else {
		connector = " AND "
	}

	return connector
}

func valueChecker(q *Query, operator string) interface{} {
	valueChecker := map[string]interface{}{}

	integerValue, integerValueErr := strconv.Atoi(q.value)
	if integerValueErr == nil {
		valueChecker["value"] = integerValue
		valueChecker["valueType"] = valueType("INT")
		return valueChecker
	}

	_, dateValueErr := time.Parse("2006-01-02", q.value)
	if dateValueErr == nil {
		valueChecker["value"] = q.value
		valueChecker["valueType"] = valueType("DATE")
		return valueChecker
	}

	inOperator, _ := getOperator("IN")
	notInOperator, _ := getOperator("NOT IN")
	if operator == inOperator || operator == notInOperator {
		var values []string
		findArrayDelimiter := strings.Index(q.value, ",")
		if findArrayDelimiter != -1 {
			values = strings.Split(q.value, ",")
		} else {
			values = []string{q.value}
		}

		valueChecker["value"] = values
		valueChecker["valueType"] = valueType("ARRAY")
		return valueChecker
	}

	valueChecker["value"] = q.value
	valueChecker["valueType"] = valueType("STRING")
	return valueChecker
}

func columnNameBuilder(q *Query, firstBracketIndex int, argumentValue map[string]interface{}) string {
	var columnName string

	if argumentValue["valueType"] == valueType("DATE") {
		columnName = "DATE_FORMAT(" + q.key[:firstBracketIndex] + ", '%Y-%m-%d')"
		return columnName
	}

	columnName = q.key[:firstBracketIndex]
	return columnName
}

func replacementNameBuilder(q *Query, firstBracketIndex int, index int) string {
	return q.key[:firstBracketIndex] + "." + strconv.Itoa(index)
}

func FilterQueryParser(c echo.Context) (string, interface{}, error) {
	var columnName string
	var operatorAbbreviation string
	var filterQueryString string
	var replacementName string
	filterArgument := map[string]interface{}{}

	url, _ := url.Parse(getURL(c))
	urlQuery := url.Query()
	queries := filterQueries(urlQuery)

	for index, element := range queries {
		firstBracketIndex, secondBracketIndex := getBracketIndex(element.key)

		if firstBracketIndex != -1 && secondBracketIndex != -1 {
			index++

			// create connector for query
			connector := connectorBuilder(queries, index)

			// create operator
			operatorAbbreviation = strings.ToUpper(string(element.key[firstBracketIndex+1 : secondBracketIndex]))
			operator, err := getOperator(operatorAbbreviation)
			if err != nil {
				return filterQueryString, filterArgument, errors.New(err.Error())
			}

			// create argument value
			argumentValue := valueChecker(element, operator).(map[string]interface{})

			// create column name
			columnName = columnNameBuilder(element, firstBracketIndex, argumentValue)

			// create replacement name
			replacementName = replacementNameBuilder(element, firstBracketIndex, index)

			// define filterArgument
			filterArgument[replacementName] = argumentValue["value"]

			// define filterQueryString
			filterQueryString = filterQueryString + filterBuilder(columnName, operator, connector, replacementName)
		}

	}

	return filterQueryString, filterArgument, nil
}
