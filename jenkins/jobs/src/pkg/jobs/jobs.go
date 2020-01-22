package jobs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	zupJenkinsURL = "https://ci.zup.com.br/createItem?name="
	bodyCiJenkins = `<?xml version='1.1' encoding='UTF-8'?>
	<org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject plugin="workflow-multibranch@2.21">
	  <actions/>
	  <description></description>
	  <properties>
		<org.jenkinsci.plugins.pipeline.modeldefinition.config.FolderConfig plugin="pipeline-model-definition@1.3.9">
		  <dockerLabel></dockerLabel>
		  <registry plugin="docker-commons@1.15"/>
		</org.jenkinsci.plugins.pipeline.modeldefinition.config.FolderConfig>
	  </properties>
	  <folderViews class="jenkins.branch.MultiBranchProjectViewHolder" plugin="branch-api@2.5.4">
		<owner class="org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject" reference="../.."/>
	  </folderViews>
	  <healthMetrics>
		<com.cloudbees.hudson.plugins.folder.health.WorstChildHealthMetric plugin="cloudbees-folder@6.9">
		  <nonRecursive>false</nonRecursive>
		</com.cloudbees.hudson.plugins.folder.health.WorstChildHealthMetric>
	  </healthMetrics>
	  <icon class="jenkins.branch.MetadataActionFolderIcon" plugin="branch-api@2.5.4">
		<owner class="org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject" reference="../.."/>
	  </icon>
	  <orphanedItemStrategy class="com.cloudbees.hudson.plugins.folder.computed.DefaultOrphanedItemStrategy" plugin="cloudbees-folder@6.9">
		<pruneDeadBranches>true</pruneDeadBranches>
		<daysToKeep>-1</daysToKeep>
		<numToKeep>6</numToKeep>
	  </orphanedItemStrategy>
	  <triggers>
		<com.cloudbees.hudson.plugins.folder.computed.PeriodicFolderTrigger plugin="cloudbees-folder@6.9">
		  <spec>H H/4 * * *</spec>
		  <interval>86400000</interval>
		</com.cloudbees.hudson.plugins.folder.computed.PeriodicFolderTrigger>
	  </triggers>
	  <disabled>false</disabled>
	  <sources class="jenkins.branch.MultiBranchProject$BranchSourceList" plugin="branch-api@2.5.4">
		<data>
		  <jenkins.branch.BranchSource>
			<source class="org.jenkinsci.plugins.github_branch_source.GitHubSCMSource" plugin="github-branch-source@2.5.5">
			  <id>0f66e540-cc5d-417c-81d6-a0ebe30697ec</id>
			  <apiUri>https://api.github.com</apiUri>
			  <credentialsId>github-zupci</credentialsId>
			  <repoOwner>ZupIT</repoOwner>
			  <repository>{{RepositoryName}}</repository>
			  <repositoryUrl>https://github.com/ZupIT/{{RepositoryName}}</repositoryUrl>
			  <traits>
				<org.jenkinsci.plugins.github__branch__source.BranchDiscoveryTrait>
				  <strategyId>1</strategyId>
				</org.jenkinsci.plugins.github__branch__source.BranchDiscoveryTrait>
				<org.jenkinsci.plugins.github__branch__source.OriginPullRequestDiscoveryTrait>
				  <strategyId>1</strategyId>
				</org.jenkinsci.plugins.github__branch__source.OriginPullRequestDiscoveryTrait>
				<net.gleske.scmfilter.impl.trait.WildcardSCMHeadFilterTrait plugin="scm-filter-branch-pr@0.4">
				  <includes>qa master release-* development demo-*</includes>
				  <excludes></excludes>
				  <tagIncludes></tagIncludes>
				  <tagExcludes>*</tagExcludes>
				</net.gleske.scmfilter.impl.trait.WildcardSCMHeadFilterTrait>
			  </traits>
			</source>
			<strategy class="jenkins.branch.DefaultBranchPropertyStrategy">
			  <properties class="empty-list"/>
			</strategy>
		  </jenkins.branch.BranchSource>
		</data>
		<owner class="org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject" reference="../.."/>
	  </sources>
	  <factory class="org.jenkinsci.plugins.workflow.multibranch.WorkflowBranchProjectFactory">
		<owner class="org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject" reference="../.."/>
		<scriptPath>Jenkinsfile</scriptPath>
	  </factory>
	</org.jenkinsci.plugins.workflow.multibranch.WorkflowMultiBranchProject>`
)

type Inputs struct {
	JobName      string
	JenkinsUser  string
	JenkinsToken string
}

type Response struct {
	Status int
	Body   string
}

func (in Inputs) Run() {
	log.Println("Zup Jenkins Job Formula Starter!")
	log.Println("Creating Job Jenkins CI")
	bodyJob := strings.ReplaceAll(bodyCiJenkins, "{{RepositoryName}}", in.JobName)
	resp := in.CreateJobsJenkins(bodyJob)
	switch resp.Status {
	case 200:
		log.Printf("Jenkins jobs request Response code: %d", resp.Status)
		log.Println("Jenkins jobs Formula Finished!!!")

	case 400:
		log.Printf("Jenkins jobs request Response code: %d", resp.Status)
		log.Println("Jenkins jobs Formula Existing!!!")

	default:
		log.Printf("Jenkins jobs request Response code: %d", resp.Status)
		log.Println("Jenkins jobs Formula Not Finished!!!")

	}

}

func (in Inputs) CreateJobsJenkins(body string) Response {
	url := fmt.Sprint(zupJenkinsURL, in.JobName)
	xmlCi := []byte(body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlCi))
	if err != nil {
		log.Fatal("Error to create Jobs Jenkins Request: ", err)
	}
	req.Header.Set("Content-Type", "application/xml")
	req.SetBasicAuth(in.JenkinsUser, in.JenkinsToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process Jobs Jenkins Request: ", err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return Response{
		Status: resp.StatusCode,
		Body:   string(bodyBytes),
	}
}
