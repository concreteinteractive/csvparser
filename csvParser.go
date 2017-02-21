package parser

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//Parser parse a csv file and returns an array of pointers of the type specified
type Parser interface {
	Parse(resultType interface{})
}

//CsvParser parses a csv file and returns an array of pointers the type specified
type CsvParser struct {
	CsvSeparator        rune
	SkipFirstLine       bool
	SkipEmptyValues     bool
	AllowIncompleteRows bool
}

//Parse creates the array of the given type from the csv file
func (parser CsvParser) Parse(filepath string, f interface{}) ([]interface{}, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parser.ParseWithReader(file, f)
}

func getIndexFromHeaders(headers []string, column string) int {
	if len(headers) == 0 {
		return -1
	}

	for idx, headerText := range headers {
		if column == headerText {
			return idx
		}
	}
	return -1
}

func getCsvColumnIndex(fieldIndex int, field reflect.StructField, headers []string) (int, error) {
	byHeaderIndex := getIndexFromHeaders(headers, strings.ToLower(field.Tag.Get("csvColumn")))
	if byHeaderIndex >= 0 {
		return byHeaderIndex, nil
	}

	csvTag := field.Tag.Get("csv")
	if len(csvTag) == 0 {
		return fieldIndex, nil
	}

	csvColumnIndex, csvTagErr := strconv.Atoi(csvTag)
	if csvTagErr != nil {
		return -1, csvTagErr
	}

	return csvColumnIndex, nil
}

//ParseWithReader creates the array of the given type from the csv file
func (parser CsvParser) ParseWithReader(r io.Reader, f interface{}) ([]interface{}, error) {
	var err error
	var csvReader = csv.NewReader(r)
	csvReader.Comma = parser.CsvSeparator

	var results = make([]interface{}, 0, 0)

	resultType := reflect.ValueOf(f).Type()
	headers := []string{}

	if parser.SkipFirstLine {
		headers, err = csvReader.Read()
		if err != nil {
			headers = []string{}
		}
		for i, h := range headers {
			headers[i] = strings.ToLower(h)
		}
	}

	for {

		rawCSVLine, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		if len(rawCSVLine) == 0 {
			continue
		}

		var newResult = reflect.New(resultType).Interface()

		// set all the struct fields
		for fieldIndex := 0; fieldIndex < resultType.NumField(); fieldIndex++ {
			var currentField = resultType.Field(fieldIndex)
			csvColumnIndex, csvTagErr := getCsvColumnIndex(fieldIndex, currentField, headers)
			if csvTagErr != nil {
				return nil, csvTagErr
			}

			if csvColumnIndex < 0 {
				return nil, fmt.Errorf("csv tag in struct field %v is less than zero", currentField.Name)
			}

			if csvColumnIndex >= len(rawCSVLine) {
				if parser.AllowIncompleteRows {
					break
				} else {
					return nil, fmt.Errorf("Trying to access csv column %v for field %v, but csv has only %v column(s)", csvColumnIndex, currentField.Name, len(rawCSVLine))
				}
			}

			var csvElement = rawCSVLine[csvColumnIndex]
			var settableField = reflect.ValueOf(newResult).Elem().FieldByName(currentField.Name)

			if csvElement == "" && parser.SkipEmptyValues {
				continue
			}

			switch currentField.Type.Name() {

			case "bool":
				var parsedBool, err = strconv.ParseBool(csvElement)
				if err != nil {
					return nil, err
				}
				settableField.SetBool(parsedBool)

			case "uint", "uint8", "uint16", "uint32", "uint64":
				var parsedUint, err = strconv.ParseUint(csvElement, 10, 64)
				if err != nil {
					return nil, err
				}
				settableField.SetUint(uint64(parsedUint))

			case "int", "int32", "int64":
				var parsedInt, err = strconv.Atoi(csvElement)
				if err != nil {
					return nil, err
				}
				settableField.SetInt(int64(parsedInt))

			case "float32":
				var parsedFloat, err = strconv.ParseFloat(csvElement, 32)
				if err != nil {
					return nil, err
				}
				settableField.SetFloat(parsedFloat)

			case "float64":
				var parsedFloat, err = strconv.ParseFloat(csvElement, 64)
				if err != nil {
					return nil, err
				}
				settableField.SetFloat(parsedFloat)

			case "string":
				settableField.SetString(csvElement)

			case "Time":
				var date, err = time.Parse(currentField.Tag.Get("csvDate"), csvElement)
				if err != nil {
					return nil, err
				}
				settableField.Set(reflect.ValueOf(date))
			}
		}

		results = append(results, newResult)
	}
	return results, nil
}
