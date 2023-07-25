package main

import (
	"bufio"
	"fmt"
	"goreloaded/ReloadFuncs"
	"os"
	"regexp"
	"strings"
)

func main() {

	if len(os.Args) > 3 {
		fmt.Println("the input must be two arguments <fileName.txt> <result.txt>")
		return
	}

	if len(os.Args) <= 2 {
		fmt.Println("the input must be two arguments <fileName.txt> <result.txt>")
		return
	}

	input := os.Args[1]
	output_file := os.Args[2]
	data, _ := os.ReadFile(input)
	string_input := string(data)
	output1 := ReloadFuncs.SmallToCap(string_input)
	output2 := ReloadFuncs.HexNumToDec(output1)
	output3 := ReloadFuncs.BinNumToDec(output2)
	output4 := ReloadFuncs.LowToUp(output3)
	output5 := ReloadFuncs.CapToLow(output4)
	output6 := ReloadFuncs.ChangeAToAn(output5)

	string_list := strings.Split(string(output6), " ")
	for idx, word := range string_list {
		if word == "(low," {
			ReloadFuncs.NumFromLow(string_list, idx)
		} else if word == "(up," {
			ReloadFuncs.NumberFromUp(string_list, idx)
		} else if word == "(cap," {
			ReloadFuncs.NumToCapitalize(string_list, idx)
		}
	}
	output := ""
	for _, word := range string_list {
		if word != "" {
			output += word
			output += " "
		}
	}
	rePunct := regexp.MustCompile(`\s+([[:punct:]]{1,})\s*`)
	output = rePunct.ReplaceAllStringFunc(output, ReloadFuncs.ModPunc)

	reQuote := regexp.MustCompile(`\'[^']+\'`)
	output = reQuote.ReplaceAllStringFunc(output, ReloadFuncs.ModQuote)

	if output[len(output)-1] == ' ' {
		output = output[0 : len(output)-1]
	}
	f_output, err := os.Create(output_file)
	if err != nil {
		fmt.Println("Error while creating the file!")
		return
	}
	w := bufio.NewWriter(f_output)
	_, err = w.WriteString(output)
	if err != nil {
		fmt.Println("Error while writing the file!")
		return
	}

	w.Flush()
}
