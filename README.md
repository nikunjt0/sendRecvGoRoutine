# sendRecvGoRoutine
playing around with GoLang and Kafka: started a simple kafka server locally via docker, created two goroutines called gSend and gRecv, gSend: to send data into kafka, gRecv: to receive data In gSend, pushed json data into the kafka server continuously, and at random time intervals In gRecv received the json data that's sent continuously 
