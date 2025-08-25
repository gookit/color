package color

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkRenderCode_SingleString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderCode("32", "Hello World")
	}
}

func BenchmarkRenderCode_TwoStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderCode("32", "Hello", "World")
	}
}

func BenchmarkRenderCode_UserPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderCode("32", strconv.Itoa(i)+". ", "test")
	}
}

func BenchmarkColorRender_SingleString(b *testing.B) {
	yellow := FgGreen.Render
	for i := 0; i < b.N; i++ {
		yellow("Hello World")
	}
}

func BenchmarkColorRender_TwoStrings(b *testing.B) {
	yellow := FgGreen.Render
	for i := 0; i < b.N; i++ {
		yellow("Hello", "World")
	}
}

func BenchmarkColorRender_UserPattern(b *testing.B) {
	yellow := FgGreen.Render
	for i := 0; i < b.N; i++ {
		yellow(strconv.Itoa(i)+". ", "test")
	}
}

// Baseline comparisons
func BenchmarkFmtSprint_SingleString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprint("Hello World")
	}
}

func BenchmarkFmtSprint_TwoStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprint("Hello", "World")
	}
}

func BenchmarkFmtSprint_UserPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprint(strconv.Itoa(i)+". ", "test")
	}
}