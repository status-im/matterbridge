package status

import (
	"testing"
)

func TestGetThreeWordName(t *testing.T) {
	contact := "0x042629a27c72bff4017dc8720e583dc5e39f229e13001c6f30fe722b9214d8ef869db4e0ac40c61456a22ca8868eeeb51e88119d4e8447cd4cc0045f3209cf43ed"
	correct := "meek corrupt koalabear"
	rval := getThreeWordName(contact)
	if (rval != correct) {
		t.Errorf("Incorrect 3-word alias: got: %s, want: %s.", rval, correct)
	}
}
