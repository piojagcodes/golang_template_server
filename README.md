# Go HTTP Server

This Go application demonstrates a simple HTTP server implementation using the Go programming language. It provides an HTTP server that listens on port 10101 and exposes an endpoint to retrieve YouTube channel statistics.

## Prerequisites

To run this application, make sure you have the following prerequisites installed:

- Go (version 1.16 or higher)

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/piojagcodes/golang_template_server.git

2. Change into the project directory:

   ```shell
   cd golang_template_server
3. Build the application:
   ```shell
   go build
4. Run the application by executing the generated binary:
   ```shell
   ./go-http-server
5. To gracefully shut down the server, use one of the following methods:
### Methods
   Press Ctrl+C in the terminal where the server is running.
   The server will attempt to complete ongoing requests and then shut down.
   
