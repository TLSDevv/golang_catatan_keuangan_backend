ENV ?= local

.PHONY: install
install:
	go mod download

.PHONY: config
config:
	ENV=$(ENV) ansible-playbook ./.ansible/playbook.yml