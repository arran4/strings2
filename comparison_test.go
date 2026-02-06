package strings2

import (
	"fmt"
	"testing"
)

type TestCase struct {
	Input    string
	Expected string
	Format   string
	Provider string
}

var comparisonTests = []TestCase{
	// iancoleman/strcase
	{"testCase", "test_case", "Snake", "iancoleman"},
	{"TestCase", "test_case", "Snake", "iancoleman"},
	{"Test Case", "test_case", "Snake", "iancoleman"},
	{" Test Case", "test_case", "Snake", "iancoleman"},
	{"Test Case ", "test_case", "Snake", "iancoleman"},
	{" Test Case ", "test_case", "Snake", "iancoleman"},
	{"test", "test", "Snake", "iancoleman"},
	{"test_case", "test_case", "Snake", "iancoleman"},
	{"Test", "test", "Snake", "iancoleman"},
	{"", "", "Snake", "iancoleman"},
	{"ManyManyWords", "many_many_words", "Snake", "iancoleman"},
	{"manyManyWords", "many_many_words", "Snake", "iancoleman"},
	{"AnyKind of_string", "any_kind_of_string", "Snake", "iancoleman"},
	{"numbers2and55with000", "numbers_2_and_55_with_000", "Snake", "iancoleman"},
	{"JSONData", "json_data", "Snake", "iancoleman"},
	{"userID", "user_id", "Snake", "iancoleman"},
	{"AAAbbb", "aa_abbb", "Snake", "iancoleman"},
	{"1A2", "1_a_2", "Snake", "iancoleman"},
	{"A1B", "a_1_b", "Snake", "iancoleman"},
	{"A1A2A3", "a_1_a_2_a_3", "Snake", "iancoleman"},
	{"A1 A2 A3", "a_1_a_2_a_3", "Snake", "iancoleman"},
	{"AB1AB2AB3", "ab_1_ab_2_ab_3", "Snake", "iancoleman"},
	{"AB1 AB2 AB3", "ab_1_ab_2_ab_3", "Snake", "iancoleman"},
	{"some string", "some_string", "Snake", "iancoleman"},
	{" some string", "some_string", "Snake", "iancoleman"},

	{"test_case", "TestCase", "Pascal", "iancoleman"}, // ToCamel in iancoleman is Pascal
	{"test.case", "TestCase", "Pascal", "iancoleman"},
	{"test", "Test", "Pascal", "iancoleman"},
	{"TestCase", "TestCase", "Pascal", "iancoleman"},
	{" test  case ", "TestCase", "Pascal", "iancoleman"},
	{"", "", "Pascal", "iancoleman"},
	{"many_many_words", "ManyManyWords", "Pascal", "iancoleman"},
	{"AnyKind of_string", "AnyKindOfString", "Pascal", "iancoleman"},
	{"odd-fix", "OddFix", "Pascal", "iancoleman"},
	{"numbers2And55with000", "Numbers2And55With000", "Pascal", "iancoleman"},
	{"ID", "Id", "Pascal", "iancoleman"},
	{"CONSTANT_CASE", "ConstantCase", "Pascal", "iancoleman"},

	{"foo-bar", "fooBar", "Camel", "iancoleman"}, // ToLowerCamel in iancoleman is Camel
	{"TestCase", "testCase", "Camel", "iancoleman"},
	{"", "", "Camel", "iancoleman"},
	{"AnyKind of_string", "anyKindOfString", "Camel", "iancoleman"},
	{"AnyKind.of-string", "anyKindOfString", "Camel", "iancoleman"},
	{"ID", "id", "Camel", "iancoleman"},
	{"some string", "someString", "Camel", "iancoleman"},
	{" some string", "someString", "Camel", "iancoleman"},
	{"CONSTANT_CASE", "constantCase", "Camel", "iancoleman"},

	{"testCase", "test-case", "Kebab", "iancoleman"},

	// ettle/strcase
	{"Hello world!", "hello_world!", "Snake", "ettle"},
	{"Hello world!", "HELLO_WORLD!", "ScreamingSnake", "ettle"},
	{"Hello world!", "hello-world!", "Kebab", "ettle"},
	{"Hello world!", "HELLO-WORLD!", "ScreamingKebab", "ettle"},
	{"Hello world!", "HelloWorld!", "Pascal", "ettle"},
	{"Hello world!", "helloWorld!", "Camel", "ettle"},
	{"Hello world!", "Hello World!", "Title", "ettle"},

	{"snake_case", "snake_case", "Snake", "ettle"},
	{"SNAKE_CASE", "snake_case", "Snake", "ettle"},
	{"kebab-case", "kebab_case", "Snake", "ettle"},
	{"PascalCase", "pascal_case", "Snake", "ettle"},
	{"camelCase", "camel_case", "Snake", "ettle"},
	{"Title Case", "title_case", "Snake", "ettle"},
	{"point.case", "point_case", "Snake", "ettle"},

	{"ID", "id", "Snake", "ettle"},
	{"ID", "ID", "ScreamingSnake", "ettle"},
	{"userID", "user_id", "Snake", "ettle"},
	{"userID", "USER_ID", "ScreamingSnake", "ettle"},
	{"JSON_blob", "json_blob", "Snake", "ettle"},
	{"HTTPStatusCode", "http_status_code", "Snake", "ettle"},
	{"http200", "http200", "Snake", "ettle"},

	// janos/casbab
	{"camelSnakeKebab", "camel_snake_kebab", "Snake", "janos"},
	{"camelSnakeKebab", "camel-snake-kebab", "Kebab", "janos"},
	{"camelSnakeKebab", "CamelSnakeKebab", "Pascal", "janos"},
	{"camelSnakeKebab", "camelSnakeKebab", "Camel", "janos"},
	{"camelSnakeKebab", "CAMEL_SNAKE_KEBAB", "ScreamingSnake", "janos"},
	{"camelSnakeKebab", "CAMEL-SNAKE-KEBAB", "ScreamingKebab", "janos"},
	{"camelSnakeKebab", "Camel Snake Kebab", "Title", "janos"},

	// golang-cz/textcase
	{"someText", "some_text", "Snake", "golang-cz"},
	{"someText", "someText", "Camel", "golang-cz"},
	{"Add updated_at to users table", "addUpdatedAtToUsersTable", "Camel", "golang-cz"},
	{"Add updated_at to users table", "add_updated_at_to_users_table", "Snake", "golang-cz"},

	// searKing/golang
	{"name____+++2", "Name2", "Pascal", "searKing"}, // UpperCamelCase
	{"one__two_+_+three.four__", "OneTwoThreeFour", "Pascal", "searKing"},
	{"name_2", "name2", "Camel", "searKing"}, // LowerCamelCase
	// {"_my_field_name_2", "xMyFieldName2", "Camel", "searKing"}, // Commented out due to arbitrary 'x' prefix behavior

	// tomedharris/caseconv
	{"foo-", "foo", "Camel", "tomedharris"},
	{"one two", "oneTwo", "Camel", "tomedharris"},
	{"one two", "one_two", "Snake", "tomedharris"},
	{"one two", "OneTwo", "Pascal", "tomedharris"},
	{"one two", "one-two", "Kebab", "tomedharris"},

	// gobeam/stringy
	{"ThisIsOne___messed up string. Can we Really Snake Case It?", "This_Is_One_messed_up_string_Can_we_Really_Snake_Case_It", "Snake", "gobeam"},
    {"ThisIsOne___messed up string. Can we Really camel-case It ?##", "thisIsOneMessedUpStringCanWeReallyCamelCaseIt", "Camel", "gobeam"},
}

