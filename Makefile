.PHONY: cfssl install run/server run test

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
	export GO111MODULE=on; cd pkg; go build .
	export GO111MODULE=on; cd cmd/cloudtables; go build .

run/server:

run:

test:
	cd pkg; ginkgo
