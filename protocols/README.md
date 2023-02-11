# Protocols

## Internet Protocol suite

The internet protocol suite is a collection of different protocols, or methods, for devices to communicate with each other.
Both TCP and UDP are major protocols within the internet protocol suite:

![Internet Protocol Suite](./imgs/ip_suite.png)

Each device that's connected to the internet has a unique IP address. And whenever two devices communicate over the internet, they're likely using either TCP or UDP to do so.

### TCP

TCP, or Transmission Control Protocol, is the most common networking protocol online.
TCP is extremely reliable, and is used for everything from surfing the web (HTTP), sending emails (SMTP), and transferring files (FTP).

TCP is used in situations where it's necessary that all data being sent by one device is received by another completely intact.

For example, when you visit a website, TCP is used to guarantee that everything from the text, images, and code needed to render the page arrives. Without TCP, images or text could be missing, or arrive in the incorrect order, breaking the page.

TCP is a connection-oriented protocol, meaning that it establishes a connection between two devices before transferring data, and maintains that connection throughout the transfer process.

To establish a connection between two devices, TCP uses a method called a three-way handshake:

![TCP Three-Way Handshake](./imgs/tcp_handshake.png)

### UDP

UDP, or User Datagram Protocol, is another one of the major protocols that make up the internet protocol suite. UDP is less reliable than TCP, but is much simpler.

UDP is used for situations where some data loss is acceptable, like live video/audio, or where speed is a critical factor like online gaming.

While UDP is similar to TCP in that it's used to send and receive data online, there are a couple of key differences.

First, UDP is a connectionless protocol, meaning that it does not establish a connection beforehand like TCP does with its three-way handshake.

Next, UDP doesn't guarantee that all data is successfully transferred.
With UDP, data is sent to any device that happens to be listening, but it doesn't care if some of it is lost along the way.
This is one of the reasons why UDP is also known as the "fire-and-forget" protocol.

UDP is more like a protester outside with a megaphone.
Everyone who is paying attention to the protester should hear most of what they're saying.
But there's no guarantee that everyone in the area will hear what the protester is saying, or that they're even listening.

### Nagle's Algorithm

Nagle's algorithm is a means of improving the efficiency of TCP/IP networks by reducing the number of packets that need to be sent over the network.

Typically, when a TCP/IP application sends data, it sends it in small chunks.
This is because the application is usually waiting for a response from the server before sending more data.
This is called "flow control".
However, this means that the application is sending a lot of small packets over the network.
This is inefficient, because the network is optimized for sending large packets.
Nagle's algorithm is a means of reducing the number of packets that need to be sent over the network.

Nagle's algorithm works by delaying the sending of small packets until there is enough data to send a full-sized packet.
This means that the network is only sending full-sized packets, which are more efficient.
However, this means that the application has to wait for the network to send the packet before it can send more data.
This can cause a delay in the application.

Nagle’s algorithm coalesces small packets and delays their delivery until an ACK is returned from the previously sent packet or an adequate amount of small packets is accumulated after a certain period.
This process usually takes milliseconds but, having a latency-sensitive service or tight latency Service Level Objectives (SLOs), shaving off a couple of milliseconds might be worthwhile.

A cross-platform TCP socket option that comes helpful here is `TCP_NODELAY`.
When enabled, it practically disables Nagle’s algorithm.
Instead of coalescing small packets, it sends them to the pipe as soon as possible.
In general, Nagle’s algorithm’s goal is to reduce the number of packets sent to save bandwidth and increase throughput with the trade-off sometimes to introduce increased latency to services.
On the other hand, `TCP_NODELAY` might decrease throughput for small writes, but there are ways to mitigate this by using buffers on the application side.

In Go, TCP_NODELAY is enabled by default, but the standard library offers the ability to disable the behavior via the net.SetNoDelay method.

To test the Nagle's algorithm, use [this example](./tcp_with_nagle_algorithm/):

```bash
# Run the server
$ go run server.go

# Run the client
$ go run client.go
```

## References

[1] [Internet Protocol Suite](https://www.sciencedirect.com/topics/computer-science/internet-protocol-suite)