func getParseOptions(provider string) []any {
	var opts []any
	if provider == "gobeam" {
		delims := map[rune]bool{
			' ': true, '_': true, '-': true, '?': true, '#': true, '.': true,
		}
		opts = append(opts, NewPartitioner(PartitionerConfig{
			Delimiters: delims,
			SplitCamel: true,
		}))
	} else if provider == "iancoleman" {
		opts = append(opts, WithSmartAcronyms(false), WithNumberSplitting(true))
	} else if provider == "searKing" {
		// searKing seems to strip leading delimiters?
		// "_my_field_name_2" -> "xMyFieldName2" (why x? Maybe empty word mapped to X?)
		// Actually the test said: "_my_field_name_2" -> "xMyFieldName2".
		// Our parser: _, my, field, name, 2.
		// If we use SnakeCase partitioner it might help?
		// For now let's try defaults.
	}
	return opts
}

func getFormatOptions(provider, format string) []Option {
	var opts []Option
	// Defaults for some providers imply strict casing (whispering)
	isSnake := format == "Snake" || format == "Kebab" // Kebab usually follows snake rules for casing
	if isSnake {
		if provider == "iancoleman" || provider == "ettle" || provider == "golang-cz" || provider == "tomedharris" || provider == "janos" {
			opts = append(opts, OptionCaseMode(CMWhispering))
		}
	}
	return opts
}

func TestComparisons(t *testing.T) {
	for i, test := range comparisonTests {
		t.Run(fmt.Sprintf("%d_%s_%s", i, test.Provider, test.Format), func(t *testing.T) {

			parseOpts := getParseOptions(test.Provider)
			words, _ := Parse(test.Input, parseOpts...)

			opts := getFormatOptions(test.Provider, test.Format)
			var got string

			switch test.Format {
			case "Snake":
				if test.Provider == "gobeam" {
					got = ToSnakeCase(words) // Preserve case for gobeam
				} else {
					got = ToSnakeCase(words, opts...)
				}
			case "ScreamingSnake":
				got = ToSnakeCase(words, append(opts, OptionCaseMode(CMScreaming))...)
			case "Kebab":
				got = ToKebabCase(words, opts...)
			case "ScreamingKebab":
				got = ToKebabCase(words, append(opts, OptionCaseMode(CMScreaming))...)
			case "Camel":
				got = ToCamelCase(words, opts...)
			case "Pascal":
				got = ToPascalCase(words, opts...)
			case "Title":
				got = ToFormattedCase(words, append(opts, OptionDelimiter(" "), OptionCaseMode(CMAllTitle))...)
			default:
				t.Fatalf("Unknown format: %s", test.Format)
			}

			if got != test.Expected {
				t.Errorf("Input: %q\nExpected: %q\nGot:      %q", test.Input, test.Expected, got)
			}
		})
	}
}
