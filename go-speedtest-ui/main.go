package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type SpeedEvent struct {
	Type      string  `json:"type"` // status, ping, download, upload, done, error
	Message   string  `json:"message,omitempty"`
	Value     float64 `json:"value,omitempty"`      // current speed Mbps
	Max       float64 `json:"max,omitempty"`        // optional gauge max
	Ping      float64 `json:"ping,omitempty"`       // ms
	FinalDown float64 `json:"final_down,omitempty"` // Mbps
	FinalUp   float64 `json:"final_up,omitempty"`   // Mbps
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/api/speedtest/stream", speedTestStreamHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func speedTestStreamHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	sendEvent(w, SpeedEvent{
		Type:    "status",
		Message: "Initializing speed test...",
	})
	flusher.Flush()

	time.Sleep(700 * time.Millisecond)

	// --------------------
	// Ping phase
	// --------------------
	sendEvent(w, SpeedEvent{
		Type:    "status",
		Message: "Testing ping...",
	})
	flusher.Flush()

	time.Sleep(500 * time.Millisecond)

	ping := 18 + rand.Float64()*20
	sendEvent(w, SpeedEvent{
		Type: "ping",
		Ping: ping,
	})
	flusher.Flush()

	time.Sleep(700 * time.Millisecond)

	// --------------------
	// Download phase
	// --------------------
	sendEvent(w, SpeedEvent{
		Type:    "status",
		Message: "Testing download speed...",
	})
	flusher.Flush()

	var finalDown float64
	for i := 0; i <= 100; i++ {
		// Simulate a realistic curve:
		// ramp up -> stabilize -> slight variation
		var current float64
		switch {
		case i < 20:
			current = float64(i)*2.8 + rand.Float64()*3
		case i < 70:
			current = 55 + rand.Float64()*20
		default:
			current = 62 + rand.Float64()*12
		}

		finalDown = current

		sendEvent(w, SpeedEvent{
			Type:    "download",
			Value:   current,
			Max:     150,
			Message: fmt.Sprintf("Download %d%%", i),
		})
		flusher.Flush()

		time.Sleep(90 * time.Millisecond)
	}

	// --------------------
	// Upload phase
	// --------------------
	sendEvent(w, SpeedEvent{
		Type:    "status",
		Message: "Testing upload speed...",
	})
	flusher.Flush()

	time.Sleep(500 * time.Millisecond)

	var finalUp float64
	for i := 0; i <= 100; i++ {
		var current float64
		switch {
		case i < 20:
			current = float64(i)*0.9 + rand.Float64()*1.5
		case i < 70:
			current = 14 + rand.Float64()*7
		default:
			current = 16 + rand.Float64()*5
		}

		finalUp = current

		sendEvent(w, SpeedEvent{
			Type:    "upload",
			Value:   current,
			Max:     50,
			Message: fmt.Sprintf("Upload %d%%", i),
		})
		flusher.Flush()

		time.Sleep(90 * time.Millisecond)
	}

	// --------------------
	// Done
	// --------------------
	sendEvent(w, SpeedEvent{
		Type:      "done",
		Message:   "Speed test completed",
		Ping:      ping,
		FinalDown: finalDown,
		FinalUp:   finalUp,
	})
	flusher.Flush()
}

func sendEvent(w http.ResponseWriter, event SpeedEvent) {
	data, _ := json.Marshal(event)
	fmt.Fprintf(w, "data: %s\n\n", data)
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
