package morestring

// Test file should be named with _test suffix

import "testing" // Testing packge containing various utilities for tests


type concatTestCase struct {
	input []string;
	expectedOutput string;
}

// Test function should be named with Test prefix
func TestConcatString(t *testing.T) {
	test_cases := []concatTestCase {
		{[]string {"",""}, ""},
		{[]string {"abc","def"}, "abcdef"},
		{[]string {"abcsdfsdfds","deffghfgh"}, "abcsdfsdfdsdeffghfghs"},
		{[]string {"","deffghfgh"}, "deffghfgh"},
		{[]string {"abcsdfsdfds",""}, "abcsdfsdfds"},
	}

	for _, test_case := range test_cases {
		output := ConcatString(test_case.input[0], test_case.input[1])
		if output != test_case.expectedOutput {
			
			// Equivalent to logf followed by fail
			// %q prints out string with a double quote
			t.Errorf("ConcatString(%q, %q) = %q, expected = %q", test_case.input[0], test_case.input[1], output, test_case.expectedOutput)
		}
	}
}
