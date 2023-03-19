package emaildomainstats

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

type testCase struct {
	input          string
	expectedOutput string
	title          string
}

var testCases = []testCase{
	{
		title: "Test invalid emails are ignored",
		input: `first_name,last_name,email,gender,ip_address
		FirstName,LastName,one@domain.com,GenderA,0.0.0.0
		FirstName,LastName,@nocustomer.com,GenderA,0.0.0.0
		FirstName,LastName,nodomain@,GenderA,0.0.0.0
		FirstName,LastName,no email,GenderA,0.0.0.0`,
		expectedOutput: "domain.com,1\n",
	},
	{
		title: "Test customers are counted per domain",
		input: `first_name,last_name,email,gender,ip_address
		FirstName,LastName,one@A-domain.com,GenderA,0.0.0.0
		FirstName,LastName,two@A-domain.com,GenderA,0.0.0.0
		FirstName,LastName,one@B-domain.com,GenderA,0.0.0.0
		FirstName,LastName,two@B-domain.com,GenderA,0.0.0.0
		FirstName,LastName,one@C-domain.com,GenderA,0.0.0.0`,
		expectedOutput: "A-domain.com,2\nB-domain.com,2\nC-domain.com,1\n",
	},
	{
		title: "Test duplicate customers are only counted once",
		input: `first_name,last_name,email,gender,ip_address
		FirstName,LastName,one@A-domain.com,GenderA,0.0.0.0
		FirstName,LastName,one@A-domain.com,GenderA,0.0.0.0`,
		expectedOutput: "A-domain.com,1\n",
	},
	{
		title: "Test output is sorted alphabetically",
		input: `first_name,last_name,email,gender,ip_address
		FirstName,LastName,one@Z-domain.com,GenderA,0.0.0.0
		FirstName,LastName,one@A-domain.com,GenderA,0.0.0.0
		FirstName,LastName,one@a-domain.com,GenderA,0.0.0.0
		FirstName,LastName,one@z-domain.com,GenderA,0.0.0.0
		FirstName,LastName,one@B-domain.com,GenderA,0.0.0.0
		FirstName,LastName,one@b-domain.com,GenderA,0.0.0.0`,
		expectedOutput: "A-domain.com,1\na-domain.com,1\nB-domain.com,1\nb-domain.com,1\nZ-domain.com,1\nz-domain.com,1\n",
	},
	{
		title: "Test invalid rows are handled gracefully",
		input: `first_name,last_name,email,gender,ip_address
		FirstName,LastName,one@A-domain.com,GenderA,0.0.0.0
		FirstName,LastName,one@A-domain.com,GenderA`,
		expectedOutput: "A-domain.com,1\n",
	},
}

// TestProcessFileTestCases tests various inputs to the ProcessFile function
func TestProcesFileTestCases(t *testing.T) {
	for _, testCase := range testCases {
		t.Logf("Running test case: %s", testCase.title)
		// Given csv data
		in := strings.NewReader(testCase.input)

		// When the data is processed and written to output
		buf := new(bytes.Buffer)
		// And no errors are expected
		if err := ProcessData(in, buf); err != nil {
			t.Fail()
		}
		// Then the output contains the expected data
		if buf.String() != testCase.expectedOutput {
			t.Errorf("Received dataStats %s does not match expected %s", buf.String(), testCase.expectedOutput)
		}
	}

}

// TestProcessFileEmptyCSV test that an empty CSV dataset produces an error
// when processed
func TestProcessFileEmptyCSV(t *testing.T) {
	// Given empty CSV data
	in := strings.NewReader(``)

	// When the data is processed
	buf := new(bytes.Buffer)
	if err := ProcessData(in, buf); err == nil {
		// Then an error is produces
		t.Fail()
	}

	// And no output is produced
	if buf.String() != "" {
		t.Fail()
	}
}

// BenchmarkTestThousandLineFile ensures the performance of the file is as expected
func BenchmarkTestThousandLineFile(b *testing.B) {
	// Given a file with a thousand CSV lines
	in, err := os.Open("../test_data/customer_data.csv")
	if err != nil {
		b.Error(err)
	}

	// When the data is processed and written to output
	buf := new(bytes.Buffer)
	// Then no errors are expected
	if err := ProcessData(in, buf); err != nil {
		b.Fail()
	}
}
