#Makefiles
TERRAFORM=aws/terraform
DARWIN=darwin/deploy
WEBHOOK=github/zup-webhook
NAVIGATE_HANDBOOK=github/navigate-handbook
SEARCH_HANDBOOK=github/search-handbook
JENKINS_JOB=jenkins/jobs
SC_COFFEE=scaffold/coffee
SC_SPRING=scaffold/spring-iti
KAFKA=kafka
VIVO=vivo/deploy
DOCKER=docker/compose
KUBERNETES=kubernetes/core
FAST_MERGE=github/fast-merge
FORMULAS=$(TERRAFORM) $(DARWIN) $(WEBHOOK) $(JENKINS_JOB) $(SC_COFFEE) $(SC_SPRING) $(KAFKA) $(VIVO) $(DOCKER) $(NAVIGATE_HANDBOOK) $(SEARCH_HANDBOOK) $(KUBERNETES) $(FAST_MERGE)

PWD_INITIAL=$(shell pwd)

push-s3:
	echo "Init pwd: $(PWD_INITIAL)"
	for formula in $(FORMULAS); do cd $$formula/src && make build && cd $(PWD_INITIAL); done
	./copy-bin-configs.sh "$(FORMULAS)"
	zip -r formulas.zip formulas
	cp formulas.zip formulas
	aws s3 cp . s3://ritchie-cli-bucket152849730126474/ --exclude "*" --include "formulas/*" --recursive
	rm -rf formulas
	rm -rf formulas.zip

bin:
	echo "Init pwd: $(PWD_INITIAL)"
	for formula in $(FORMULAS); do cd $$formula/src && make build && cd $(PWD_INITIAL); done
	./copy-bin-configs.sh "$(FORMULAS)"

test-local: bin
	rm -rf ~/.rit/formulas
	rm -rf ~/.rit/.cmd_tree.json
	mv formulas ~/.rit
	cp tree/tree.json  ~/.rit/.cmd_tree.json