ENV ?= local

.PHONY: install
install:
	go mod download

.PHONY: config
config:
	ENV=$(ENV) ansible-playbook ./.ansible/playbook.yml

.PHONY: setup-dir
setup-dir:
	mkdir .keys

.PHONY: gen-keys
gen-keys:
	ssh-keygen -m PEM -t rsa -f ./.keys/id_rsa_local.pem -q -P ""
	openssl rsa -in ./.keys/id_rsa_local.pem -pubout -out ./.keys/id_rsa_local.pub.pem

.PHONY: keys
keys: setup-dir gen-keys