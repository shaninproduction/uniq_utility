package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	uniq_strings := MainLogic(data, FLags)

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

func GetData(file *os.File) []string {

	scanner := bufio.NewScanner(file)

	var strings []string
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}

func BadFlags(flags []string) bool {
	counter := 0

	for i := 0; i < len(flags); i++ {
		if flags[i] == "-u" || flags[i] == "-c" || flags[i] == "-d" {
			counter++
		}
	}

	if counter > 1 {
		return true
	}

	return false
}

func CflagAlg(data []string, copied_data []string) []string {

	uniq_strings := make([]string, 0)
	indexes_values := make(map[string]int)
	flag := true

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(uniq_strings); j++ {
			if data[i] == uniq_strings[j] {
				flag = false
			}
		}

		if flag {

			uniq_strings = append(uniq_strings, data[i])
			indexes_values[data[i]] = 1

		} else {

			flag = true
			indexes_values[data[i]] += 1

		}
	}

	result := make([]string, 0)

	if *i == true {
		for i := 0; i < len(uniq_strings); i++ {
			for j := 0; j < len(copied_data); j++ {
				if strings.ToLower(copied_data[j]) == uniq_strings[i] {

					uniq_strings[i] = copied_data[j]
					break

				}
			}
		}

		for i := 0; i < len(uniq_strings); i++ {
			result = append(result, strconv.Itoa(indexes_values[strings.ToLower(uniq_strings[i])])+" "+uniq_strings[i])
		}

		return result
	}

	for i := 0; i < len(uniq_strings); i++ {
		result = append(result, strconv.Itoa(indexes_values[uniq_strings[i]])+" "+uniq_strings[i])
	}

	return result
}

func DflagAlg(data []string) []string {

	not_uniq_strings := make([]string, 0)
	flag := true

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {

			if data[i] == data[j] {

				flag = true
				for k := 0; k < len(not_uniq_strings); k++ {
					if data[i] == not_uniq_strings[k] {
						flag = false
						break
					}
				}

				if flag == true {
					not_uniq_strings = append(not_uniq_strings, data[i])
					break
				}

			}

		}
	}

	return not_uniq_strings
}

func UflagAlg(data []string) []string {

	uniq_strings := make([]string, 0)
	not_uniq_strings := make([]string, 0)
	var flag bool

	for i := 0; i < len(data); i++ {

		flag = false
		for j := i + 1; j < len(data)-1; j++ {
			if data[i] == data[j] {
				flag = true
				break
			}
		}

		if flag == true {
			not_uniq_strings = append(not_uniq_strings, data[i])

		} else {

			for k := 0; k < len(not_uniq_strings); k++ {
				if data[i] == not_uniq_strings[k] {
					flag = true
					break
				}
			}

			if flag == false {
				uniq_strings = append(uniq_strings, data[i])
			}

		}
	}

	return uniq_strings

}

func MainLogic(data []string, flags FLags) []string {

	copied_data := make([]string, len(data))

	copy(copied_data, data)

	uniq_strings := make([]string, 0)

	if *f != 0 { // Поведение при флаге -f

		for i := 0; i < len(data); i++ {

			temp := strings.Split(data[i], " ")

			if *f > len(temp) {

				data[i] = ""
				copied_data[i] = ""

			} else {

				temp = temp[*f:]

				data[i] = strings.Join(temp, " ")
				copied_data[i] = strings.Join(temp, " ")
			}
		}

	}

	if *s != 0 {
		for i := 0; i < len(data); i++ {

			if *s > len(data[i]) {
				data[i] = ""
			} else {

				data[i] = data[i][*s:]
			}
		}
	}

	if *i == true { // Поведение при флаге -i

		for i := 0; i < len(data); i++ {
			data[i] = strings.ToLower(data[i])
		}
	}

	if *c == true { // Поведение при флаге -c

		uniq_strings = CflagAlg(data, copied_data)

	} else if *d == true { // Поведение при флаге -d

		uniq_strings = DflagAlg(data)

	} else if *u == true { // Поведение при флаге -u

		uniq_strings = UflagAlg(data)

	} else { // Поведение без флагов

		uniq_strings = WithoutParams(data)
	}

	if flags.c == true {

		if *c == false {
			for i := 0; i < len(uniq_strings); i++ {
				for j := 0; j < len(copied_data); j++ {
					if strings.ToLower(copied_data[j]) == uniq_strings[i] {

						uniq_strings[i] = copied_data[j]
						break

					}
				}
			}

		}

	}

	return uniq_strings
}

var (
	c = flag.Bool("c", false, "Подсчитывает кол-во встречаний строки во входных данных.")
	d = flag.Bool("d", false, "Выводит только те строки, который повторились во входных данных")
	u = flag.Bool("u", false, "Выводит только те строки, которые не повторились во входных данных")
	i = flag.Bool("i", false, "Игнорирует регистр")
	f = flag.Int("f", 0, "Игнорирует num полей в строке")
	s = flag.Int("s", 0, "Игнорирует num символов в строке")
)

type FLags struct {
	c bool
	d bool
	u bool
	i bool
	f int
	s int
}

func main() {

	Flags := Flags{
		c: c,
		d: d,
		u: u,
		i, i,
		f: f,
		s: s,
	}

	flag.Parse()

	args := flag.Args()

	// Список флагов
	flags := make([]string, 0)
	for i := 1; i < len(os.Args[1:])+1; i++ {

		if string(os.Args[i][0]) == "-" {
			flags = append(flags, os.Args[i])
		}
	}

	// Проверка не введены ли неправильные комбинации флагов
	if BadFlags(flags) == true {
		HelpMessage()
		return
	}

	var data []string

	if len(args) == 0 {

		// Если ввод с stdin
		PrintFromStdIn(data)

	} else if len(args) == 1 {

		// Вывод в лог
		PrintStdIn(data, args)

	} else if len(args) == 2 {

		// Вывод в файл
		PrintInFile(data, args)

	} else {

		// Сообщение при неправильном использовании скрипта
		HelpMessage()

	}
}
