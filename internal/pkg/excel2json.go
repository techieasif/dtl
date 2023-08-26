package pkg

import (
	"log"

	"github.com/FerdinaKusumah/excel2json"
)

func FromEXCELPath(path string, sheetName string) {
	var result []*map[string]interface{}
	var err error
	if result, err = excel2json.GetExcelFilePath(path, sheetName, nil); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	jsonBytes, _ := ReadAsJson(result)
	err = Save2Json(path, jsonBytes)
	if err != nil {
		panic(err)
	}
}

func FromEXCELURL(url string, path2Save string) {
	var result []*map[string]interface{}
	var err error
	if result, err = excel2json.GetExcelFileUrl(url, "", nil); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	jsonBytes, _ := ReadAsJson(result)
	err = Save2Json(path2Save, jsonBytes)
	if err != nil {
		panic(err)
	}
}

func FromCSVPath(path string) {
	var result []*map[string]interface{}
	var err error
	if result, err = excel2json.GetCsvFilePath(path, "", nil); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	jsonBytes, _ := ReadAsJson(result)
	err = Save2Json(path, jsonBytes)
	if err != nil {
		panic(err)
	}
}

func FromCSVURL(url string, path2Save string) {
	var result []*map[string]interface{}
	var err error
	if result, err = excel2json.GetCsvFileUrl(url, "", nil); err != nil {
		log.Fatalf(`unable to parse file, error: %s`, err)
	}
	jsonBytes, _ := ReadAsJson(result)
	err = Save2Json(path2Save+"jsonData", jsonBytes)
	if err != nil {
		panic(err)
	}
}
