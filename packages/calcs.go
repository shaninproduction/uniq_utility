package calcs

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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

func MainLogic(data []string) []string {

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

	if *i == true {

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
