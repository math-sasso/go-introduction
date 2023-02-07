package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitorings = 5
const delay = 5

func main() {
	showIntroduction()
	for {

		command := readCommand()
		checkCommand(command)
	}

}

func showIntroduction() {
	var nome string = "Douglas"
	idade := 26
	var versao float32 = 1.1
	fmt.Println("Olá ", nome, "sua idade é", idade)
	fmt.Println("Versao ", versao)

	fmt.Println("O tipo da variavel nome é ", reflect.TypeOf(nome))
}

func readCommand() int {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit Program")

	var command int
	fmt.Scanf("%d", &command)
	fmt.Println("O comando escolhido foi", command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring")
	sites := readSitesFromFile()

	for i := 0; i < monitorings; i++ {
		for _, site := range sites {
			testSite(site)

		}
		time.Sleep(delay * time.Second)
	}

}

func testSite(site string) {
	ans, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	if ans.StatusCode == 200 {
		fmt.Println("Site", ans.Request.URL, "loaded correctly")
		registerLog(site, true)
	} else {
		fmt.Println("Site", ans.Request.URL, "with problems", ans.StatusCode)
		registerLog(site, false)
	}
}

func checkCommand(command int) {
	// if command == 1 {
	// 	fmt.Println("Monitoring")
	// } else if command == 2 {
	// 	fmt.Println("Show logs")
	// } else if command == 3 {
	// 	fmt.Println("Exit")
	// } else {
	// 	fmt.Println("Command not valid")
	// }

	switch command {
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	case 3:
		fmt.Println("Exit")
		os.Exit(0)
	default:
		fmt.Println("Command not valid")
		os.Exit(-1)
	}
}

func readSitesFromFile() []string {
	sites := []string{}
	file, err := os.Open("sites.txt")
	// file, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		// return
	}

	reader := bufio.NewReader(file)
	for {
		// content := file.Read()

		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}

	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	time_now := time.Now().Format("02/01/2006 15:04:05")
	file.WriteString(time_now + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {
	fmt.Println("Show logs")
	file, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
