package average

import (
	"fmt"
 )

type Calculator interface {
	Input()
	CalculateEvenAvg()
	CalculateOddAvg()
	CalculateAll()
}

type Entities struct {
	sumOfOdd float64 
	sumOfEven float64;
	sumOccurrance int8;
	oddOccurrance int8;
	
}
func (calci *Entities)Input(nums float64){

	if int(nums)%2 == 0 {
			calci.sumOfEven += nums
			calci.sumOccurrance++
		} else {
			calci.sumOfOdd += nums
			calci.oddOccurrance++
		}
}

func (calci *Entities) CalculateEvenAvg() (float64, error) {
	if calci.sumOccurrance == 0 {
		return 0, fmt.Errorf("no even numbers provided")
 	}
	return  calci.sumOfEven / float64(calci.sumOccurrance), nil
}

func (calci *Entities) CalculateOddAvg() (float64, error) {
	if calci.oddOccurrance == 0 {
		return 0, fmt.Errorf("no odd numbers provided")
	}
	return float64(calci.sumOfOdd) / float64(calci.oddOccurrance), nil
}

func (calci *Entities) CalculateAll() (float64, error) {
	totalCount := calci.sumOccurrance + calci.oddOccurrance
	if totalCount == 0 {
		return 0, fmt.Errorf("no numbers provided")
	}
	totalSum := calci.sumOfEven + calci.sumOfOdd // This should sum both even and odd
	return totalSum / float64(totalCount), nil
}
