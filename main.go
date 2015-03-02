package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/opesun/goquery"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
)

func main() {
	args := os.Args[1:]
	if len(args[0:]) != 0 {
		var x goquery.Nodes
		if args[0] == "all" {
			x, _ = goquery.ParseUrl("http://referats.yandex.ru/referats/?t=astronomy+geology+gyroscope+literature+marketing+mathematics+music+polit+agrobiologia+law+psychology+geography+physics+philosophy+chemistry+estetica")
		} else {
			x, _ = goquery.ParseUrl("http://referats.yandex.ru/referats/?t=" + args[0])
		}
		text := x.Find(".referats__text").Text()
		title_regexp, _ := regexp.Compile("Тема: «(.*)»")
		title := title_regexp.FindStringSubmatch(text)[1]
		c_regexp, _ := regexp.Compile("»(.*)")
		c := c_regexp.FindStringSubmatch(text)[1]

		referat := title + "\n\n" + c

		if len(args[1:]) != 0 {
			switch args[1] {
			default:
				fmt.Println("Неверная опция!")
			case "copy":
				// copy referat to clipboard
				clipboard.WriteAll(referat)
				fmt.Println("Реферат на тему \"" + title + "\" скопирован в буфер обмена")
			case "file":
				usr, err := user.Current()
				if err != nil {
					fmt.Println(err)
				}
				e := ioutil.WriteFile(usr.HomeDir+"/"+title, []byte(referat), 0777)
				if e != nil {
					panic(e)
				}
				fmt.Println("Реферат записан в файл \"" + usr.HomeDir + "/" + title + "\"")
			}
		} else {
			fmt.Println(referat)
		}
	} else {
		fmt.Println("Yandex Referats by CubexX\n\nАстрономия - astronomy\nГеология - geology\nГироскопия - gyroscope\nЛитература - literature\nМаркетинг - marketing\nМатематика - mathematics\nМузыковедение - music\nПолитологии - polit\nПочвоведение - agrobiologia\nПравоведение - law\nПсихология - psychology\nСтрановедение - geography\nФизика - physics\nХимия - chemistry\nЭстетика - estetica\n\nНапример: referat physics+music+law [print|file|copy]")
	}
}
