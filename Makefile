.PHONY: build/dev build cfssl install run/server run test

build/dev: build
	$(shell cd pkg; go-bindata -debug -pkg cloudtables -prefix "../" ../ui/...)

build:
	$(shell export GO111MODULE=on; go build pkg/*)
	$(shell export GO111MODULE=on; go build mock/*)

cert/server: cfssl
	$(shell cd tls; cfssl gencert -initca ca_csr.json | cfssljson -bare ca)
	$(shell cd tls; cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=ca_config.json \
		-profile=client-server \
		server_csr.json | cfssljson -bare cert)

cert/client:
	$(shell cd tls; cfssl gencert \
		-ca=ca.pem \
		-ca-key=ca-key.pem \
		-config=ca_config.json \
		-profile=client \
		client_csr.json | cfssljson -bare client)
	$(shell cd tls; openssl pkcs12 -export -out cloudtables_user.p12 \
		-inkey client-key.pem -in client.pem -certfile ca.pem)

cert/clean:
	$(shell rm -f tls/*.pem)
	$(shell rm -f tls/*.csr)
	$(shell rm -f tls/*.p12)

cfssl:
	$(shell go get -u github.com/cloudflare/cfssl/cmd/...)

init/test:
	$(shell cd pkg; ginkgo bootstrap)
	$(shell cd pkg; ginkgo generate cloudtables)

install:

run/server:

run/dev: build/dev
	$(shell go run cmd/cloudtables/cloudtables.go)

test:
	cd pkg; ginkgo

test/debug:
	cd pkg; DEBUG=true ginkgo