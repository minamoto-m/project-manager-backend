package main

import (
	"github.com/minamoto-m/project-manager-backend/internal/infrastructures"
	"github.com/minamoto-m/project-manager-backend/internal/interfaces"
)

func main() {
	infrastructures.InitEnvironment()
	db := infrastructures.SetupDB()
	r := interfaces.SetupRouter(db)
	r.Run(":8080")
}
