# palwell-distributed-services

This repository serves as an introduction to
building distributed services using Go.

The services are built with the following architecture and business objectives in mind.

1.) The service has an in memory commit log
2.)

### JSON and HTTP

JSON over HTTP application programmable interfaces (APIs) are the most common APIs on the web. 
They’re simple to build because modern day programming languages have out of the box support for 
JSON I/O, JSON is simple to understand and read, and are a defacto 
standard for exchanging data over the web. What's more, one can call HTTP APIs 
via the terminal with curl, by visiting the site with your browser, or using any of 
the plethora of good HTTP clients currently available. If you have an idea for a web 
service that you want to hack up and have people try as soon as possible, 
then implementing it with JSON/HTTP is the way to go [1]

### Protocol Buffers

For their internal web APIs, an enterprise may take advantage of technologies like 
protobuf for features that JSON/HTTP doesn’t provide, e.g. type checking and versioning. However,
their public APIs will still use JSON/HTTP for accessibility.

The advantages of using protobuf are that it: 

1. Guarantees type-safety; 
2. Prevents schema-violations; 
3. Enables fast serialization; 
4. Offers backward compatibility.



###Resources
[1] Jeffery, Travis. Distributed Services with Go (p. 20). Pragmatic Bookshelf.

