# go-microservices
a very minimalist microservices example

repo includes a "monolith" server, and which gets the time and the date from two other microservices calling each microservice using go-routines, to emphasize parallelism.

## Getting started
Assuming working directory to be `./go-microservices`:

### The Monolith
copy to /go-microservices/monolith/monolith.go
open a terminal
`$ go run ./monolith/main.go`

### The Date Microservice
copy to /go-microservices/dateService/main.go
open a second terminal
`$ go run ./dateService/main.go`

### The Time Microservice
copy to /go-microservices/timeService/main.go
open a third terminal
`$ go run ./dateService/main.go`
