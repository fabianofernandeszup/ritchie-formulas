package dependencies

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func additionDependencies(xmlFile *os.File) string {
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
	}
	result := string(byteValue)

	repositoryBeagle :=
		`	<repositories>
		<repository>
			<id>beagle-nexus</id>
			<name>Beagle Nexus Repository</name>
			<url>https://repo-iti.zup.com.br/repository/beagle-jars-all/</url>
		</repository>
	</repositories>
</project>`

	dependencyBeagle :=
		`	<dependency>
			<groupId>br.com.zup.beagle</groupId>
			<artifactId>framework</artifactId>
			<version>0.0.1</version>
		</dependency>
	</dependencies>`

	result = strings.Replace(result, "</project>", repositoryBeagle, -1)
	result = strings.Replace(result, "</dependencies>", dependencyBeagle, 1)

	return result
}
