# BeatHub

A backend service built in **Go** to manage and filter data from the osu! API v2. This project focuses on building a scalable, containerized architecture using modern cloud-native patterns.


> [!NOTE]  
> Project has not been in development for quite some time. Planning to get back to it at some point...

---

## Tools

* **Language:** Go (Golang)
* **API Framework:** Echo v4
* **Database:** PostgreSQL 16
* **Caching:** Valkey 8
* **Containerization:** Docker / Docker Compose
* **Hot Reload:** Air

---

## Project Structure

* `cmd/api/`: Application entry point and router initialization.
* `internal/api/`: HTTP client implementation for osu! API v2 integration.
* `internal/handlers/`: REST controllers for beatmap and user logic.
* `internal/schemas/`: Go structs for API responses (JSON Marshaling).

---

## Deployment

```bash
docker-compose up --build
