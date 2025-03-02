# Faire une tabulation pour chaque commande pour que le Makefile fonctionne correctement

.PHONY: create
create:
	cobra-cli add $(COMMAND)

.PHONY: test
test:
	go test -v ./...