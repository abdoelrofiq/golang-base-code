package core

import (
	"errors"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type Filter struct {
	FilterQuery    string
	FIlterArgument string
}

func getBracketIndex(value string) (int, int) {
	firstBracketIndex := strings.Index(value, "[")
	secondBracketIndex := strings.Index(value, "]")

	return firstBracketIndex, secondBracketIndex
}

func valueTypeChecker(value string) string {

	types := map[string]interface{}{
		"INT":    "int",
		"STRING": "string",
	}

	if types[value] != nil {
		return types[value].(string)
	}

	return ""

}

func FQPBuilder(c echo.Context) (string, interface{}, error) {
	var replacementName string
	var filterQueryString string
	var valueType string

	filterQuery := c.QueryParam("filter-query")
	rawfilterArgument := c.QueryParam("filter-argument")
	if len([]rune(filterQuery)) == 0 && len([]rune(rawfilterArgument)) == 0 {
		return "", nil, nil
	}

	if len([]rune(filterQuery)) == 0 && len([]rune(rawfilterArgument)) > 0 {
		return "", nil, errors.New("filter-query parameter not found")
	}

	if len([]rune(filterQuery)) > 0 && len([]rune(rawfilterArgument)) == 0 {
		return "", nil, errors.New("filter-argument parameter not found")
	}

	filter := Filter{FilterQuery: filterQuery, FIlterArgument: rawfilterArgument}
	filterArgument := map[string]interface{}{}
	filterQueryString = filter.FilterQuery

	//convert filter argument from string to array
	newFilterArgument := strings.Split(filter.FIlterArgument, "&&")

	for _, element := range newFilterArgument {

		//convert element  from string to array
		newElement := strings.Split(element, "=")

		for _, rawValue := range newElement {
			firstBracketIndex, secondBracketIndex := getBracketIndex(rawValue)

			if firstBracketIndex != -1 && secondBracketIndex != -1 {
				// value type checker
				rawValueType := strings.ToUpper(string(rawValue[firstBracketIndex+1 : secondBracketIndex]))
				valueType = valueTypeChecker(rawValueType)
				if len([]rune(valueType)) == 0 {
					return filterQueryString, filterArgument, errors.New("value type of " + rawValueType + " is not supported for now.")
				}

				// create replacement name
				replacementName = rawValue[1:firstBracketIndex]
			}

			// create argument value
			switch valueType {
			case valueTypeChecker("INT"):
				integerValue, _ := strconv.Atoi(rawValue)
				filterArgument[replacementName] = integerValue
			default:
				// default value type is string
				filterArgument[replacementName] = rawValue
			}

		}

	}

	return filterQueryString, filterArgument, nil
}

func FQP(DB *gorm.DB, c echo.Context) (*gorm.DB, error) {
	var queryDB *gorm.DB

	filterQueryString, filterArgument, err := FQPBuilder(c)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if len([]rune(filterQueryString)) == 0 && filterArgument == nil {
		queryDB = DB
	} else {
		queryDB = DB.Where(filterQueryString, filterArgument)
	}

	return queryDB, nil
}
