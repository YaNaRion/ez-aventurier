package class

import (
	"log"
	"sync"
	"time"
)

var (
	montrealLoc *time.Location
	once        sync.Once
)

// GetMontrealLocation returns the America/Toronto timezone (Montreal)
func GetMontrealLocation() *time.Location {
	once.Do(func() {
		var err error
		montrealLoc, err = time.LoadLocation("America/Toronto")
		if err != nil {
			log.Printf("Warning: Failed to load Montreal timezone: %v, using fixed offset", err)
			// Fallback to fixed EST (not ideal but works as fallback)
			montrealLoc = time.FixedZone("EST", -5*3600)
		}
	})
	return montrealLoc
}

// NowMontreal returns current time in Montreal timezone
func NowMontreal() time.Time {
	return time.Now().In(GetMontrealLocation())
}

// ToUTC converts Montreal time to UTC for storage
func ToUTC(montrealTime time.Time) time.Time {
	if montrealTime.IsZero() {
		return montrealTime
	}
	// If the time doesn't have a location set, assume it's Montreal
	if montrealTime.Location() == time.UTC {
		// It might be UTC already, convert to Montreal first to ensure correct conversion
		montrealTime = montrealTime.In(GetMontrealLocation())
	}
	return montrealTime.UTC()
}

// FromUTC converts UTC time to Montreal time for display
func FromUTC(utcTime time.Time) time.Time {
	if utcTime.IsZero() {
		return utcTime
	}
	return utcTime.In(GetMontrealLocation())
}

// ParseMontrealTime parses a string in "YYYY-MM-DDTHH:mm" format as Montreal time
func ParseMontrealTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}

	// Parse as Montreal time
	t, err := time.ParseInLocation("2006-01-02T15:04", timeStr, GetMontrealLocation())
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

// FormatMontrealTime formats a UTC time as Montreal time string for display
func FormatMontrealTime(utcTime time.Time) string {
	if utcTime.IsZero() {
		return ""
	}
	return utcTime.In(GetMontrealLocation()).Format("2006-01-02 15:04 MST")
}

// FormatMontrealTimeInput formats a UTC time as Montreal time string for datetime-local input
func FormatMontrealTimeInput(utcTime time.Time) string {
	if utcTime.IsZero() {
		return ""
	}
	// datetime-local expects format: YYYY-MM-DDTHH:mm (no timezone)
	return utcTime.In(GetMontrealLocation()).Format("2006-01-02T15:04")
}
