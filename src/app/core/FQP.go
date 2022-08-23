package core

import (
	"errors"
	"fmt"
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

func valueTypeList(value string) string {

	types := map[string]interface{}{
		"INT":     "int",
		"STRING":  "string",
		"DATE":    "date",
		"ARRAY":   "array",
		"BOOLEAN": "boolean",
	}

	if types[value] != nil {
		return types[value].(string)
	}

	return ""

}

func extractValueType(value string) string {
	firstBracketIndex, secondBracketIndex := getBracketIndex(value)
	return strings.ToUpper(string(value[firstBracketIndex+1 : secondBracketIndex]))
}

func valueTypeChecker(value string) (string, error) {
	rawValueType := extractValueType(value)
	valueType := valueTypeList(rawValueType)
	if len([]rune(valueType)) == 0 {
		return valueType, errors.New(fmt.Sprint("value type of ", rawValueType, " is not supported for now."))
	}

	return valueType, nil
}

func valueConverter(value string, valueType string) interface{} {
	var newValue interface{}

	switch valueType {
	case valueTypeList("INT"):
		integerValue, _ := strconv.Atoi(value)
		newValue = integerValue
	case valueTypeList("DATE"):
		// value type of date will be string always
		newValue = value
	case valueTypeList("ARRAY"):
		newValue = arrayValueBuilder(value)
	case valueTypeList("BOOLEAN"):
		booleanValue, _ := strconv.ParseBool(strings.TrimSpace(value))
		newValue = booleanValue
	default:
		// default value type is string
		newValue = value
	}

	return newValue
}

func arrayValueBuilder(value string) interface{} {
	firstBracketIndex, secondBracketIndex := getBracketIndex(value)
	if firstBracketIndex == 0 && secondBracketIndex > 0 {
		valueType, _ := valueTypeChecker(value)
		if valueType != valueTypeList("ARRAY") {
			var arrayValue []interface{}
			for _, element := range strings.Split(value[secondBracketIndex+2:len([]rune(value))-1], ",") {
				arrayValue = append(arrayValue, valueConverter(element, valueType))
			}

			return arrayValue
		}
	}

	return nil
}

func argumentValueBuilder(element []string, valueType string) (interface{}, error) {
	value := element[1]

	return valueConverter(value, valueType), nil
}

func replacementNameBuilder(element []string) (string, string, error) {
	var valueType string
	rawReplacementName := element[0]
	firstBracketIndex, secondBracketIndex := getBracketIndex(rawReplacementName)

	if firstBracketIndex > 0 && secondBracketIndex > 0 {
		// value type checker
		valueType, err := valueTypeChecker(rawReplacementName)
		if err != nil {
			return rawReplacementName, valueType, errors.New(err.Error())
		}

		// create replacement name
		return rawReplacementName[1:firstBracketIndex], valueType, nil
	}

	return rawReplacementName, valueType, errors.New("failed to create replacement name")
}

func FQPBuilder(c echo.Context) (string, interface{}, error) {
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
	filterQueryString := filter.FilterQuery

	//convert filter argument from string to array
	filterArgumentInArray := strings.Split(filter.FIlterArgument, "&&")

	for _, element := range filterArgumentInArray {

		//convert element  from string to array
		rawArgument := strings.Split(element, "=")

		replacementName, valueType, err := replacementNameBuilder(rawArgument)
		if err != nil {
			return filterQueryString, filterArgument, errors.New(err.Error())
		}

		argumentValue, err := argumentValueBuilder(rawArgument, valueType)
		if err != nil {
			return filterQueryString, filterArgument, errors.New(err.Error())
		}

		// create argument value
		filterArgument[replacementName] = argumentValue

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
