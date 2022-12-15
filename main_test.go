package main

import (
	"testing"
)

/**func TestCalculaSaque(t *testing.T) {
	got := calculaSaque(100)
	want := 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
*/

func TestCalculaRestoRetorna5(t *testing.T) {
    
    got := calculaResto(25,20)
    wantobreakfree := 5

    if got != wantobreakfree {
		t.Errorf("got %q, wanted %q", got, wantobreakfree)
	}
}

func TestNumeroDeNotasRetornaInteiro(t *testing.T) {
    
    got := numeroDeNotas(10, 99)
    wantobreakfree := 9

    if got != wantobreakfree {
		t.Errorf("got %d, wanted %d", got, wantobreakfree)
	}
}
