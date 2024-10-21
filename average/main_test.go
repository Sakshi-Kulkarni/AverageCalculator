package average

import (
	"testing"
)

func processInput(calci *Entities, inputs []float64) {
	for _, num := range inputs {
		calci.Input(num)
	}
}

func TestInput(t *testing.T) {
	tests := []struct {
		name              string
		input             []float64
		expectedSumEven   float64
		expectedCountEven int8
		expectedSumOdd    float64
		expectedCountOdd  int8
	}{
		{
			name :"Test with even and odd numbers", 
			input  :[]float64{2, 4, 1, 3},
			expectedSumEven: 6, 
			expectedCountEven:2,
			expectedSumOdd: 4,
			expectedCountOdd:   2},
		{
			name :"Test with only even numbers", 
			input  :[]float64{2, 4, 6}, 
			expectedSumEven:12, 
			expectedCountEven:3, 
			expectedSumOdd:0, 
			expectedCountOdd  :0},
		{
			name :"Test with only odd numbers",
			input  : []float64{1, 3, 5}, 
			expectedSumEven:0,
			expectedCountEven: 0, 
			expectedSumOdd:9,
			expectedCountOdd  : 3},
		{
			name :"Test with no input", 
			input  :[]float64{},
			expectedSumEven: 0,
			expectedCountEven: 0, 
			expectedSumOdd:0,
			expectedCountOdd  : 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calci := &Entities{}
			processInput(calci, tt.input)

			if calci.sumOfEven != tt.expectedSumEven {
				t.Errorf("sumOfEven: got %f, want %f", calci.sumOfEven, tt.expectedSumEven)
			}
			if calci.sumOccurrance != tt.expectedCountEven {
				t.Errorf("evenCount: got %d, want %d", calci.sumOccurrance, tt.expectedCountEven)
			}
			if calci.sumOfOdd != tt.expectedSumOdd {
				t.Errorf("sumOfOdd: got %f, want %f", calci.sumOfOdd, tt.expectedSumOdd)
			}
			if calci.oddOccurrance != tt.expectedCountOdd {
				t.Errorf("oddCount: got %d, want %d", calci.oddOccurrance, tt.expectedCountOdd)
			}
		})
	}
}

func TestCalculateEvenAvg(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		expected  float64
		errorMsg string
	}{
		{
			name :"Even average of 2 and 4", 
			input  :[]float64{2, 4}, 
			expected :3, 
			errorMsg:""},
		{
			name :"Even average of 6 and 10", 
			input  :[]float64{6, 10}, 
			expected :8,
			errorMsg :""},
		{
			name :"No even numbers",
			input  : []float64{1, 3, 5},
			expected : 0,
			errorMsg :"no even numbers provided"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calci := &Entities{}
			processInput(calci, tt.input)

			avg, err := calci.CalculateEvenAvg()
			if tt.errorMsg != "" {  
				if err == nil {
					t.Error("Expected an error, but got nil")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error message: '%s', got: '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if avg != tt.expected {
					t.Errorf("got %f, want %f", avg, tt.expected)
				}
			}
		})
	}
}

func TestCalculateOddAvg(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		expected  float64
	    errorMsg string
	}{
		{
			name:      "Odd average of 1 and 3",
			input:     []float64{1, 3},
			expected:  2,
			errorMsg: ""},
		{
			name:      "Odd average of 5 and 9",
			input:     []float64{5, 9},
			expected:  7,
			errorMsg: ""},
		{
			name: "No odd numbers", 
			input: []float64{2, 4, 6}, 
			expected: 0, 
			errorMsg: "no odd numbers provided"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calci := &Entities{}
			processInput(calci, tt.input)

			avg, err := calci.CalculateOddAvg()
			if tt.errorMsg != "" {  
				if err == nil {
					t.Error("Expected an error, but got nil")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error message: '%s', got: '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if avg != tt.expected {
					t.Errorf("got %f, want %f", avg, tt.expected)
				}
			}
 
			
		})
	}
}

func TestCalculateAll(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		expected  float64
		errorMsg string
	}{
		{
			name :"Average of 2, 4, 1, 3", 
			input  :[]float64{2, 4, 1, 3}, 
			expected :2.5,
			errorMsg: ""}, 
		{
			name :"Average of 6, 8, 10",
			input  : []float64{6, 8, 10}, 
			expected :8,
			errorMsg: ""},    
		{
			name :"No numbers", 
			input  :[]float64{},
			expected : 0, 
			errorMsg:"no numbers provided"},            
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calci := &Entities{}
			processInput(calci, tt.input)

			avg, err := calci.CalculateAll()
			if tt.errorMsg != "" {  
				if err == nil {
					t.Error("Expected an error, but got nil")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error message: '%s', got: '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if avg != tt.expected {
					t.Errorf("got %f, want %f", avg, tt.expected)
				}
			}
		})
	}
}


func BenchmarkCalculateOddAvg(b *testing.B) {
	calci := &Entities{}

 	for i := 1; i <= 1000; i += 2 {
		calci.Input(float64(i))
	}

 	for i := 0; i < b.N; i++ {
		_, err := calci.CalculateOddAvg()
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkCalculateEvenAvg(b *testing.B) {
	calci := &Entities{}

 	for i := 1; i <= 1000; i += 2 {
		calci.Input(float64(i))
	}

 	for i := 0; i < b.N; i++ {
		_, err := calci.CalculateOddAvg()
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkCalculateAll(b *testing.B) {
	calci := &Entities{}

 	for i := 1; i <= 1000; i += 2 {
		calci.Input(float64(i))
	}

 	for i := 0; i < b.N; i++ {
		_, err := calci.CalculateOddAvg()
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	}
}