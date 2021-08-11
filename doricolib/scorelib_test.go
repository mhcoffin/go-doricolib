package doricolib

import (
	"io/ioutil"
	"testing"
)

func TestUnmarshall(t *testing.T) {
	bytes, err := ioutil.ReadFile("/Users/michaelcoffin/Downloads/ParseTest1.doricolib")
	if err != nil {
		t.Fatalf("failed to read file: %s", err)
	}
	scorelib, err := ReadXml(bytes)
	if err != nil {
		t.Fatalf("failed to unmarshall XML: %s", err)
	}
	emaps := scorelib.ExpressionMaps.Entities.Contents
	if len(emaps) != 1 {
		t.Fatalf("number of expression maps is incorrect")
	}
	if emaps[0].Name != "ParseTest1" {
		t.Fatalf("wrong name")
	}
	combos := emaps[0].Combinations.Combos
	if len(combos) != 2 {
		t.Fatalf("wrong number of combos")
	}
}
