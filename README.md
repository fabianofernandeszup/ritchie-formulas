# [ Planilha de Teste ](https://docs.google.com/spreadsheets/d/1fnfxnDMJxnjYJ-OLqZuGPVUhUQvyU9CLBCQl2-sPaTc/edit?usp=sharing)

# HOW TO WRITE A FORMULA

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
