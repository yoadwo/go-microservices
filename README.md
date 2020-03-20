# go-microservices
a very minimalist microservices example

repo includes a "monolith" server, and which gets the time and the date from two other microservices calling each microservice using go-routines, to emphasize parallelism.

## Getting started
Assuming working directory to be `./go-microservices`:

### The Monolith
open the first terminal window
`$ go run ./monolith/monolith.go`

### The Date Microservice
open the second terminal window
`$ go run ./dateService/main.go`

### The Time Microservice
open the third terminal window
`$ go run ./dateService/main.go`
