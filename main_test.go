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

func TestSaque(t *testing.T) {
    
  got := calculaSaque(180)
  wantobreakfree := register{nota100 : 1,nota50 : 1, nota20:1, nota10:1 } 

  if  got != wantobreakfree {
		t.Errorf("got %+v, wanted %+v", got, wantobreakfree)
	}
}
