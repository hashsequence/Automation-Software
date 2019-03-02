package main
import (
     "fmt"
     _ "github.com/denisenkom/go-mssqldb"
   "database/sql"
//   "net/url"
  "log"
)
/*

https://github.com/Microsoft/sql-server-samples/blob/master/samples/tutorials/go/crud.go
*/

func main() {
  fmt.Println("hello world")
  fmt.Println(sql.Drivers())
  db, err := sql.Open("sqlserver","odbc:server=ODSSQLCVSDEV;IntegratedSecurity = false;;database=cvsCareMarkPB;app name=MyAppName")
  if err != nil {
  log.Fatal("Open connection failed:", err.Error())
  }
  fmt.Printf("Connected!\n")
  defer db.Close()

  tsql := fmt.Sprintf("SELECT distinct Top(3) memberrecid ,firstname, lastname from members (nolock);")
  rows, _ := db.Query(tsql)
  cols, _ := rows.Columns()

  data := make(map[string]string)

  for _ , colname := range cols {
    fmt.Print(colname + " | ")
  }
  fmt.Print("\n")

  for rows.Next() {
      columns := make([]string, len(cols))
      columnPointers := make([]interface{}, len(cols))
      for i, _ := range columns {
          columnPointers[i] = &columns[i]
      }

      rows.Scan(columnPointers...)

      for i, colName := range cols {
          data[colName] = columns[i]
      }

      for _, colname := range cols {
        fmt.Print(data[colname] + " | ")

      }
      fmt.Print("\n")
   }
}
