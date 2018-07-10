package main

import "github.com/HAL-RO-Developer/caseTeamB/router"

func main() {
	r := router.GetRouter()
	r.Run(":8080")
}
