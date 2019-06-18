package console

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/logrusorgru/aurora"
)

// ReadString 从命令行中读取用户输入
func ReadStringVar(v *string, defaultV string, pt string) error {
	prompt(pt, defaultV)
	input, err := readLine()
	if err != nil {
		return err
	}
	if input == "" {
		input = defaultV
	}
	*v = input
	return nil
}

func ReadBoolVar(v *bool, defaultV bool, pt string) error {
	prompt(fmt.Sprintf("%s (Y/N)", pt), defaultV)
	input, err := readLine()
	if err != nil {
		return err
	}
	if strings.ToUpper(input) == "Y" {
		*v = true
	} else if input == "" {
		*v = defaultV
	} else {
		*v = false
	}
	return nil
}

func ReadIntVar(v *int, defaultV int, pt string) error {
	prompt(pt, defaultV)
	input, err := readLine()
	if err != nil {
		return err
	}

	if input == "" {
		*v = defaultV
		return nil
	}

	i, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	*v = i

	return nil
}

func prompt(str string, defaultV interface{}) {
	var suffix string
	v := reflect.ValueOf(defaultV)
	switch v.Kind() {
	case reflect.String:
		if v2 := v.String(); v2 != "" {
			suffix = fmt.Sprintf("[%s]", v2)
		}
		break
	case reflect.Bool:
		v2 := v.Bool()
		if v2 {
			suffix = fmt.Sprintf("[%s]", "Y")
		} else {
			suffix = fmt.Sprintf("[%s]", "N")
		}
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		suffix = fmt.Sprintf("[%d]", v.Int())
		break
	default:
		break
	}

	fmt.Println(aurora.BgBlack(str), aurora.Cyan(suffix).Italic())
}

func readLine() (string, error) {
	in := bufio.NewReader(os.Stdout)
	v, err := in.ReadString('\n')
	v2 := strings.Trim(v, " \n")
	return v2, err
}
