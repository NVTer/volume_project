package main

import (
	"fmt"

	"github.com/NVTer/volume_project/internal/calculator"
	"github.com/NVTer/volume_project/internal/cli"
)

func main() {
	cliParams := cli.Load()
	volume, err := calculator.Calculate(cliParams.Shape, cliParams.Radius, cliParams.Length,
		cliParams.Width, cliParams.Hieght)
	if err == nil {
		fmt.Println(volume)
	} else {
		fmt.Printf("Error: %v", err)
	}
}
