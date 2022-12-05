package repository

import (
	"reflect"
	"testing"
)

func TestSliceToMap(t *testing.T) {

	slice := []string{
		"Илья",
		"324",
		"Кукуруза",
	}

	res, err := sliceToMap(slice)

	if err != nil {
		t.Error(err)
	}

	testData := map[string]bool{
		"Илья":     true,
		"324":      true,
		"Кукуруза": true,
	}

	if !reflect.DeepEqual(res, testData) {
		t.Log("Не прошел тест на сравнение с оригиналом")
		t.Fail()
	}
}

func TestMapToSlice(t *testing.T) {

	mapData := map[string]bool{
		"Илья":     true,
		"324":      true,
		"Кукуруза": true,
	}
	res, err := mapToSlice(mapData)

	if err != nil {
		t.Error(err)
	}

	testData := []string{
		"Илья",
		"324",
		"Кукуруза",
	}

	if !reflect.DeepEqual(res, testData) {
		t.Log("Не прошел тест на сравнение с оригиналом")
		t.Fail()
	}
}
