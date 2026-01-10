package main

import (
	"context"
	"fmt"
	"log"
	"middleware/Scheduler/internal/events"
	"os"
	"os/signal"
	"time"

	"github.com/zhashkevych/scheduler"
)

func mainLogic() {
	url := "https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=13295,13345&projectId=3&calType=ical&nbWeeks=8&displayConfigId=128"

	icalData, err := events.FetchICalData(url)
	if err != nil {
		log.Printf("Erreur lors de la récupération des données iCal : %s", err)
		return
	}

	parsedEvents, err := events.ParseIcalToJSON(icalData)
	if err != nil {
		log.Printf("Erreur lors du parsing des données iCal : %s", err)
		return
	}

	err = events.PublishEvents(parsedEvents, "EVENTS.schedule")
	if err != nil {
		log.Printf("Erreur lors de la publication des événements : %s", err)
		return
	}

	for _, event := range parsedEvents {
		fmt.Printf("Événement publié : %s\n", event.Summary)
	}
}

func main() {
	ctx := context.Background()
	sc := scheduler.NewScheduler()

	sc.Add(ctx, func(ctx context.Context) {
		mainLogic()
	}, time.Minute*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	sc.Stop()
}
