package detector

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
	"regexp"
)

func TestIsPotentialDelimiter(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{
			byte('a'),
			false,
		},
		{
			byte('A'),
			false,
		},
		{
			byte('1'),
			false,
		},
		{
			byte('|'),
			true,
		},
		{
			byte('$'),
			true,
		},
	}

	detector := &detector{
		nonDelimiterRegex : regexp.MustCompile(nonDelimiterRegexString),
	}
	for _, test := range tests {
		assert.Equal(t, test.expected, !detector.nonDelimiterRegex.MatchString(string(test.input)))
	}
}

func TestFrequencyTable(t *testing.T) {
	ft := createFrequencyTable()

	ft.increment(',', 1).increment(',', 2).increment('|', 3).increment('|', 3)

	assert.Equal(t, 1, ft[','][1])
	assert.Equal(t, 1, ft[','][2])
	assert.Equal(t, 2, ft['|'][3])
}

func TestDetectDelimiter(t *testing.T) {
	detector := New()

	file, err := os.OpenFile("./Fixtures/test1.csv", os.O_RDONLY, os.ModePerm)
	assert.NoError(t, err)
	defer file.Close()

	delimiters := detector.DetectDelimiter(file, '"')

	assert.Equal(t, []string{","}, delimiters)
}

func TestDetectorSample(t *testing.T) {
	detector := &detector{}

	file, err := os.OpenFile("./Fixtures/test1.csv", os.O_RDONLY, os.ModePerm)
	assert.NoError(t, err)
	defer file.Close()

	actual, line := detector.sample(file, 15, '"')
	expected := frequencyTable{
		'.' : map[int]int{
			2: 1,
			3: 1,
			4: 1,
			5: 1,
		},
		' ' : map[int]int{
			5: 1,
		},
		',' : map[int]int{
			1: 4,
			2: 4,
			3: 4,
			4: 4,
			5: 4,
		},
	}

	for k, v := range expected {
		assert.Equal(t, v, actual[k])
	}
	assert.Equal(t, 5, line)
}

func TestDetectAnalyze(t *testing.T) {
	ft := frequencyTable{
		'.' : map[int]int{
			2: 1,
			3: 1,
			4: 1,
			5: 1,
		},
		' ' : map[int]int{
			5: 1,
		},
		',' : map[int]int{
			1: 4,
			2: 4,
			3: 4,
			4: 4,
			5: 4,
		},
	}

	detector := &detector{}
	candidates := detector.analyze(ft, 5)

	assert.Equal(t, []byte{','}, candidates)
}
