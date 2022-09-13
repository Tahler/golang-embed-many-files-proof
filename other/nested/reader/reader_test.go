package reader

import (
	"encoding/json"
	"testing"
)

func TestReadTable(t *testing.T) {
	b, err := ReadTable()
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]string
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	got := m["name"]
	want := "Samples"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
