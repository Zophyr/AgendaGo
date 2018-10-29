package entity

import (
	"encoding/json"
	"fmt"
	"os"
	"utils"
)

type storage struct {
	path string 
}

func (s *storage) load(ptr interface{}) {
	file, err := os.Open(s.path)
	defer file.Close()
	if err != nil{
		fmt.Println(os.Stderr,"err:%s",err)
		return		
	}
	err := json.NewDecoder(file).Decode(ptr)
	if err!=nil{
		fmt.Println(os.Stderr,"err:%s",err)
		return
	}
}

func (s *storage) dump(ptr interface{}) {
	file, err := os.OpenFile(s.path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(os.Stderr,"err:%s",err)
		return
	}
	err := json.NewEncoder(file).Encode(ptr)
	if err != nil {
		fmt.Println(os.Stderr,"err:%s",err)
		return
	}
}