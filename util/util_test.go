package util

import "testing"

func TestGetShortId(t *testing.T){
	shortId, err := GetShortId()
	if shortId == "" || err != nil {
		t.Error("GetShortId failed!")
	}
	t.Log("GetShortId test pass")
}

func BenchmarkGetShortId(b *testing.B){
	for i:=0; i <b.N; i++{
		GetShortId()
	}
}

func BenchmarkGetShortIdTimeConsuming(b *testing.B){
	b.StopTimer()
	shortId, err := GetShortId()
	if shortId == "" || err != nil {
		b.Error(err)
	}
	b.StartTimer()

	for i:=0; i < b.N; i++{
		GetShortId()
	}


}
