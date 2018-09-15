.PHONY: install run/server run test

certs:

certs/ui:

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
