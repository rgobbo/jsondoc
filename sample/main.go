package main

import (
	"github.com/rgobbo/jsondoc"
	"log"

)

func main() {
	j := jsondoc.NewJson()
	str := `{"name": "test", "age": 30}`
	err := j.LoadFromString(str)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("name=" ,j.GetString("name"))
	log.Println("age=", j.GetInt("age"))

	j2 := jsondoc.NewJson()
	j2.LoadFromFile("sample/test.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("name=" ,j2.GetString("name"))
	log.Println("count=", j2.GetInt("count"))
	d, err := j2.GetDate("DD/MM/YYYY","date")
	if err != nil{
		log.Println(err)
	}
	log.Println("date=", d)

}
