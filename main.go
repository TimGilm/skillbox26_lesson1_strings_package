package main

import (
	"fmt"
	"strings"
)

func main() {
	//перевод строк в нижний и верхний регистр
	str := "Some String"

	strInLowerCase := strings.ToLower(str)
	fmt.Println(strInLowerCase)
	strInUpperCase := strings.ToUpper(str)
	fmt.Println(strInUpperCase)

	//функции поиска подстроки
	haveSome := strings.HasPrefix(str, "Some")
	fmt.Println(haveSome) //true
	haveIng := strings.HasSuffix(str, "ing")
	fmt.Println(haveIng) //true
	haveSt := strings.Contains(str, "St")
	fmt.Println(haveSt) //true
	haveS := strings.Count(str, "S")
	fmt.Println(haveS) //2

	//определение длины строки
	str = "string"
	fmt.Println(len(str)) //6
	str = "строка"
	fmt.Println(len(str)) //12
}
