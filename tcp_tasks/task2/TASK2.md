Task 2 
Create a new version of your Client-Server communication, which uses UDP.

Create a Readme.md file explaining the difference between TCP and UDP. Use code comments/annotations to highlight where the differences are.

_* TCP vs. UDP *_

-  TCP (Transmission Control Protocol) and UDP (User Datagram Protocol) are two communication protocols used top send and receive data over a network.
-  TCP is a connection-oriented protocol that ensures data is delivered reliably and in order.
-  UDP is a connectionless protocol that does not guarantee reliability or ordering, but it's faster and more efficient in scenarios where speed is critical and some loss of data is acceptable (like streaming).

Summary:

-  TCP is useful when you need reliable communication, such as for file transfers or web browsing.
-  UDP is ideal for real-time applications where speed is more critical than reliability, such as for online gaming or video streaming.

Steps:

1. Compile and Run the TCP Server and Client in their respective folders:
   -  Run the server: go run tcpServer.go 8080
   -  Run the client: go run tcpClient.go localhost:8080
2. Compile and Run the UDP Server and Client in their respective folders:
   -  Run the server: go run udpServer.go 8080
   -  Run the client: go run udpServer.go localhost:8080
