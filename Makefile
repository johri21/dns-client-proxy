build:
	go build -o dns-client-proxy
	docker build -t johri21/dns-client-proxy .
	rm -f dns-client-proxy
