package modules

import (
	"testing"
)

// TestGroup ...
func TestFindIndex( t *testing.T)  {
	log.Noticef("Start Test Group...")
	g := new(Group)
	i:=g.FindIndex("abd")
	if i!=0{
		t.Errorf( "FindGroup Error %d",i)
	}
	
	i=g.FindIndex("abcd")
	if i!=1{
		t.Errorf( "FindGroup Error %d",i)
	}

	i=g.FindIndex("abd")
	if i!=0{
		t.Errorf( "FindGroup Error %d",i)
	}

	if g.Len() != 2{
		t.Error( "Err Length")
	}
}
