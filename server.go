package main

import (
	"net/http"
	"fmt"
	"html/template"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/astaxie/beego/session"
	"reflect"
)

const(// 基本配置
	mongoUrl = "mongodb://localhost:27017"// 数据库地址
	PortNumber = ":9000"// 端口号
	Database = "forum"// 数据库名
	User = "user"// 用户集合
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

/*type mgoMap struct{
	Arr []testMgo
}*/
var globalSessions *session.Manager

func init() {
	var err error
	sessionConfig := &session.ManagerConfig{
		CookieName: "jigsessionid",
		Gclifetime: 3600,
	}
	globalSessions, err = session.NewManager("memory", sessionConfig)
	if err == nil {}
	go globalSessions.GC()
}

func static(){// 静态目录
	http.Handle("/js/", http.FileServer( http.Dir("public") ))
	http.Handle("/css/", http.FileServer( http.Dir("public") ))
	http.Handle("/images/", http.FileServer( http.Dir("public") ))
	http.Handle("/dist/", http.FileServer( http.Dir("") ))
}

func routing(){// 路由
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

func mgoTest(){// mgo 测试
	session, err := mgo.Dial(mongoUrl)
	/*defer session.Close()
	session.SetMode(mgo.Monotonic, true)*/
	if err != nil{
		panic(err)
	}
	collection := session.DB(Database).C(User)
	
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

func (this *MainController) Get() {
		var intcountintsess:= this.StartSession()
		count:= sess.Get("count")
		if count == nil {
			intcount = 0
		} else {
			intcount = count.(int)
		}
		intcount = intcount + 1
		sess.Set("count", intcount)
		this.Data["Username"] = "astaxie"
		this.Data["Email"] = "astaxie@gmail.com"
		this.Data["Count"] = intcount
		this.TplNames = "index.tpl"
	}

func pageInterface(){// 接口
	//var sess *session.MemSessionStore
	http.HandleFunc("/api/Login/", func(w http.ResponseWriter, r *http.Request){// 登录接口
		w.Header().Set("content-type", "application/json")

		/*sess, err := globalSessions.SessionStart(w, r)
		fmt.Println("type:", reflect.TypeOf(sess))

		defer sess.SessionRelease(w)
		if sess.Get("username") != nil {
			out :=  &Login{"已登录", 2}
			jsonData, err := json.Marshal( out )
			if err == nil {
				w.Write(jsonData)
			}
			return
		}
		sess.Set("username", r.FormValue("name"))
		uesename := sess.Get("username")
		fmt.Printf("%s\n", uesename)*/

		name, password := r.FormValue("name"), r.FormValue("password")
		println(name, password)
		out := &Login{"成功", 1}
		jsonData, err := json.Marshal(out)
		if err == nil {
			w.Write(jsonData)
		}
	})


	http.HandleFunc("/api/registered/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("content-type", "application/json")

		// sess, err := globalSessions.SessionStart(w, r)
		// defer sess.SessionRelease(w)
		//username := sess.Get("username")
		//fmt.Printf("%s", username)
		//if err == nil {}
	})
}

func main() {

	static()// 静态目录
	routing()// 路由
	// mgoTest()// 数据库测试
	pageInterface()// 接口
	fmt.Printf("启动成功\n")
	http.ListenAndServe(PortNumber, nil)
}
