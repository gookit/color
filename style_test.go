package color

import (
	"testing"
	"regexp"
	"fmt"
	"log"
)

func TestMatchTag(t *testing.T) {
	s1 := "<err>text</>"

	reg := regexp.MustCompile(TagExpr)
	r1 := reg.FindAllStringSubmatch(s1, -1)

	fmt.Printf("ret %+v\n", r1)

	s2 := "abc <err>err-text</> def <info>info text</>"
	reg, err := regexp.Compile(TagExpr)

	if err != nil {
		log.Fatal(err)
	}

	r2 := reg.FindAllStringSubmatch(s2, -1)
	fmt.Printf("ret %v\n", r2)

	s3 := `abc <err>err-text</> 
def <info>info text
</>`
	g3, err := regexp.Compile(TagExpr)

	if err != nil {
		log.Fatal(err)
	}

	r3 := g3.FindAllStringSubmatch(s3, -1)
	fmt.Printf("ret %v\n", r3)

	// some tag and content
	s4 := "abc <err>err-text</> def <err>err-text</> "
	g4, err := regexp.Compile(TagExpr)

	if err != nil {
		log.Fatal(err)
	}

	r4 := g4.FindAllStringSubmatch(s4, -1)
	fmt.Printf("ret %v\n", r4)

}

func TestClearTag(t *testing.T) {
	s1 := "<err>text</>"
	ClearTag(s1)

	s2 := "abc <err>err-text</> def <info>info text</>"
	ClearTag(s2)

	s3 := `abc <err>err-text</> 
def <info>info text
</>`
	ClearTag(s3)
}
