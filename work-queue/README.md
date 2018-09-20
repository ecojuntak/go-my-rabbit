On this tutorial, I learned
- How to create multiple subscription in one channel.
- How RabbitMQ implement Round-robin dispatching, to make sure that the consumer consume the same total of message.
```
By default, RabbitMQ will send each message to the next consumer, in sequence. 
On average every consumer will get the same number of messages. 
This way of distributing messages is called round-robin
- RabbitMQ Tutorial
```
- How RabbitMQ avoid if there is one consumer failed to process the received message. 
```
If a consumer dies (its channel is closed, connection is closed, or TCP connection is lost) without sending 
an ack, RabbitMQ will understand that a message wasn't processed fully and will re-queue it. 
If there are other consumers online at the same time, it will then quickly redeliver it to another consumer. 
That way you can be sure that no message is lost, even if the workers occasionally die.
- RabbitMQ Tutorial
```
- How RabbitMQ handle the message durability, in case the RabbitMQ server is down and we don't want to lost our messages
