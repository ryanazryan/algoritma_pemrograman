package main

import "fmt"

const (
	bigParcelPaperWidth    = 60
	bigParcelPaperHeight   = 80
	smallParcelPaperWidth  = 40
	smallParcelPaperHeight = 60
	paperSizeWidth         = 80
	paperSizeHeight        = 80
	bigParcelCost          = 75000
	smallParcelCost        = 45000
	paperCost              = 750
)

func checkEmployees(position string, highLevel, entryLevel *int) {
	if position == "manager" || position == "supervisor" {
		*highLevel++ 
	} else if position == "officer" || position == "staff" || position == "intern" {
		*entryLevel++ 
	}
}

func totalEmployees(highLevel, entryLevel int) int {
	return highLevel + entryLevel
}

func totalPaperNeededSize(highLevel, entryLevel int) int {
	return (highLevel * bigParcelPaperWidth * bigParcelPaperHeight) + (entryLevel * smallParcelPaperWidth * smallParcelPaperHeight)
}

func totalPaperNeeded(totalPaperNeedSize int) int {
	totalPaper := totalPaperNeedSize / (paperSizeWidth * paperSizeHeight)
	if totalPaperNeedSize%(paperSizeWidth*paperSizeHeight) != 0 {
		totalPaper++
	}
	return totalPaper
}

func needMore(totalPaperNeededSize int) bool {
	return totalPaperNeededSize%(paperSizeWidth*paperSizeHeight) != 0
}

func updateTotalPaperNeeded(totalPaperNeeded *int, needMore bool) {
	if needMore {
		*totalPaperNeeded++
	}
}

func totalCost(totalPaperNeeded, highLevel, entryLevel int) int {
	bigParcelCount := highLevel
	smallParcelCount := entryLevel

	return (totalPaperNeeded * paperCost) + (bigParcelCount * bigParcelCost) + (smallParcelCount * smallParcelCost)
}

func main() {
	var position string
	var highLevel, entryLevel, totalEmployeesVal, totalPaperNeededVal, totalCostVal int
	var totalPaperSize int

	employeeList := []string{}

	for {
		fmt.Scanln(&position)
		if position == "#" {
			break
		}
		employeeList = append(employeeList, position)
	}

	for _, position := range employeeList {
		checkEmployees(position, &highLevel, &entryLevel)
	}

	totalEmployeesVal = totalEmployees(highLevel, entryLevel)

	totalPaperSize = totalPaperNeededSize(highLevel, entryLevel)

	totalPaperNeededVal = totalPaperNeeded(totalPaperSize)

	updateTotalPaperNeeded(&totalPaperNeededVal, needMore(totalPaperSize))

	totalCostVal = totalCost(totalPaperNeededVal, highLevel, entryLevel)

	fmt.Printf("Vcorrp membutuhkan anggaran sebesar: Rp %d, untuk memberikan parcel kepada %d pegawai.\n", totalCostVal, totalEmployeesVal)
	fmt.Printf("Yang terdiri dari %d bigParcel dan %d smallParcel\n", highLevel, entryLevel)
}