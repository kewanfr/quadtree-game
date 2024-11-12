package floor

import (
	"reflect"
	"testing"
)

func TestReadExemple(t *testing.T) {
	var want = [][]int{{1, 1, 3, 4}, {1, 1, 4, 3}, {0, 0, 2, 2}, {0, 0, 2, 2}}
	var result = readFloorFromFile("../floor-files/exemple")

	if result == nil {
		t.Error("Erreur: result is nil")
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Erreur: expected %v, got %v", want, result)
	}
}

func TestReadEmpty(t *testing.T) {
	var want [][]int = make([][]int, 0)
	var result = readFloorFromFile("../tests/emptyFile")

	if result == nil {
		t.Error("Erreur: result is nil")
	}

	if !reflect.DeepEqual(result, want) {
		t.Errorf("Erreur: expected %v, got %v", want, result)
	}
}


func TestReadInvalidFile(t *testing.T) {
	var result = readFloorFromFile("invalid")

	if result != nil {
		t.Error("Erreur: expected nil, got a result")
	}
}