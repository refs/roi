package human

import (
	"fmt"
	"math"
	"time"
)

// Duration is a poor-mans attempt to humanize a duration. It has poor granularity, but it does the job pretty well.
// inspiration: https://gist.github.com/harshavardhana/327e0577c4fed9211f65 slightly modified.
func Duration(duration time.Duration) string {
	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))

	if days == 0 && hours < 1 {
		return "within an hour"
	} else if days == 0 {
		return "within a day"
	} else if weeks := days / 7; weeks > 0 {
		if years := days / 365; years > 0 {
			if years == 1 {
				return "within a year"
			}
			return fmt.Sprintf("within %v years", years)
		}

		return fmt.Sprintf("within %v weeks", weeks)
	} else if weeks == 0 {
		return "within a week"
	}

	return ""
}
