#Makefiles
TERRAFORM=aws/terraform
DARWIN=darwin/deploy
WEBHOOK=github/zup-webhook
JENKINS_JOB=jenkins/jobs
SC_COFFEE=scaffold/coffee
SC_SPRING=scaffold/spring-iti
KAFKA_TOPIC_LIST=kafka/list/topic
KAFKA_TOPIC_CREATE=kafka/create/topic
KAFKA_CONSUME=kafka/consume
VIVO=vivo/deploy

FORMULAS=$(TERRAFORM) $(DARWIN) $(WEBHOOK) $(JENKINS_JOB) $(SC_COFFEE) $(SC_SPRING) $(KAFKA_TOPIC_LIST) $(KAFKA_TOPIC_CREATE) $(KAFKA_CONSUME) $(VIVO)
PWD_INITIAL=$(shell pwd)

push-s3:
	echo "Init pwd: $(PWD_INITIAL)"
	for formula in $(FORMULAS); do cd $$formula/src && make build && cd $(PWD_INITIAL); done
	for formula in $(FORMULAS); do mkdir -p formulas/$$formula && cp $$formula/config.json formulas/$$formula && cp -rf $$formula/bin formulas/$$formula; done
	zip -r formulas.zip formulas
	aws s3 sync . s3://ritchie-cli-bucket152849730126474/formulas --exclude "*" --include "formulas.zip"
	rm -rf formulas
	rm -rf formulas.zip

bin:
	echo "Init pwd: $(PWD_INITIAL)"
	for formula in $(FORMULAS); do cd $$formula/src && make build && cd $(PWD_INITIAL); done
	for formula in $(FORMULAS); do mkdir -p formulas/$$formula && cp $$formula/config.json formulas/$$formula && cp -rf $$formula/bin formulas/$$formula; done