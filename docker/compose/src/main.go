package main

import (
	"compose/pkg/compose"
	"compose/pkg/prompt"
)

func main() {
	var selectItems []string
	selectItem := ""
	postgresDB := ""
	postgresUser := ""
	postgresPassword := ""
	items := []string{"awsclivl", "consul", "dynamoDB", "jaeger", "kafka", "postgres", "redis", "stubby4j", "finish!"}
	for selectItem != "finish!" {
		selectItem, _ = prompt.List("Select docker image: ", items)
		if selectItem == "postgres" {
			postgresDB, _ = prompt.String("Type DB name: ", true)
			postgresUser, _ = prompt.String("Type DB user: ", true)
			postgresPassword, _ = prompt.String("Type DB password: ", true)
		}
		selectItems = append(selectItems, selectItem)
		for i, item := range items {
			if item == selectItem { //Remove input to list
				items = append(items[:i], items[i+1:]...)
				break
			}
		}
	}

	compose.GenerateYml(selectItems, postgresDB, postgresUser, postgresPassword)
}