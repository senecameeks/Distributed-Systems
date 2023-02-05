/*
 *  Brown University, CS138, Spring 2018
 *
 *  Purpose: contains all interval related logic.
 */

package liteminer

// Represents [Lower, Upper)
type Interval struct {
	Lower uint64 // Inclusive
	Upper uint64 // Exclusive
}

// GenerateIntervals divides the range [0, upperBound] into numIntervals
// intervals.
func GenerateIntervals(upperBound uint64, numIntervals int) (intervals []Interval) {
	// TODO: Students should implement this.

	overFilled := int(upperBound) % numIntervals
	var offset uint64
	offset = 0

	for i := 0; i < numIntervals; i++ {
		if i < overFilled {
			intervals = append(intervals, Interval{offset, offset + (upperBound / uint64(numIntervals)) + 1})
			offset += (upperBound / uint64(numIntervals)) + 1
		} else {
			intervals = append(intervals, Interval{offset, offset + (upperBound / uint64(numIntervals))})
			offset += (upperBound / uint64(numIntervals))
		}
	}

	intervals[numIntervals-1].Upper += 1
	return
}
