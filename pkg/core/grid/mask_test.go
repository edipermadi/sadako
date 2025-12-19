package grid_test

import (
	"testing"

	"github.com/edipermadi/sadako/pkg/core/cell"
	"github.com/edipermadi/sadako/pkg/core/grid"
	"github.com/stretchr/testify/require"
)

func TestMask_Set(t *testing.T) {
	type testCase struct {
		Title         string
		GivenValue    cell.Value
		ExpectedValue grid.Mask
	}

	testCases := []testCase{
		{
			Title:         "One",
			GivenValue:    cell.ValueOne,
			ExpectedValue: grid.Mask(0b10),
		},
		{
			Title:         "Two",
			GivenValue:    cell.ValueTwo,
			ExpectedValue: grid.Mask(0b100),
		},
		{
			Title:         "Three",
			GivenValue:    cell.ValueThree,
			ExpectedValue: grid.Mask(0b1000),
		},
		{
			Title:         "Four",
			GivenValue:    cell.ValueFour,
			ExpectedValue: grid.Mask(0b10000),
		},
		{
			Title:         "Five",
			GivenValue:    cell.ValueFive,
			ExpectedValue: grid.Mask(0b100000),
		},
		{
			Title:         "Six",
			GivenValue:    cell.ValueSix,
			ExpectedValue: grid.Mask(0b1000000),
		},
		{
			Title:         "Seven",
			GivenValue:    cell.ValueSeven,
			ExpectedValue: grid.Mask(0b10000000),
		},
		{
			Title:         "Eight",
			GivenValue:    cell.ValueEight,
			ExpectedValue: grid.Mask(0b100000000),
		},
		{
			Title:         "Nine",
			GivenValue:    cell.ValueNine,
			ExpectedValue: grid.Mask(0b1000000000),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			var mask grid.Mask
			require.Equal(t, tc.ExpectedValue, mask.Set(tc.GivenValue))
		})
	}
}

func TestMask_IsSet(t *testing.T) {
	type testCase struct {
		Title         string
		GivenMask     grid.Mask
		GivenValue    cell.Value
		ExpectedValue bool
	}

	testCases := []testCase{
		{
			Title:         "One",
			GivenMask:     grid.Mask(0b10),
			GivenValue:    cell.ValueOne,
			ExpectedValue: true,
		},
		{
			Title:         "Two",
			GivenMask:     grid.Mask(0b100),
			GivenValue:    cell.ValueTwo,
			ExpectedValue: true,
		},
		{
			Title:         "Three",
			GivenMask:     grid.Mask(0b1000),
			GivenValue:    cell.ValueThree,
			ExpectedValue: true,
		},
		{
			Title:         "Four",
			GivenMask:     grid.Mask(0b10000),
			GivenValue:    cell.ValueFour,
			ExpectedValue: true,
		},
		{
			Title:         "Five",
			GivenMask:     grid.Mask(0b100000),
			GivenValue:    cell.ValueFive,
			ExpectedValue: true,
		},
		{
			Title:         "Six",
			GivenMask:     grid.Mask(0b1000000),
			GivenValue:    cell.ValueSix,
			ExpectedValue: true,
		},
		{
			Title:         "Seven",
			GivenMask:     grid.Mask(0b10000000),
			GivenValue:    cell.ValueSeven,
			ExpectedValue: true,
		},
		{
			Title:         "Eight",
			GivenMask:     grid.Mask(0b100000000),
			GivenValue:    cell.ValueEight,
			ExpectedValue: true,
		},
		{
			Title:         "Nine",
			GivenMask:     grid.Mask(0b1000000000),
			GivenValue:    cell.ValueNine,
			ExpectedValue: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Title, func(t *testing.T) {
			require.Equal(t, tc.ExpectedValue, tc.GivenMask.IsSet(tc.GivenValue))
			require.Equal(t, !tc.ExpectedValue, (tc.GivenMask ^ grid.CompletedGridMask).IsSet(tc.GivenValue))
		})
	}
}
