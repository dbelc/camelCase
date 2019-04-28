package stringutils

import (
	"github.com/dbelc/camelCase/dictionary"
	"strings"
	"testing"
)

type subtest struct {
	name   string
	testFn func(*testing.T)
}

func TestCamelCase(t *testing.T) {
	for _, subtest := range []subtest{
		{"Empty", empty},
		{"SingleWord", singleWord},
		{"MultiWord", multiWord},
		{"NonWords", nonWords},
	} {
		t.Run(subtest.name, subtest.testFn)
	}
}

// Test helpers

type testCase struct {
	in, want string
}

func testCamelCase(t *testing.T, cases []testCase, assert func(testCase, string, error)) {
	for _, c := range cases {
		got, err := CamelCase(c.in, testDictInst)
		assert(c, got, err)
	}
}

func testCamelCaseSuccess(t *testing.T, cases []testCase) {
	testCamelCase(t, cases, func(c testCase, got string, err error) {
		format := "CamelCase(%q) == %q, want %q"
		if err != nil {
			t.Errorf(format, c.in, err.Error(), c.want)
		} else if got != c.want {
			t.Errorf(format, c.in, got, c.want)
		}
	})
}

func testCamelCaseError(t *testing.T, cases []testCase) {
	testCamelCase(t, cases, func(c testCase, got string, err error) {
		if err == nil {
			t.Errorf("CamelCase(%q) == %q, did not return error, want InvalidArgument", c.in, got)
		}
	})
}

// Test cases

func empty(t *testing.T) {
	testCamelCaseSuccess(t, []testCase{
		// Empty string should be empty string
		{"", ""},
	})
}

func singleWord(t *testing.T) {
	testCamelCaseSuccess(t, []testCase{
		// Single words should return as single words, lowercased
		{"word", "word"},
		{"Title", "title"},
		{"CAPITAL", "capital"},
		{"mIxEd", "mixed"},
	})
}

func multiWord(t *testing.T) {
	testCamelCaseSuccess(t, []testCase{
		// Multi words should return as camelCased
		{"onetwo", "oneTwo"},
		{"ThreeFour", "threeFour"},
		{"FIVESIX", "fiveSix"},
		{"sEvEnEiGhT", "sevenEight"},

		{"chaircouchsink", "chairCouchSink"},
		{"GuitarLampRemote", "guitarLampRemote"},
		{"VACUUMTREEGRILLFOOD", "vacuumTreeGrillFood"},

		{"thisisasentence", "thisIsASentence"},
	})
}

func nonWords(t *testing.T) {
	testCamelCaseError(t, []testCase{
		// If the entire sentence is nonwords
		{"awefawefaawefa", ""},
		{"oioiwefwoinomxqeq", ""},

		// If the sentence starts with nonwords and contains valid word(s)
		{"asdfawefasdhello", ""},
		{"fawefourfawoitwowoeif", ""},

		// If the sentence starts with words and contains nonwords
		{"phoneaweofiwef", ""},
		{"testCaseOfawaefiaoweif", ""},
	})
}

// Create a test dictionary for the test cases
type testDictionary struct {
	cache map[string]bool
}

func (dict *testDictionary) IsWord(str string) bool {
	lower := strings.ToLower(str)
	if _, prs := dict.cache[lower]; prs {
		return true
	}

	return false
}

func newTestDictionary(init ...string) *testDictionary {
	cache := make(map[string]bool)
	for _, word := range init {
		cache[strings.ToLower(word)] = true
	}

	return &testDictionary{cache}
}

var testDictInst dictionary.Dictionary = newTestDictionary(
	"word",
	"title",
	"capital",
	"mixed",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"chair",
	"couch",
	"sink",
	"guitar",
	"lamp",
	"remote",
	"vacuum",
	"tree",
	"grill",
	"food",
	"no",
	"hello",
	"four",
	"phone",
	"test",
	"case",
	"this",
	"is",
	"a",
	"sentence",
)
