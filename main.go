package main

import (
	//"fmt"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/renoinn/ent_atlas/ent"
	"github.com/renoinn/ent_atlas/ent/ogent"
)

func main() {
    // Create ent client.
    //dataSourceName := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True", "sample_user", "sample_password", "localhost:3306", "sample_db")
    client, err := ent.Open(dialect.MySQL, "sample_user:sample_password@tcp(localhost:3306)/sample_db")
    if err != nil {
        log.Fatal(err)
    }
    // Run the migrations.
    //if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
    //    log.Fatal(err)
    //}
    // Start listening.
    srv, err := ogent.NewServer(ogent.NewOgentHandler(client))
    if err != nil {
        log.Fatal(err)
    }
    if err := http.ListenAndServe(":8080", srv); err != nil {
        log.Fatal(err)
    }
}
