package util

import "testing"

func TestCaculateDbAndTableIndex(t *testing.T) {
	expectDbIndex, expectTableIndex := 0, 1
	if dbIndex, tableIndex := CaculateDbAndTableIndex(3);
		expectDbIndex != dbIndex || expectTableIndex != tableIndex {
		t.Error("expectDbIndex =", expectDbIndex, "dbIndex =", dbIndex,
			"expectTableIndex =", expectTableIndex, "tableIndex =", tableIndex)
	}
}

func TestCaculateDbAndTable(t *testing.T) {
	InitDbs()
	expectTable := "relation_1"
	if _, table := CaculateDbAndTable(3);
		expectTable != table {
		t.Error("expectTable =", expectTable, "table = ", table)
	}
}
