#Makefiles
HELLO_WORLD=github/hello-world
FORMULAS=$(HELLO_WORLD)

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