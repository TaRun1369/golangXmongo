package golangxmongo

import (
	// Add your imports here
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/tarun1369/golangXmongo/controllers"

)
	


func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession()) // getsession returns pointer to s of mongodb session
	r.GET("/user/:id",uc.GetUser)
	r.POST("/user",uc.CreateUser,)
	r.DELETE("/user/:id",uc.DeleteUser)
	http.ListenAndServe("localhost:8080",r)

}

func getSession() *mgo.Session{
	// pointer of mongodb session
	s,err := mgo.Dial("mongodb://localhost:27017")  // s is pointer session jo get session se aaya
	if err != nil {
		panic(err)
	}
	return s
}