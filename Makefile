install:
	./install.sh

bin/mackerel-agent:
	./author/build.sh

qnap-mackerel-agent.tar.gz:
	git archive HEAD --output=$@

.PHONY: install
