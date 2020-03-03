# [ Planilha de Teste ](https://docs.google.com/spreadsheets/d/1fnfxnDMJxnjYJ-OLqZuGPVUhUQvyU9CLBCQl2-sPaTc/edit?usp=sharing)

# HOW TO WRITE A FORMULA

for create a ritchie formula you can choose a automatic way or a manual way.

## Use ritchie to create a ritchie formula (Automatic Way)

on the ritchie-formulas home dir:

#### run:
```
rit scaffold generate ritchie
```
#### and after:
 - inform how the people will can your formula.
 - formula description.

#### if everything goes well:

you can run the generated formula with a hello world message and after change the generated code you can build your changes if:
```
make test-local form=$name
```
- $name is the last name of the formula.

## Create formula manually
To make sure ritchie-cli will be able to process and use your formulas, the following directory structure must be respected:

```yaml
<formula_path>
  └ bin
    └ <formula_bin>-linux
    └ <formula_bin>-darwin
    └ <formula_bin>-windows.exe
  └ src
  config.json
```

Inside the `bin` directory you put the executable file ritchie-cli will run. This binary **must** be able to take all inputs from environment variables to run the formula properly.
Also you need to generate the binaries that will run on Linux, Mac and Windows. 

Inside the `src` directory put all the code and templates required to run your formula.

To handle all external inputs you must create a file located at the root of your formula called `config.json`. The following snippet can be used as base to create your configuration.

```json
{
  "description": "Create your microservice scaffold",
  "inputs" : [
    {
      "name" : "packaging",
      "type" : "text",
      "default" : "jar",
      "label" : "Pick your packaging type[jar]: "
    },
    {
      "name" : "java_version",
      "type" : "text",
      "default" : "11",
      "label" : "Pick your java version [11]: "
    },
    {
      "name" : "language",
      "type" : "text",
      "default" : "kotlin",
      "items" : ["kotlin", "java"],
      "label" : "Select your language: "
    },
    {
      "name" : "group_id",
      "type" : "text",
      "label" : "Type your groupId[ex.: itau.iti.demo]: "
    },
    {
      "name" : "artifact_id",
      "type" : "text",
      "label" : "Type your artifactId[ex.: demo]: "
    },
    {
      "name" : "version",
      "type" : "text",
      "default" : "0.0.1-SNAPSHOT",
      "label" : "Type your version[0.0.1-SNAPSHOT]: "
    },
    {
      "name" : "name",
      "type" : "text",
      "label" : "Type your project name[ex.: demo]: "
    },
    {
      "name" : "description",
      "type" : "text",
      "label" : "Type the project`s description [ex.: project demo]: "
    },
    {
      "name" : "package_name",
      "type" : "text",
      "label" : "Type your package name[ex.: itau.iti.demo]: "
    }
  ]
}
```

Keep in mind you can write your formula using any language you are familiar with(bash,python,golang,etc).

# HOW TO CONFIGURE THE CLI TREE FILE

Modify the `tree/tree.json` file to make sure the cli knows where to put your formula on the execution tree. Inside the command list insert the following snippet:

```json
{
  "commands": [
    ...
    {
      "usage": "<command_name>",
      "help": "<helper>",
      "formula": {
        "path": "<formula_path>",
        "bin": "<formula_bin>-${so}"
      },
      "parent": "<command_tree_parent>"
    },
    ...
    ],
    ...
}
```

As an example, you can use the coffee formula located at `scaffold/coffee`.

# HOW BUILDING AND TESTING A NEW FORMULA

## Configuring the main Makefile

Modify the `ritchie-formulas/Makefile` file by adding the referenes of the new formula.

In the example below we added:
1. New variable with file path of the new formula, e.g. `NEW_FORMULA=new-formula/formula`
2. Add into FORMULAS the new variable, e.g. `$(NEW_FORMULA)`

```Makefile
#Makefiles
TERRAFORM=aws/terraform
DARWIN=darwin/deploy
SC_COFFEE=scaffold/coffee
KUBERNETES=kubernetes/core
NEW_FORMULA=new-formula/formula

FORMULAS=$(TERRAFORM) $(DARWIN) $(SC_COFFEE) $(KUBERNETES) $(NEW_FORMULA)
```

## Building and testing local new formula

After changing the `ritchie-formulas/Makefile` file, you can execute one of the commands:

1. `make test-local`: Build all formulas by moving to the `~/.rit/formulas` folder;

2. `make test-local form=NEW_FORMULA`: Build the specific formula by moving to the folder `~/.rit/formulas`;

The above commands will also update the ~/.rit/.cmd_tree.json file locally.

**Now the formula will be ready to be tested locally.**