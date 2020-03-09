#Makefiles
BRANCH=
TERRAFORM=aws/terraform
DARWIN=darwin/deploy
WEBHOOK=github/zup-webhook
NAVIGATE_HANDBOOK=github/navigate-handbook
SEARCH_HANDBOOK=github/search-handbook
JENKINS_JOB=jenkins/jobs
SC_COFFEE=scaffold/coffee
SC_SPRING=scaffold/spring-iti
SC_SPRING_STARTER=scaffold/spring-starter
SC_RITCHIE=scaffold/ritchie
KAFKA=kafka
VIVO=vivo/deploy
DOCKER=docker/compose
KUBERNETES=kubernetes/core
FAST_MERGE=github/fast-merge
FORMULAS=$(TERRAFORM) $(DARWIN) $(WEBHOOK) $(JENKINS_JOB) $(SC_COFFEE) $(SC_SPRING) $(SC_SPRING_STARTER) $(KAFKA) $(VIVO) $(DOCKER) $(NAVIGATE_HANDBOOK) $(SEARCH_HANDBOOK) $(KUBERNETES) $(FAST_MERGE) $(SC_RITCHIE)

PWD_INITIAL=$(shell pwd)

FORM = $($(form))

push-s3:
	echo $(RITCHIE_AWS_BUCKET)
	echo "Init pwd: $(PWD_INITIAL)"
	for formula in $(FORMULAS); do cd $$formula/src && make build && cd $(PWD_INITIAL) || exit; done
	./copy-bin-configs.sh "$(FORMULAS)"
	aws s3 cp . s3://$(RITCHIE_AWS_BUCKET)/ --exclude "*" --include "formulas/*" --recursive
	rm -rf formulas

bin:
	echo "Init pwd: $(PWD_INITIAL)"
	echo "Formulas bin: $(FORMULAS)"
	for formula in $(FORMULAS); do cd $$formula/src && make build && cd $(PWD_INITIAL); done
	./copy-bin-configs.sh "$(FORMULAS)"

test-local:
ifneq "$(FORM)" ""
	echo "true: $(FORM)"
	$(MAKE) bin FORMULAS=$(FORM)
	mkdir -p ~/.rit/formulas
	rm -rf ~/.rit/formulas/$(FORM)
	./unzip-bin-configs.sh
	cp -r formulas/* ~/.rit/formulas
	rm -rf formulas
else
	echo "true: $(FORM)"
	$(MAKE) bin
	rm -rf ~/.rit/formulas
	./unzip-bin-configs.sh
	mv formulas ~/.rit
endif
	rm -rf ~/.rit/.cmd_tree.json
	cp tree/tree.json  ~/.rit/.cmd_tree.json