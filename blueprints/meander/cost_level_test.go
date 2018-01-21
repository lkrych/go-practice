package meander_test

//by separating the tests into their own package we can ensure that we are truly testing
//the code like a user. We may only call exported methods and only have visibility to exported
//types. We cannot mess around with internals and do things that our users cannot.

import (
	"meander"
	"testing"

	"github.com/cheekybits/is"
)

func TestCostValues(t *testing.T) {
	is := is.New(t)
	is.Equal(int(meander.Cost1), 1)
	is.Equal(int(meander.Cost2), 2)
	is.Equal(int(meander.Cost3), 3)
	is.Equal(int(meander.Cost4), 4)
	is.Equal(int(meander.Cost5), 5)
}

func TestCostString(t *testing.T) {
	is := is.New(t)
	is.Equal(meander.Cost1.String(), "$")
	is.Equal(meander.Cost2.String(), "$$")
	is.Equal(meander.Cost3.String(), "$$$")
	is.Equal(meander.Cost4.String(), "$$$$")
	is.Equal(meander.Cost5.String(), "$$$$$")
}

func TestParseCost(t *testing.T) {
	is := is.New(t)
	is.Equal(meander.ParseCost("$"), meander.Cost1)
	is.Equal(meander.ParseCost("$$"), meander.Cost2)
	is.Equal(meander.ParseCost("$$$"), meander.Cost3)
	is.Equal(meander.ParseCost("$$$$"), meander.Cost4)
	is.Equal(meander.ParseCost("$$$$$"), meander.Cost5)
}

func TestParseCostRange(t *testing.T) {
	is := is.New(t)
	var l meander.CostRange
	var err error
	l, err = meander.ParseCostRange("$$...$$$")
	is.NoErr(err)
	is.Equal(l.From, meander.Cost2)
	is.Equal(l.To, meander.Cost3)
	l, err = meander.ParseCostRange("$...$$$$$")
	is.NoErr(err)
	is.Equal(l.From, meander.Cost1)
	is.Equal(l.To, meander.Cost5)
}

func TestCostRangeString(t *testing.T) {
	is := is.New(t)
	r := meander.CostRange{
		From: meander.Cost2,
		To:   meander.Cost4,
	}
	is.Equal("$$...$$$$", r.String())
}
