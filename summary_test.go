package appoptics

import "testing"

func TestUpdateValue(t *testing.T) {
	s := &Summary{
		Count: 3,
		Sum:   8,
		Min:   3,
		Max:   5,
		Last:  3,
	}

	t.Run("UpdateValue", func(t *testing.T) {
		newValue := int64(4)
		preUpdate := *s
		s.UpdateValue(newValue)

		newCount := preUpdate.Count + 1
		if s.Count != newCount {
			t.Errorf("expected Count to be %d but was %d", newCount, s.Count)
		}

		newSum := preUpdate.Sum + newValue
		if s.Sum != newSum {
			t.Errorf("expected Sum to be %d but was %d", newSum, s.Sum)
		}

		if s.Min != 3 {
			t.Errorf("expected Min to be 3 but was %d", s.Min)
		}

		if s.Max != 5 {
			t.Errorf("expected Max to be 5 but was %d", s.Max)
		}

		if s.Last != 4 {
			t.Errorf("expected Last to be 4 but was %d", s.Last)
		}
	})

	t.Run("UpdateValue with new Min", func(t *testing.T) {
		newMin := int64(1)
		s.UpdateValue(newMin)

		if s.Min != newMin {
			t.Errorf("expected Min to be %d but was %d", newMin, s.Min)
		}
	})

	t.Run("UpdateValue with new Max", func(t *testing.T) {
		newMax := int64(7)
		s.UpdateValue(newMax)
		if s.Max != newMax {
			t.Errorf("expected Max to be %d but was %d", newMax, s.Max)
		}
	})
}

func TestUpdateWithZeroValues(t *testing.T) {
	newSummary := Summary{
		Count: 2,
		Sum:   3,
		Min:   1,
		Max:   2,
		Last:  2,
	}

	emptySummary := &Summary{}

	emptySummary.Update(newSummary)

	if emptySummary.Count != newSummary.Count {
		t.Errorf("expected Count to match but %d != %d", emptySummary.Count, newSummary.Count)
	}

	if emptySummary.Sum != newSummary.Sum {
		t.Errorf("expected Sum to match but %d != %d", emptySummary.Sum, newSummary.Sum)
	}

	if emptySummary.Min != newSummary.Min {
		t.Errorf("expected Min to match but %d != %d", emptySummary.Min, newSummary.Min)
	}

	if emptySummary.Max != newSummary.Max {
		t.Errorf("expected Max to match but %d != %d", emptySummary.Max, newSummary.Max)
	}

	if emptySummary.Last != newSummary.Last {
		t.Errorf("expected Last to match but %d != %d", emptySummary.Last, newSummary.Last)
	}

}

func TestUpdateAggregation(t *testing.T) {
	oldSummary := Summary{
		Count: 2,
		Sum:   3,
		Min:   1,
		Max:   2,
		Last:  2,
	}

	newSummary := Summary{
		Count: 2,
		Sum:   5,
		Min:   2,
		Max:   3,
		Last:  3,
	}

	oldSummary.Update(newSummary)

	if oldSummary.Count != 4 {
		t.Errorf("expected Count to be aggregate but was %d", oldSummary.Count)
	}

	if oldSummary.Sum != 8 {
		t.Errorf("expected Sum to be aggregate but was %d", oldSummary.Sum)
	}
}

func TestUpdateWithNewMin(t *testing.T) {
	oldSummary := Summary{
		Count: 2,
		Sum:   6,
		Min:   2,
		Max:   4,
		Last:  2,
	}

	newSummary := Summary{
		Count: 2,
		Sum:   4,
		Min:   1,
		Max:   3,
		Last:  3,
	}

	oldSummary.Update(newSummary)

	if oldSummary.Min != newSummary.Min {
		t.Errorf("expected Min to be reset to %d but was %d", newSummary.Min, oldSummary.Min)
	}

}

func TestUpdateWithNewMax(t *testing.T) {
	oldSummary := Summary{
		Count: 2,
		Sum:   3,
		Min:   1,
		Max:   2,
		Last:  2,
	}

	newSummary := Summary{
		Count: 2,
		Sum:   4,
		Min:   1,
		Max:   3,
		Last:  3,
	}

	oldSummary.Update(newSummary)

	if oldSummary.Max != newSummary.Max {
		t.Errorf("expected Max to be reset to %d but was %d", newSummary.Max, oldSummary.Max)
	}
}
