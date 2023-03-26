package main

import (
	"bufio"
	"crypto/md5"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func parseRemoteCSV() []map[string]interface{} {
	resp, err := http.Get("https://raw.githubusercontent.com/OtusTeam/highload/master/homework/people.csv")
	if err != nil {
		log.Fatal("cant send get request: ", err)
	}

	scanner := bufio.NewScanner(resp.Body)

	result := make([]map[string]interface{}, 0)

	for scanner.Scan() {
		firstLine := scanner.Text()
		values := make(map[string]interface{})

		fields := strings.Split(firstLine, ",")
		if len(fields) < 3 {
			log.Print("wrong format of input line, continue...")
			continue
		}
		firstSecond := strings.Split(fields[0], " ")
		if len(firstSecond) < 2 {
			log.Print("wrong format of input line, continue...")
			continue
		}
		values["first_name"] = firstSecond[0]
		values["second_name"] = firstSecond[1]
		ageString := fields[1]
		age, err := strconv.Atoi(ageString)
		if err != nil {
			log.Printf("wrong format of ages - %s, continue...", ageString)
			continue
		}
		values["age"] = age
		values["city"] = fields[2]
		result = append(result, values)
		pass, err := md5.New().Write([]byte(firstSecond[0]))
		if err != nil {
			log.Printf("cant create pass: %s", err.Error())
			continue
		}
		values["password"] = pass
	}
	return result
}
