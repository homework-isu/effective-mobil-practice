package flagparser

import (
	"flag"
	"fmt"
)

const (
	ActionUp     = "up"
	ActionDown   = "down"
	ActionActual = "actual"
)

type Parser struct{}

func (Parser) GetAction() (string, int, error) {
	upFlag := flag.Int(ActionUp, 0, "steps to update database")
	downFlag := flag.Int(ActionDown, 0, "steps to down database")
	actualFlag := flag.Bool(ActionActual, false, "update database to actual state")

	flag.Parse()

	initsCount := 0 
	if *upFlag != 0 {
		initsCount++
	}
	if *downFlag != 0 {
		initsCount++
	}
	if *actualFlag {
		initsCount++
	}

	if initsCount == 0 {
		return "", 0, fmt.Errorf("must be init with one of up, down, actual")
	}
	if initsCount > 1 {
		return "", 0, fmt.Errorf("must be init with only one flag")
	}

	if *upFlag != 0 {
		return ActionUp, *upFlag, nil
	}
	if *downFlag != 0 {
		return ActionDown, *downFlag, nil
	}
	if *actualFlag{
		return ActionActual, 1, nil
	}
	
	return "", 0, fmt.Errorf("must be init with one of up, down, actual")
}
