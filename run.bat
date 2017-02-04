set SERVICE_HOST=:8080

set DB_HOST=192.168.1.2
set DB_PORT=5432
set DB_DATABASE=TruckMonitor
set DB_USER=postgres
set DB_PASSWORD=postgres

go run cmd/server.go
