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
	PortNumber = ":9000"
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

func static(){
	http.Handle("/js/", http.FileServer( http.Dir("public") ))
	http.Handle("/css/", http.FileServer( http.Dir("public") ))
	http.Handle("/images/", http.FileServer( http.Dir("public") ))
	http.Handle("/dist/", http.FileServer( http.Dir("") ))
}

func routing(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {// 首页与404路由
        if r.URL.Path == "/" {
        	t, err := template.ParseFiles("index.html");
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

}

func mgoTest(){
	session, err := mgo.Dial(mongoUrl)
	/*defer session.Close()
	session.SetMode(mgo.Monotonic, true)*/
	if err != nil{
		panic(err)
	}
	collection := session.DB("godb").C("user")
	
	 result := testMgo{} // 查询数据
	collection.Find( bson.M{"name":"海错图"} ).One(&result)
	//fmt.Printf("内容：%s",result.AUTHOR)

	/*downData := &testMgo{"悲惨世界", 60, "雨果"}// 插入数据
	err = collection.Insert( &testMgo{"如果宅", 40, "有时右逝"}, downData )
	if err != nil{
		panic(err)
	}*/

	/*var mgoArr mgoMap
	iter := collection.Find( bson.M{"name": bson.M{"$regex":"海错"} } ).Iter();
	for iter.Next(&result){
		mgoArr.Arr = append(mgoArr.Arr, result)
	}
	fmt.Printf("%s", mgoArr.Arr)
*/
	// collection.UpdateAll( bson.M{"name":"悲惨世界"}, bson.M{ "$set": bson.M{"author":"雨果"} } )
	// collection.RemoveAll( bson.M{"name":"悲惨世界"} )
}

func pageInterface(){
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

func main() {
	static()// 静态目录
	routing()// 路由
	mgoTest()// 数据库测试
	pageInterface()// 接口
	fmt.Printf("启动成功\n")
	http.ListenAndServe(PortNumber, nil)
}
