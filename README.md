**Dependencies:**
- "database/sql"
- "io"
- "log"
- "net/http"
- "strconv"
- "modernc.org/sqlite"
____
**Build:**
1. Install sqlite and init testDB.db from init-db.sql 
(for linux you may do `sqlyte3 testDB.db < init-db.sql`)
2. Run server.go with `go run server.go`
3. Open tab at http://127.0.0.1:8080