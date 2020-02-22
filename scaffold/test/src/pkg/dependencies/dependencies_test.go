package dependencies

import (
	"fmt"
	"os"
	"testing"
)

func TestAdditionDependencies(t *testing.T) {

	xmlFile, err := os.Open("pom_test.xml")
	if err != nil {
		t.Errorf("Error: %d", err)
	}
	defer xmlFile.Close()

	str := additionDependencies(xmlFile)
	if err != nil {
		t.Errorf("Error: %d", err)
	}

	fmt.Println(str)
}
