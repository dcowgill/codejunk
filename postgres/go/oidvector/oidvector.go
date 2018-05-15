// Experiment with selecting column type OIDVECTOR.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

type oidVector []pgtype.OID

// The binary representation of an oidvector doesn't seem to be documented
// anywhere, so we rely on its text representation instead. (It isn't
// documented, either, but it does seem to be straightforward: a sequence of
// integer strings, separated by whitespace.)
func (v *oidVector) DecodeText(ci *pgtype.ConnInfo, src []byte) error {
	for _, s := range strings.Fields(string(src)) {
		n, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return err
		}
		*v = append(*v, pgtype.OID(n))
	}
	return nil
}

func main() {
	connInfo := flag.String("conninfo", "host=localhost", "Postgres conninfo string or URI")
	flag.Parse()

	connConf, err := pgx.ParseConnectionString(*connInfo)
	if err != nil {
		fatalf("invalid Postgres conninfo string %q: %s", *connInfo, err)
	}

	conn, err := pgx.Connect(connConf)
	if err != nil {
		fatalf("pgx.Connect: %s", err)
	}
	defer conn.Close()

	// Collect the unique OIDs in "pg_index.indclass".
	rows, err := conn.Query(`
SELECT i.indclass
  FROM pg_class c
  JOIN pg_index i ON i.indexrelid = c.oid
`)
	if err != nil {
		fatalf("conn.Query: %s", err)
	}
	uniqOID := make(map[pgtype.OID]bool)
	for rows.Next() {
		var indClass oidVector
		if err := rows.Scan(&indClass); err != nil {
			fatalf("rows.Scan: %s", err)
		}
		fmt.Printf("indclass = %+v\n", indClass)
		for _, oid := range indClass {
			uniqOID[oid] = true
		}
	}
	if err := rows.Err(); err != nil {
		fatalf("rows.Err returned an error: %s", err)
	}

	// Look up, in pg_opclass, the OIDs we gathered above.
	for oid := range uniqOID {
		row := conn.QueryRow(`SELECT opcname FROM pg_opclass WHERE oid = $1`, oid)
		var opcname string
		if err := row.Scan(&opcname); err != nil {
			fatalf("error scanning pg_opclass where oid=%v: %s", oid, err)
		}
		fmt.Printf("oid = %v => opcname = %q\n", oid, opcname)
	}
}

func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
