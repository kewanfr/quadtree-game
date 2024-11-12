package floor

import (
	"reflect"
	"testing"
)

func TestReadExemple(t *testing.T) {
	var testArr = [][]int{{1, 1, 3, 4}, {1, 1, 4, 3}, {0, 0, 2, 2}, {0, 0, 2, 2}}
	var result = readFloorFromFile("exemple")

	if result == nil {
		t.Error("Erreur: result is nil")
	}

	if !reflect.DeepEqual(result, testArr) {
		t.Errorf("Erreur: expected %v, got %v", testArr, result)
	}
}


func TestReadInvalidFile(t *testing.T) {
	var result = readFloorFromFile("invalid")

	if result != nil {
		t.Error("Erreur: expected nil, got a result")
	}
}