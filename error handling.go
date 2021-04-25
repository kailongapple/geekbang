/*
import (
    "database/sql"
    "fmt"
 )
 
 func aaa() error {
    return sql.ErrNoRows
 }
 
 func bbb() error {
    return aaa()
 }
 
 func main() {
    err := bbb()
    if err != nil {
       fmt.Printf("got err, %+v\n", err)
    }
 }
 //Outputs:
 // got err, sql: no rows in result set

 */



/*
第三方库包装error方法
github.com/pkg/errors
*/
import (
    "database/sql"
    "fmt"
 
    "github.com/pkg/errors"
 )
 
 func aaa() error {
    return errors.Wrap(sql.ErrNoRows, "aaa failed")
 }
 
 func bbb() error {
    return errors.WithMessage(aaa(), "bbb failed")
 }
 
 func main() {
    err := bbb()
    if errors.Cause(err) == sql.ErrNoRows {
       fmt.Printf("data not found, %v\n", err)
       fmt.Printf("%+v\n", err)
       return
    }
    if err != nil {
       // unknown error
    }
 }
/* Outputs:
data not found, bbb failed: aaa failed: sql: no rows in result set
bbb failed:
    main.bbb
        /usr/four/main.go:12
  - aaa failed:
    main.aaa
        /usr/four/main.go:18
  - sql: no rows in result set
*/




 /*
Go 1.13 内置支持

 */

 import (
   "database/sql"
   "errors"
   "fmt"
)

func bbb() error {
   if err := aaa(); err != nil {
      return fmt.Errorf("bbb failed: %w", aaa())
   }
   return nil
}

func aaa() error {
   return fmt.Errorf("aaa failed: %w", sql.ErrNoRows)
}

func main() {
   err := bbb()
   if errors.Is(err, sql.ErrNoRows) {
      fmt.Printf("data not found,  %+v\n", err)
      return
   }
   if err != nil {
      // unknown error
   }
}
/* Outputs:
data not found,  bbb failed: aaa failed: sql: no rows in result set
*/


