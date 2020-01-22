package main

import (
	"deploy/pkg/build"
	"deploy/pkg/circle"
	"deploy/pkg/deployment"
	"deploy/pkg/login"
	"deploy/pkg/modules"
	"deploy/pkg/prompt"
	"deploy/pkg/tags"
	"deploy/pkg/user"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	username := os.Getenv("USERNAME")
	pwd := os.Getenv("PASSWORD")

	token := login.Login(username, pwd)
	user := user.GetUserId(token, username)

	circleSelect := readCircle(token)

	moduleSelect := readModule(token)

	buildModule := readBuildModule(token, moduleSelect)

	releaseName := readReleaseName()

	buildGenerate := build.Build{
		ReleaseName: releaseName,
		AuthorId:    user.Id,
		Modules:     []build.Module{buildModule},
	}

	responseBuild := build.CreateBuild(buildGenerate, token)
	deploymentRequest := deployment.Deployment{
		AuthorId: user.Id,
		CircleId: circleSelect.Id,
		BuildId:  responseBuild.Id,
	}
	response := deployment.CreateDeployment(deploymentRequest, token)
	log.Printf("Id deploy: %v\nStatus: %v\n", response.Id, response.Status)

	for {
		log.Println("Waiting for deploy.")
		time.Sleep(time.Second * 20)
		token = login.Login(username, pwd)
		response = deployment.GetStatus(token, response.Id)
		log.Printf("Id deploy: %v\nStatus: %v\n", response.Id, response.Status)
		if "DEPLOYED" == response.Status {
			log.Println("Deploy finished!")
			break
		}
		if "DEPLOY_FAILED" == response.Status || "NOT_DEPLOYED" == response.Status {
			log.Println("Failed in Deploy!")
			break
		}
	}
}

func readBuildModule(token string, mod modules.Module) build.Module {
	componentsSelect := readComponent(mod.Components)
	buildModule := build.Module{
		Id:         mod.Id,
		Components: []build.Component{},
	}

	for _, c := range componentsSelect {
		buildComponent := readTag(token, c)
		buildModule.Components = append(buildModule.Components, buildComponent)
	}
	return buildModule
}

func readReleaseName() string {
	releaseName, err := prompt.String("Type ReleaseName: ", true)
	if err != nil {
		log.Fatal(err)
	}
	return releaseName
}

func readTag(token string, c modules.Component) build.Component {
	partName := readPartNameTag()
	tagList := tags.GetTagsByComponentId(token, c.Id, partName)
	if len(tagList.Tags) == 0 {
		log.Printf("No result to tag: %v\n", partName)
		return readTag(token, c)
	}
	var items []string
	for _, tag := range tagList.Tags {
		items = append(items, fmt.Sprint(tag.Version, " ", tag.Artifact))
	}
	selectItem, err := prompt.List("Select artifact to deploy of component:", items)
	if err != nil {
		log.Fatal(err)
	}
	itemSelect := strings.Split(selectItem, " ")
	return build.Component{
		Id:       c.Id,
		Version:  itemSelect[0],
		Artifact: itemSelect[1],
	}
}

func readPartNameTag() string {
	moduleName, err := prompt.String("Type name or part name of tag: ", false)
	if err != nil {
		log.Fatal(err)
	}
	return moduleName
}


func readPartNameModule() string {
	moduleName, err := prompt.String("Type name or part name of module to deploy: ", false)
	if err != nil {
		log.Fatal(err)
	}
	return moduleName
}

func readComponent(componentList []modules.Component) []modules.Component {
	var items []string
	for _, component := range componentList {
		items = append(items, fmt.Sprint(component.Id, " ", component.Name))
	}
	selectItem, err := prompt.List("Select component to deploy:", items)
	if err != nil {
		log.Fatal(err)
	}
	idItem := strings.Split(selectItem, " ")[0]
	var components []modules.Component
	for _, c := range componentList {
		if idItem == c.Id {
			components = append(components, c)
			break
		}
	}
	if len(components) == len(componentList) {
		return components
	}
	for selectMore("Add more component? ") {
		selectItem, err := prompt.List("Select component to deploy:", items)
		if err != nil {
			log.Fatal(err)
		}
		idItem := strings.Split(selectItem, " ")[0]
		for _, c := range componentList {
			if idItem == c.Id {
				components = append(components, c)
				break
			}
		}
		if len(components) == len(componentList) {
			return components
		}
	}
	return components
}

func selectMore(name string) bool {
	items := []string{"yes", "no"}
	selectItem, err := prompt.List(name, items)
	if err != nil {
		log.Fatal(err)
	}
	return selectItem == "yes"
}

func readModule(token string) modules.Module {
	moduleName := readPartNameModule()
	modulesList := modules.SearchModules(token, moduleName)
	if len(modulesList) == 0 {
		log.Printf("No result to module: %v\n", moduleName)
		return readModule(token)
	}
	var items []string
	for _, module := range modulesList {
		items = append(items, fmt.Sprint(module.Id, " ", module.Name))
	}
	selectItem, err := prompt.List("Select module to deploy:", items)
	if err != nil {
		log.Fatal(err)
	}
	idItem := strings.Split(selectItem, " ")[0]
	for _, module := range modulesList {
		if idItem == module.Id {
			return module
		}
	}
	log.Fatal("Failed to select module!")
	return modules.Module{}
}

func readCircle(token string) circle.Circle {
	circleName := readPartNameCircle()
	circles := circle.SearchCircles(token, circleName)
	if len(circles) == 0 {
		log.Printf("No result to cicle: %v\n", circleName)
		return readCircle(token)
	}
	var items []string
	for _, circle := range circles {
		items = append(items, fmt.Sprint(circle.Id, " ", circle.Name))
	}
	selectItem, err := prompt.List("Select circle to deploy:", items)
	if err != nil {
		log.Fatal(err)
	}
	idItem := strings.Split(selectItem, " ")[0]
	for _, circle := range circles {
		if idItem == circle.Id {
			return circle
		}
	}
	log.Fatal("Failed to select circle!")
	return circle.Circle{}
}

func readPartNameCircle() string {
	circleName, err := prompt.String("Type name or part name of circle: ", false)
	if err != nil {
		log.Fatal(err)
	}
	return circleName
}



