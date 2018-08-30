FROM ubuntu:18.04

RUN apt-get update 

#update root-ca
RUN apt-get install -y ca-certificates

ADD dns-client-proxy /bin/dns-client-proxy
EXPOSE 53
EXPOSE 53/udp

ENTRYPOINT ["/bin/dns-client-proxy"]
