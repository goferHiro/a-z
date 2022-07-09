package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string `json:"coursename"` //aliases
	Price int     
  Password string `json:"-"` //dont reflect 
	Tags  []string `json:"tags,omitempty"` //omit empty removes null
}


func EncodeJson() {
	courses := []course{
		{"React", 120, "1234",[]string{"FrontEnd", "Web App"}},
		{"Node", 1220,"1234", []string{"BackEnd", "Web App"}},
	}
	//finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses,"","\t") //prefix in every line is "" 
	try(err)
	fmt.Printf("%s",finalJson)
}

func main() {
	fmt.Println("JSON")
  EncodeJson()
  DecodeJSON()
}

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func DecodeJSON(){
  reqJSON := []byte(
    `{
        "coursename": "React",
        "Price": 120,
        "tags": [
            "FrontEnd",
            "Web App"
        ]
    }
`) 
   var sampleCourse course

   checkValid := json.Valid(reqJSON)
   if checkValid{
     fmt.Println("Valid JSON")
     json.Unmarshal(reqJSON,&sampleCourse)
     fmt.Printf("%#v\n",sampleCourse)
   }else{
     fmt.Println("JSON invalid")
   }

  //Alternate loading json to key-value pair
  var myData map[string]interface{} //not sure which datatypes

  json.Unmarshal(reqJSON,&myData)
   fmt.Printf("%#v\n",myData)
  for k,v:=range myData{
    fmt.Printf("Key %v Val %v Type %T\n",k,v,v)
  }
  //no need to convert to string
}