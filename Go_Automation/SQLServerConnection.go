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
  servername := ""
  databasename := ""
  query := ""
  fmt.Println(sql.Drivers())
  db, err := sql.Open("sqlserver","odbc:server="+ servername + ";IntegratedSecurity = false;;database=" + databasename + ";app name=MyAppName")
  if err != nil {
  log.Fatal("Open connection failed:", err.Error())
  }
  fmt.Printf("Connected!\n")
  defer db.Close()

  tsql := fmt.Sprintf(query)
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
     /*multiple rows sets
       store := []map[string]interface{}{}
  for n := true; n; n = rows.NextResultSet() {
    cols, _ := rows.Columns()

    for rows.Next() {

        columns := make([]interface{}, len(cols))
        columnPointers := make([]interface{}, len(cols))
        for i, _ := range columns {
            columnPointers[i] = &columns[i]
        }

        if err := rows.Scan(columnPointers...); err != nil {
            log.Fatal(err)
        }
        m := make(map[string]interface{})
        for i, colName := range cols {
            val := columnPointers[i].(*interface{})
            m[colName] = *val
        }
        store = append(store, m)
    }
    fmt.Println(store)
  }
  js, _ = json.MarshalIndent(store, "", "  ")
     
     
     */
}
