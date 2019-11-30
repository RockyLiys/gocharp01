package httpweb
/*
	做一个简单的 ftp
*/
import (
	"fmt"
	"net/http"
)

func Webftp() {
	http.Handle("/", http.FileServer(http.Dir("/Users/oyo00451/Documents/")))
	fmt.Println("请访问ftp: http://localhost:8008")
	http.ListenAndServe("0.0.0.0:8008", nil)
}
