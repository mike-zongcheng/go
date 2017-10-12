package main

import (
	"net/http"
	"fmt"
	"html/template"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const(
	mongoUrl = "mongodb://localhost:27017"
)

var (
	mogSessin *mgo.Session
	dataBase = "mydb"
)

type Login struct{
	Text string
	Id int
}

type testMgo struct{
	//ID bson.ObjectId `bson:"_id"`
	NAME string `bson:"name"`
	NUM int `bson:"num"`
	AUTHOR string `bson:"author"`
}

type mgoMap struct{
	Arr []testMgo
}

func routing(){
	http.Handle("/js/", http.FileServer( http.Dir("content") ))
	http.Handle("/css/", http.FileServer( http.Dir("content") ))
	http.Handle("/images/", http.FileServer( http.Dir("content") ))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {// 首页与404路由
        if r.URL.Path == "/" {
        	t, err := template.ParseFiles("view/home/index.html");
			if err == nil {
				t.Execute(w, nil)
			}
        	return
        }
        t, err := template.ParseFiles("view/404/index.html")
        if err == nil {
        	t.Execute(w, nil)
        }
    })


	http.HandleFunc("/search/", func(w http.ResponseWriter, r*http.Request){
		w.Write([]byte("这是搜索"))
	})

	http.HandleFunc("/api/test/", func(w http.ResponseWriter, r*http.Request){
		w.Header().Set("content-type", "application/json")

		name, password := r.FormValue("name"), r.FormValue("password")

		println(name, password)

		out := &Login{"成功", 15}
		jsonData, err := json.Marshal(out)
		if err == nil {
			w.Write(jsonData)
		}
	})

}

func mgoTest(){
	session, err := mgo.Dial(mongoUrl)
	/*defer session.Close()
	session.SetMode(mgo.Monotonic, true)*/
	if err != nil{
		panic(err)
	}
	db := session.DB("godb")
	collection := db.C("user")
	
	result := testMgo{} // 查询数据
	/*collection.Find( bson.M{"name":"海错图"} ).One(&result)
	fmt.Printf("内容：%s",result.AUTHOR)*/

	/*downData := &testMgo{"悲惨世界", 60, "雨果"}// 插入数据
	err = collection.Insert( &testMgo{"如果宅", 40, "有时右逝"}, downData )
	if err != nil{
		panic(err)
	}*/

	var mgoArr mgoMap
	iter := collection.Find( bson.M{"name":"悲惨世界"} ).Iter();
	for iter.Next(&result){
		mgoArr.Arr = append(mgoArr.Arr, result)
	}
	fmt.Printf("%s", mgoArr.Arr)

}

func main() {
	routing()
	fmt.Printf("启动成功\n")
	mgoTest()
	http.ListenAndServe(":9000",nil)
}
