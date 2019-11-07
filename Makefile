.PHONY: e2e operator

operator:
	./scripts/run-test.sh operator

e2e:
	./scripts/run-test.sh e2e
