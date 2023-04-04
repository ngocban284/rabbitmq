const amqp = require("amqplib");

async function sendToGo() {
  const connection = await amqp.connect("amqp://admin:password@localhost:5672");

  const channel = await connection.createChannel();

  const queue = "js-to-go";
  const exchange = "contactJS";
  // create exchange
  channel.assertExchange(exchange, "fanout", { durable: false });
  channel.assertQueue(queue, { durable: false });
  channel.bindQueue(queue, exchange, "");

  // publish to exchange
  channel.publish(exchange, "", Buffer.from("Hello from Node.js!"));
  console.log("Sent message to Go");
}

async function receivedFromGo() {
  const connection = await amqp.connect("amqp://admin:password@localhost:5672");
  const channel = await connection.createChannel();

  const queue = "go-to-js";
  const exchange = "contactGo";
  channel.assertQueue(queue, { durable: false });
  channel.assertExchange(exchange, "fanout", { durable: false });

  console.log(`Waiting for messages on ${queue}`);

  channel.consume(
    queue,
    (msg) => {
      console.log(`Received message from Go: ${msg.content.toString()}`);
      // Process the message here
    },
    { noAck: true }
  );
}

sendToGo();
receivedFromGo();
