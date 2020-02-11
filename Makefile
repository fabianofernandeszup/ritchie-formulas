#Makefiles
TERRAFORM=aws/terraform
DARWIN=darwin/deploy
WEBHOOK=github/zup-webhook
NAVIGATE_HANDBOOK=github/navigate-handbook
SEARCH_HANDBOOK=github/search-handbook
JENKINS_JOB=jenkins/jobs
SC_COFFEE=scaffold/coffee
SC_SPRING=scaffold/spring-iti
KAFKA_LIST_TOPIC=kafka/list/topic
KAFKA_CREATE_TOPIC=kafka/create/topic
KAFKA_CONSUME=kafka/consume
KAFKA_PRODUCE=kafka/produce
VIVO=vivo/deploy
DOCKER=docker/compose
FORMULAS=$(TERRAFORM) $(DARWIN) $(WEBHOOK) $(JENKINS_JOB) $(SC_COFFEE) $(SC_SPRING) $(KAFKA_LIST_TOPIC) $(KAFKA_CREATE_TOPIC) $(KAFKA_CONSUME) $(KAFKA_PRODUCE) $(VIVO) $(DOCKER) $(NAVIGATE_HANDBOOK) $(SEARCH_HANDBOOK)

PWD_INITIAL=$(shell pwd)

push-s3:
	echo "Init pwd: $(PWD_INITIAL)"
	for formula in $(FORMULAS); do cd $$formula/src && make build && cd $(PWD_INITIAL); done
	for formula in $(FORMULAS); do mkdir -p formulas/$$formula && cp $$formula/config.json formulas/$$formula && find $$formula/bin -type f -exec md5sum {} \; | sort -k 2 | md5sum | cut -f1 -d ' ' > formulas/$$formula".md5" && cp -rf $$formula/bin formulas/$$formula; done
	zip -r formulas.zip formulas
	aws s3 cp . s3://ritchie-cli-bucket152849730126474/formulas --exclude "*" --include "formulas.zip"
	aws s3 cp . s3://ritchie-cli-bucket152849730126474/ --exclude "*" --include "formulas/*" --recursive
	rm -rf formulas
	rm -rf formulas.zip

bin:
	echo "Init pwd: $(PWD_INITIAL)"
	for formula in $(FORMULAS); do cd $$formula/src && make build && cd $(PWD_INITIAL); done
	for formula in $(FORMULAS); do mkdir -p formulas/$$formula && cp $$formula/config.json formulas/$$formula && cp -rf $$formula/bin formulas/$$formula; done

test-local: bin
	for formula in $(FORMULAS); do mkdir -p formulas/$$formula && cp $$formula/config.json formulas/$$formula && cp -rf $$formula/bin formulas/$$formula; done
	rm -rf ~/.rit/formulas
	rm -rf ~/.rit/.cmd_tree.json
	mv formulas ~/.rit
	cp tree/tree.json  ~/.rit/.cmd_tree.json