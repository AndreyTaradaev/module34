package main

import (
  "fmt"
  "regexp"
  "strings"
)

func main() {
  cities := []string{
     "Moscow",
     "Washington",
     "New-York",
     "Kiev",
     "Vitebsk",
     "Kishinev",
     "Vladivostok",
  }

  fmt.Println(cities[0])

  for i := 1; i < len(cities); i++ {
     p, n := cities[i-1], cities[i]
     if !isCorrectCityName(p, n) {
        fmt.Println(cities[i], "-", "WRONG!")
        break
     }

     fmt.Println(cities[i], "-", "OK!")
  }
  
  for i := 1; i < len(cities); i++ {
	p, n := cities[i-1], cities[i]
	
	if !isCorrectCityName1(p, n) {
	   fmt.Println(cities[i], "-", "WRONG!")
	   break
	}

	fmt.Println(cities[i], "-", "OK!")
 }
}


func isCorrectCityName1(prev, next string) bool {
 	re := regexp.MustCompile(`^([A-Z])[a-zA-Z-]+([a-z])$`)
	submatches := re.FindAllStringSubmatch(prev, -1)
//fmt.Println(submatches[0][2])
parse := "(?i)"+ string(next[0])
reg := regexp.MustCompile(parse)
return reg.MatchString(submatches[0][2])
} // (

func isCorrectCityName(prev, next string) bool {
prevlen := len(prev)

if prevlen ==0 {
	return false
}
if(len(next)==0){
	return false
}


end := prev[prevlen-1]
if strings.ToLower(string(end)) == strings.ToLower(string(next[0])){
	return true
}
 return false // TODO: implement me
}