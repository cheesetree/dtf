package test

import(
	"gorm.io/gorm"
	"github.com/cheesetree/oracle"
	"fmt"
)

func main(){
	db, err := gorm.Open(oracle.Open("oadmin/123456@10.116.0.41:1521/ORCL"), &gorm.Config{})
    if err != nil {
        fmt.Println(err)
    }
}