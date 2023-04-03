package io

import (
	"bufio"
	"fmt"
	"os"
)

func print_log(data []string, args []string) {

	file, err := os.Open(args[0])

	if err != nil {
		fmt.Println("Error: Bad input")
		os.Exit(1)
	}

	defer file.Close()

	data = get_data(file)

	uniq_strings := main_logic(data)

	for i := 0; i < len(uniq_strings); i++ {
		fmt.Println(uniq_strings[i])
	}

}
func print_from_stdin(data []string) {

	data = get_data(os.Stdin)

	uniq_strings := main_logic(data)

	for i := 0; i < len(uniq_strings); i++ {
		fmt.Println(uniq_strings[i])
	}

}

func print_in_file(data []string, args []string) {

	file_in, err_in := os.Open(args[0])

	if err_in != nil {
		fmt.Println("Error: Bad input")
		return
	}

	defer file_in.Close()

	file_out, err_out := os.Create(args[1])

	if err_out != nil {
		fmt.Println("Error: Bad input")
		return
	}

	defer file_out.Close()

	data = get_data(file_in)

	uniq_strings := main_logic(data)

	for j := 0; j < len(uniq_strings); j++ {

		fmt.Fprintln(file_out, uniq_strings[j])

	}
}

func get_data(file *os.File) []string {

	scanner := bufio.NewScanner(file)

	var strings []string
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}
