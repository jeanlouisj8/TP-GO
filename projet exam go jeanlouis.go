package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

type task struct {
	Description string
	Done bool
}


//verif constante POST

	const (
		MethodePost  = "POST"	

	)



	func main() {

varglobtask := [] string{("t0 faire courses" id:false), "t1 payer factures" id:false) "Done" ==false}

}

//header envoie client
rw.WriteHeader(http.StatusOK)
Write([]byte) (int, error) chaine:= "ma chaine"
enBytes:= []byte(chaine)
}

		hande1 := func (w http.ResponseWriter, _ *http.Request) {
				io.WriteString(w, "hello from a hande1 #1!\n")
		}

		hande2 := func (w http.ResponseWriter, _ *http.Request) {
				io.WriteString(w, "hello from a hande2 #2!\n")
		}	

		hande3 := func (w http.ResponseWriter, _ *http.Request) {
				io.WriteString(w, "hello from a hande3 #3!\n")
		}

		http.HandleFunc("/", list, hande1)
		http.HandleFunc("/done", done, hande2)
		http.HandleFunc("/add", add, hande3)
	}

	//hello hande, server web

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "ListenAndServe\n")
	}

}

//function list
func  list(rw http.ResponseWriter, _ *http.Request)


http.HandleFunc("/", /)
log.Fatal(http.ListenAndServe(":8080", nil))
}
http.HandleFunc("/done", done)
log.Fatal(http.ListenAndServe(":8080", nil))
}
http.HandleFunc("/add", add)
log.Fatal(http.ListenAndServe("8080", nil))


body, err :=
ioutil.ReadAll(r.Body) if err !=
nil {
fmt.Printf("Error reading body: %v", err)
http.Error(
rw,
"can't read body", http.StatusBadRequest,
)
return
