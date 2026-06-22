# Go Real-Time Internet Speed Test with Live Speedometer

A real-time internet speed test application built with **Go** and a browser-based UI.
It provides a **live animated speedometer**, **ping measurement**, **download/upload speed tracking**, and a clean frontend dashboard.

> This project is designed as the foundation for a full internet speed testing tool similar to Fast.com / Speedtest-style apps, with a custom Go backend and live UI updates using **Server-Sent Events (SSE)**.

---

## Features

* **Real-time speedometer UI**

    * Animated needle-based speedometer
    * Live speed updates during the test
    * Download and upload phase visualization

* **Ping measurement**

    * Shows ping before transfer testing begins

* **Live test status**

    * Displays current phase:

        * Initializing
        * Ping
        * Download
        * Upload
        * Done

* **Go backend with SSE**

    * Streams test progress to the browser in real time
    * No page refresh needed

* **Modern UI**

    * Responsive glassmorphism-style dashboard
    * Final download / upload / ping cards

---

## Current Project Status

This version is the **real-time UI foundation** of the project.

### What is implemented

* Go backend server
* SSE event streaming endpoint
* Real-time frontend speedometer UI
* Live phase/status updates
* Ping / download / upload / final result rendering

### What is not yet fully implemented

The current backend logic is structured like a **real speed test flow**, but if you are using the version from the initial setup, the transfer values may still be **simulated** or **not yet tuned to match production-grade services** like:

* Fast.com
* Ookla Speedtest
* LibreSpeed

To make it a **fully accurate real internet speed test**, the backend should be upgraded to:

* real download throughput testing using multiple parallel streams
* real upload throughput testing
* exact download/upload data usage tracking
* stable sampling and averaging logic
* server tuning for higher accuracy

---

## Project Structure

```bash
go-realtime-speedometer/
├─ go.mod
├─ main.go
└─ web/
   └─ index.html
```

---

## Tech Stack

### Backend

* **Go**
* **net/http**
* **Server-Sent Events (SSE)** for real-time updates

### Frontend

* **HTML**
* **CSS**
* **Vanilla JavaScript**
* SVG/CSS-based speedometer UI logic

---

## How It Works

### 1. Frontend starts the speed test

When the user clicks **Start Test**, the browser opens an SSE connection to:

```http
GET /api/speedtest/stream
```

### 2. Backend streams test events

The backend sends JSON events such as:

* `status`
* `ping`
* `download`
* `upload`
* `done`

### 3. Frontend updates the UI in real time

The browser listens to the SSE stream and updates:

* speedometer needle
* current Mbps
* ping
* final download/upload values
* status text

---

## Event Flow

Typical event sequence:

1. `status` → Initializing speed test
2. `status` → Testing ping
3. `ping` → Ping result
4. `status` → Testing download speed
5. `download` → live download speed updates
6. `status` → Testing upload speed
7. `upload` → live upload speed updates
8. `done` → final result

---

## API

## `GET /api/speedtest/stream`

Streams real-time speed test updates using **Server-Sent Events (SSE)**.

### Response type

```http
Content-Type: text/event-stream
```

### Example event payload

```json
{
  "type": "download",
  "value": 72.4,
  "max": 150,
  "message": "Download 54%"
}
```

### Event fields

| Field        | Type   | Description                                                          |
| ------------ | ------ | -------------------------------------------------------------------- |
| `type`       | string | Event type (`status`, `ping`, `download`, `upload`, `done`, `error`) |
| `message`    | string | Human-readable status message                                        |
| `value`      | float  | Current speed in Mbps                                                |
| `max`        | float  | Maximum gauge range                                                  |
| `ping`       | float  | Ping in milliseconds                                                 |
| `final_down` | float  | Final measured download speed                                        |
| `final_up`   | float  | Final measured upload speed                                          |

---

## UI Overview

The frontend dashboard includes:

### Main Speedometer Panel

* Large circular speedometer
* Animated needle
* Current Mbps display
* Current phase label (`PING`, `DOWNLOAD`, `UPLOAD`, `DONE`)

### Metrics Panel

* Ping
* Final Download
* Final Upload

---

## Installation

## 1. Clone the project

```bash
git clone https://github.com/aliraza910/golang-session.git
cd go-realtime-speedometer
```

## 2. Run the server

```bash
go run main.go
```

## 3. Open in browser

```bash
http://localhost:8080
```

---

## Example Output

During the test, the UI will update live with values like:

* **Ping:** `21.40 ms`
* **Download:** `68.20 Mbps`
* **Upload:** `17.50 Mbps`

Final output might look like:

* **Final Download:** `72.84 Mbps`
* **Final Upload:** `18.92 Mbps`
* **Ping:** `19.60 ms`

---

## Important Notes About Accuracy

This project currently focuses on the **real-time UI + streaming architecture**.

If you want production-level speed test accuracy similar to Fast.com or Ookla, you should improve the backend test engine with:

* multiple parallel download connections
* multiple parallel upload streams
* warm-up exclusion
* stable sample averaging
* stronger test server bandwidth
* accurate byte counters for data usage reporting

---

## Roadmap

Planned upgrades for the next version:

* [x] Real internet download speed measurement
* [x] Real internet upload speed measurement
* [x] Accurate **download/upload data usage** tracking
* [x] Ping + jitter statistics
* [ ] Test history
* [ ] SQLite storage for results
* [x] Better gauge rendering
* [x] Cancel test support
* [x] Multi-server selection
* [ ] React frontend version
* [x] Production-grade accuracy tuning

---

## Future Version Goals

The next major version of this project can include:

### Backend

* real download endpoint using random in-memory buffers
* real upload endpoint with byte counting
* multi-stream testing
* bandwidth sampling every 200ms
* final average speed calculation
* test result persistence

### Frontend

* premium animated speedometer
* live graphs
* phase timeline
* data used counters
* jitter display
* responsive dashboard cards
* dark/light themes

---

## Development Notes

### Why SSE?

Server-Sent Events are used because they are simple and work well for **one-way real-time streaming** from backend to frontend.

This makes them a good fit for:

* progress updates
* speed samples
* status messages
* final result delivery

### Why a browser UI?

A browser-based UI makes it easy to:

* visualize speed live
* animate the speedometer
* make the project portable
* later convert it into a React/Vite dashboard or even a desktop app wrapper

---

## License

This project is open-source and can be used as a base for custom internet speed testing dashboards, educational networking projects, or performance monitoring tools.

Add your preferred license here, for example:

* MIT
* Apache 2.0
* GPL

---

## Author

Built with Go + real-time streaming UI.

If you extend this project into a full production-grade speed test, consider adding:

* real throughput engine
* exact bandwidth usage tracking
* multi-region test servers
* historical reporting dashboard

---
