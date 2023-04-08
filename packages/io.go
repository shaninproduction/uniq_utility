package input_output

import (
	"fmt"
	"os"
)

func HelpMessage() {

	fmt.Println("`uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]`")
	fmt.Println("Использование")
	fmt.Print("Параметры\n`-с` - подсчитать количество встречаний строки во входных данных.Вывести это число перед строкой отделив пробелом." +
		"`-d` - вывести только те строки, которые повторились во входных данных.\n" +
		"`-u` - вывести только те строки, которые не повторились во входных данных.\n" +
		"`-f num_fields` - не учитывать первые `num_fields` полей в строке.\nПолем в строке является непустой набор символов отделённый пробелом.\n" +
		"`-s num_chars` - не учитывать первые `num_chars` символов в строке.\nПри использовании вместе с параметром `-f` учитываются первые символы\nпосле `num_fields` полей (не учитывая пробел-разделитель после\nпоследнего поля)." +
		"`-i` - не учитывать регистр букв.\n")
}

func WithoutParams(data []string) []string {
	uniq_strings := make([]string, 0)
	flag := true

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(uniq_strings); j++ {
			if data[i] == uniq_strings[j] {
				flag = false
				break
			}
		}

		if flag {

			uniq_strings = append(uniq_strings, data[i])

		} else {

			flag = true

		}
	}

	return uniq_strings
}

func PrintStdIn(data []string, args []string) {

	file, err := os.Open(args[0])

	if err != nil {
		fmt.Println("Error: Bad input")
		os.Exit(1)
	}

	defer file.Close()

	data = GetData(file)

	uniq_strings := MainLogic(data)

	for i := 0; i < len(uniq_strings); i++ {
		fmt.Println(uniq_strings[i])
	}

}
func PrintFromStdIn(data []string) {

	data = GetData(os.Stdin)

	uniq_strings := MainLogic(data)

	for i := 0; i < len(uniq_strings); i++ {
		fmt.Println(uniq_strings[i])
	}

}

func PrintInFile(data []string, args []string) {

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

	data = GetData(file_in)

	uniq_strings := MainLogic(data)

	for j := 0; j < len(uniq_strings); j++ {

		fmt.Fprintln(file_out, uniq_strings[j])

	}
}
