package tapestry

import (
	//"fmt"
	"testing"
)

func TestSharedPrefixLength(t *testing.T) {
	SetDebug(true)

	a := RandomID()
	b := a

	//Debug.Printf("%v", a)
	//Debug.Printf("%v", b)

	if SharedPrefixLength(a,b) != DIGITS {
		t.Errorf("SharedPrefixLength not working for identical IDs")
	}

	b[7] += Digit(1)
	b[7] %= BASE
	//Debug.Printf("%v", a)
	//Debug.Printf("%v", b)

	if SharedPrefixLength(a,b) != 7 {
		t.Errorf("SharedPrefixLength not working for partially identical IDs")
	}

	b[0] += Digit(1)
	b[0] %= BASE
	//Debug.Printf("%v", a)
	//Debug.Printf("%v", b)

	if SharedPrefixLength(a,b) != 0 {
		t.Errorf("SharedPrefixLength not working for IDs with different starts")
	}	
}

func TestCloser(t *testing.T) {
	SetDebug(true)

	a := RandomID()
	b := a
	c := b

	if a.Closer(b,c) {
		t.Errorf("Closer not working for identical IDs")
	}

	c[DIGITS-1] += Digit(1)
	c[DIGITS-1] %= BASE

	if !a.Closer(b,c) {
		t.Errorf("Closer not working for slightly different IDs")
	}

	if a.Closer(c,b) {
		t.Errorf("Closer not working for slightly different IDs")
	}

	b[3] += Digit(1)
	b[3] %= BASE

	if a.Closer(b,c) {
		t.Errorf("Closer not working for significantly different IDs")
	}

	if !a.Closer(c,b) {
		t.Errorf("Closer not working for significantly different IDs")
	}

	c = b 
	if a.Closer(b,c) {
		t.Errorf("Closer not working for identical IDs")
	}
}

func TestBetterChoice(t *testing.T) {
	SetDebug(true)

	id := RandomID()
	first := id
	first[7] -= Digit(2)
	second := id
	second[7] += Digit(2)

	/*Debug.Printf("%v", id)
	Debug.Printf("%v", first)
	Debug.Printf("%v", second)*/

	if id.BetterChoice(first, second) != false {
		t.Errorf("BetterChoice does not choose correct choice. Should be second.")
	}

	first[7] += Digit(3)

	//Debug.Printf("%v", first)
	//Debug.Printf("%v", second)

	if id.BetterChoice(first, second) != true {
		t.Errorf("BetterChoice does not choose correct choice. Should be first.")
	}

	first[0] += Digit(1)
	//Debug.Printf("%v", first)
	//Debug.Printf("%v", second)

	if id.BetterChoice(first, second) != false {
		t.Errorf("BetterChoice does not choose correct choice. Should be second.")
	}
}
