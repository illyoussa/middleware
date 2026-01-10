package events

import (
	"strings"
	"time"
)

type Event struct {
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}

func ParseIcalToJSON(icalData string) ([]Event, error) {
	var events []Event
	lines := strings.Split(icalData, "\n")
	var currentEvent *Event

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "BEGIN:VEVENT") {
			currentEvent = &Event{}
		} else if strings.HasPrefix(line, "SUMMARY:") {
			if currentEvent != nil {
				currentEvent.Summary = strings.TrimPrefix(line, "SUMMARY:")
			}
		} else if strings.HasPrefix(line, "DESCRIPTION:") {
			if currentEvent != nil {
				currentEvent.Description = strings.TrimPrefix(line, "DESCRIPTION:")
			}
		} else if strings.HasPrefix(line, " ") && currentEvent != nil && currentEvent.Description != "" {
			currentEvent.Description += strings.TrimSpace(line)
		} else if strings.HasPrefix(line, "DTSTART:") {
			if currentEvent != nil {
				date, err := parseICalDate(strings.TrimPrefix(line, "DTSTART:"))
				if err != nil {
					return nil, err
				}
				currentEvent.Start = date
			}
		} else if strings.HasPrefix(line, "DTEND:") {
			if currentEvent != nil {
				date, err := parseICalDate(strings.TrimPrefix(line, "DTEND:"))
				if err != nil {
					return nil, err
				}
				currentEvent.End = date
			}
		} else if strings.HasPrefix(line, "END:VEVENT") {
			if currentEvent != nil {
				events = append(events, *currentEvent)
				currentEvent = nil
			}
		}
	}

	return events, nil
}

func parseICalDate(rawDate string) (time.Time, error) {
	return time.Parse("20060102T150405Z", rawDate)
}
