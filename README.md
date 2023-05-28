# Problem
Send / Receive data using Kafka server via goroutines

# Explanation
- Setup a simple kafka server locally via docker
- Start two goroutines, call them gSend, gRecv
- gSend: to send data into kafka, gRecv: to receive data
- In gSend, push json data into the kafka server continuously, and at random time intervals
- In gRecv, receive the json data that's sent continuously

json data format:

{
    "datetime": <UTC datetime string>,
    "data" : <integer>

}

Ensure that the goroutines are actively processing information i.e. unless we force kill the go process, it would be running forever


