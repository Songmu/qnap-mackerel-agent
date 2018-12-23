install:
	./install.sh

bin/mackerel-agent:
	./author/build.sh

bin/mkrinst:
	./author/mkrinst.sh

qnap-mackerel-agent.tar.gz:
	git archive HEAD --prefix=qnap-mackerel-agent/ --output=$@

.PHONY: install
