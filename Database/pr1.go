package main()
import fmt
import database/sql
import _"github.com/mattn/go-sqlite3"

func main()
{
con, e1:= sql.open("sqlite3", "st1.db")
if e1!= nil{
	fmt.Println("issue", e1)
	return
}
fmt.Println("connected")
defer con.Close()
}

sql:= "create table student(rno int primary key, name text)"
res,e2 := con.Exec(sql)
if e2!= nil{
	fmt.Println("creating issue",e2)
	return
}
fmt.Println("connected")
}