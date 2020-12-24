package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDack()
	dLen := len(d)
	if dLen != 52 {
		t.Errorf("Expected 52 cards in new deck, but got %d", dLen)
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected Spades of Ace at the first position, but got %s", d[0])
	}

	if d[dLen-1] != "Clubs of King" {
		t.Errorf("Expected Clubs of King at the first position, but got %s", d[dLen-1])
	}
}

func TestSaveFileAndLoadDeckFromFile(t *testing.T) {
	d := newDack()
	testFile := "deck_file.test"
	// Clean any old file
	os.Remove(testFile)

	d.saveToFile(testFile)

	dLoadedFromFile := newDeckFromFile(testFile)
	// Verify if deck has correct number of elements
	if len(dLoadedFromFile) != 52 {
		t.Errorf("Expected 52 cards in new deck, but got %d", len(dLoadedFromFile))
	}
	// Cleanup
	err := os.Remove(testFile)
	if err != nil {
		t.Errorf("Unable to clean the %v file", testFile)
	}
}
