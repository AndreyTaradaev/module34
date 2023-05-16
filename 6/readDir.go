package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main (){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку данных: ")
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)	  
	}	
	str = strings.TrimSpace(str)
	files, err := ioutil.ReadDir(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)	  
	}

	f, err := os.OpenFile("ListDir.txt", os.O_WRONLY|os.O_CREATE, os.ModePerm)

	if(err!= nil){
		fmt.Println(err)
		os.Exit(-3)	  
	}

	writer := bufio.NewWriter(f)

for _, v := range files {
	writer.Write([]byte(v.Name()) ) 
	writer.Write([]byte("\t") ) 
	writer.WriteString( "("+ fmt.Sprintf("%d",v.Size())+")\t") 

	if(v.IsDir()) {
		writer.WriteString("[DIRECTORY]\n")
	}else{
		writer.WriteString("[FILE]\n")
	}
	writer.Flush()
}



	/* writer.Write([]byte("Hello"))
	writer.Flush()
	writer.Write([]byte(", "))
	writer.Write([]byte("World"))
	writer.Write([]byte("!!!")) // break
	writer.Flush()



	/* for _, file := range files {
	  fmt.Println(file.Name(), file.IsDir())
	} */
 

}