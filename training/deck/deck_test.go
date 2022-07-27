package main

import ( 
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d:=newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck leng of 20 but got %v",len(d))
	} 

	if d[0] != "1 of a" {
		t.Errorf("Error %v",d[0])
	}
	if d[len(d)-1] != "4 of d" {
		t.Errorf("Error %v",d[len(d)-1])
	}
}

func TestSaveToDeck (t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	ld := newDeckFromFile("_decktesting")

	if len(ld) != 16 {
		t.Errorf("Expected 16, got %v", len(ld))
	}
	os.Remove("_decktesting")

}