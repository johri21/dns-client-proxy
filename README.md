# DNS-Over-TLS

The implementation is based on an awesome [dns](https://github.com/miekg/dns) library from [Miek Gieben](https://github.com/miekg). The implementation starts a tcp and a udp server on 53 port and uses [Cloudflare DNS over TLS](https://developers.cloudflare.com/1.1.1.1/dns-over-tls/) as a upstream DNS resolver.
## Build
To build the project Makefile is provided. cd into the directory and run the following command.
```sh
$ make build
```
The above command builds the project and creates a docker image abhinav/dns-client-proxy. Thereafter removing the build binary.

## Run
To run the container use the following command.
```sh
$ sudo docker run -d --rm -p53:53 -p53:53/udp abhinav/dns-client-proxy
```
This will start the server on 127.0.0.1:53 binding localhost 53 port.

## Test
To test the application on linux you can use [dig](https://en.wikipedia.org/wiki/Dig_(command)) command.
```sh
$ dig @127.0.0.1 www.zomato.com
```
If everything goes well you may see the dig output. The dig is communicating to localhost at 53 port which in turn uses cloudflare DNS for resolution
