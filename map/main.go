package main

import "fmt"

func main() {
	company := map[string]interface{}{
		"name": "TechCorp",
		"departments": map[string]interface{}{
			"IT": map[string]interface{}{
				"head":  "Rahul",
				"staff": 20,
			},
			"HR": map[string]interface{}{
				"head":  "Priya",
				"staff": 8,
			},
		},
	}

	depts := company["departments"].(map[string]interface{})
	it := depts["IT"].(map[string]interface{})
	fmt.Println("IT Head:", it["head"])
	fmt.Println("IT Staff:", it["staff"])
}
