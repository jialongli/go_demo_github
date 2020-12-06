package main

import (
	"fmt"
	"strings"
)

func main() {
	namingAndEnv := strings.Split("dsf_afdas_saf.bj.env", ".")
	fmt.Println(namingAndEnv[0])

	fmt.Sprintf(`["%s"]`, "fadf")
}
