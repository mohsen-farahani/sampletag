package tag_test

import (
	"testing"

	"github.com/mohsen-farahani/sampletag/internal/tag"
)

func TestItConvertStringToArrayByAndSeparated(t *testing.T) {
	//arrange
	str := " a, b , c"

	//act
	//tag is sut
	result := tag.Convertor(str)

	//assert
	expected := []string{"a", "b", "c"}
	if len(result) != len(expected) {
		t.Errorf("result is %v, expected is %v", result, expected)
	}
}

func TestItConvertStringToArrayByPipeSeparated(t *testing.T) {
	//arrange
	str := "a | b | c"

	//act
	//tag is sut
	result := tag.Convertor(str)

	//assert
	expected := []string{"a", "b", "c"}
	if len(result) != len(expected) {
		t.Errorf("result is %v, expected is %v", result, expected)
	}
}
