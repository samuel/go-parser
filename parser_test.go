package parser

import (
	"io"
	"testing"
)

func TestComments(t *testing.T) {
	spec := Spec{
		CommentStart: "/*",
		CommentEnd:   "*/",
		CommentLine:  String("//"),
	}
	in := NewStringInput(`// this is a test
	    // only a test
	    /* this is
	       a multiline comment */`)
	st := &State{Spec: spec, Input: in}
	p := OneLineComment()
	out, d, err := p(st)
	if err != nil {
		t.Fatalf("OneLinecomment returned error %s", err.Error())
	}
	if !d {
		t.Fatal("OneLineComment returned !ok")
	}
	exp := " this is a test"
	if outString, ok := out.(string); !ok {
		t.Fatal("OneLinecomment returned non-string")
	} else if outString != exp {
		t.Fatalf("OneLineComment returned '%s' instead of '%s'", outString, exp)
	}
}

func TestStringInput(t *testing.T) {
	testString := "tes†ing mitä"

	in := NewStringInput(testString)
	outString := ""
	for {
		r, err := in.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatalf("StringInput.Next returned error %s", err.Error())
		}
		in.Pop(1)
		outString += string(r)
	}

	if testString != outString {
		t.Fatalf("Next/Pop produced unmatched output of %#v instead of %#v", outString, testString)
	}

	in = NewStringInput(testString)
	in.Pop(1)
	outString, err := in.Get(5)
	if err != nil {
		t.Fatal("Get(5) returned error %s", err.Error())
	}
	if "es†in" != outString {
		t.Fatalf("Get produced unmatched output of %#v instead of %#v", outString, "es†in")
	}

	in = NewStringInput(testString)
	outString, err = in.Get(12)
	if err != nil {
		t.Fatal("Get(12) returned error %s", err.Error())
	}
	if testString != outString {
		t.Fatalf("Get(len) produced unmatched output of %#v instead of %#v", outString, testString)
	}
	outString, err = in.Get(13)
	if err != io.EOF {
		t.Fatal("Get(len+1) returned error %+v but should have returned EOF", err)
	}
}
