package main

import (
	"bufio"
	"fmt"
	"strconv"
    "log"
	"os"
	"strings"
	
)
//Struct for packages
type Package struct{
	 name string
	 installDate string
	 lastUpdate string
	 numUpdates int
	 removalDate string
}

	var installedPack int
	var removedPack int
	var upgradedPack int
	var currentPack int
	

func main() {
	fmt.Println("Pacman Log Analyzer")

	if len(os.Args) < 2 {
		fmt.Println("You must send at least one pacman log file to analize")
		fmt.Println("usage: ./pacman_log_analizer <logfile>")
		os.Exit(1)
	}
	// Your fun starts here.
	
	file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	//var text []string
	var lines []string
	packs := map[string]Package{}

	//Read File
    for scanner.Scan() {
		
		var instDate string
		var upgradeCont int
		var lstUp string

		lines=strings.Split(scanner.Text()," ")
		if lines[3]=="installed" {
			//fmt.Println(lines[0])
			packs[lines[4]]=Package{lines[4],lines[0]+" "+lines[1]," - ",0," - "}
			installedPack++

		}else if lines[3] =="upgraded"{
			instDate= packs[lines[4]].installDate
			upgradeCont=packs[lines[4]].numUpdates
			upgradeCont=upgradeCont+1
			packs[lines[4]]=Package{lines[4],instDate,lines[0]+" "+lines[1],upgradeCont," - "}
			upgradedPack++
		}else if lines[3]=="removed"{
			instDate = packs[lines[4]].installDate
			upgradeCont=packs[lines[4]].numUpdates
			lstUp =packs[lines[4]].lastUpdate
			packs[lines[4]]=Package{lines[4],instDate,lstUp,upgradeCont,lines[0]+" "+lines[1]}
			removedPack++
		}
	}
	
	currentPack :=installedPack-removedPack
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	finalInstalled := strconv.Itoa(installedPack)
	finalRemoved := strconv.Itoa(removedPack)
	finalUpgraded := strconv.Itoa(upgradedPack)
	finalCurred := strconv.Itoa(currentPack)
	

	//Write file
	f, err:=os.Create("packages_report.txt")
    if err != nil {
        fmt.Println(err)
        return
	}
	l, err := f.WriteString("Pacman Packages Report\n-----------------------\n - Installed packages  : "+finalInstalled+"\n - Removed packages    : "+finalRemoved+"\n - Upgraded packages   : "+finalUpgraded+"\n - Current installed   : "+finalCurred+"\n")
	_, err = f.WriteString(" List of packages \n ----------------\n")

	for _,pack:=range packs{
		_, err = f.WriteString("- Package Name       : "+ pack.name+"\n - Install date      : "+ pack.installDate+"\n - Last update date  : "+pack.lastUpdate+"\n - How many updates  : "+strconv.Itoa(pack.numUpdates)+"\n - Removal date      : "+ pack.removalDate+"\n")

	}



	if err != nil {
        fmt.Println(err)
        f.Close()
        return
	}

	fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }

}
