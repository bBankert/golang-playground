package maps

import "fmt"

func DisplayMaps() {
	websites := map[string]string{
		"google": "https://google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites)

	websites["linkedin"] = "https://linkedin.com"

	fmt.Println(websites)
}
