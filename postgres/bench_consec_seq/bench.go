// Benchmarks gap-free consecutive number generation in Postgres.
// Written to win an argument. Run setup.sql first.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx"
)

// Configuration for the PG connection.
type pgConfig struct {
	host     string
	port     int
	database string
	user     string
	password string
}

func (c *pgConfig) dsn() string {
	return fmt.Sprintf("host=%s port=%d database=%s user=%s password=%s",
		c.host, c.port, c.database, c.user, c.password)
}

// Returns an initial PostgreSQL configuration. It uses the same
// defaults (including environment variable names) as psql.
func initPGConfig() pgConfig {
	user := os.Getenv("USER")
	var env = map[string]string{
		"PGHOST":     "localhost",
		"PGPORT":     "5432",
		"PGDATABASE": user,
		"PGUSER":     user,
		"PGPASSWORD": "",
	}
	for name := range env {
		if v := os.Getenv(name); v != "" {
			env[name] = v
		}
	}
	port, _ := strconv.Atoi(env["PGPORT"]) // port=0 on error
	return pgConfig{
		host:     env["PGHOST"],
		port:     port,
		database: env["PGDATABASE"],
		user:     env["PGUSER"],
		password: env["PGPASSWORD"],
	}
}

func main() {
	// Read the environment and handle command line flags.
	var (
		pgConf      = initPGConfig()
		verbose     bool
		trace       bool
		numCounters int
		numWorkers  int
	)
	flag.StringVar(&pgConf.host, "host", pgConf.host, "server host name")
	flag.IntVar(&pgConf.port, "port", pgConf.port, "server listen port")
	flag.StringVar(&pgConf.database, "db", pgConf.database, "database name")
	flag.StringVar(&pgConf.user, "user", pgConf.user, "database user")
	flag.StringVar(&pgConf.password, "password", pgConf.password, "password for auth")
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.BoolVar(&trace, "trace", false, "enable trace logging (very verbose)")
	flag.IntVar(&numCounters, "c", 1, "number of counters")
	flag.IntVar(&numWorkers, "n", 1, "number of workers per counter")
	flag.Parse()

	// Connect to postgresql.
	connConf, err := pgx.ParseDSN(pgConf.dsn())
	must(err)
	poolConf := pgx.ConnPoolConfig{
		ConnConfig:     connConf,
		MaxConnections: numCounters * numWorkers, // No more, no less.
	}
	pool, err := pgx.NewConnPool(poolConf)
	must(err)

	// First initialize the counters.
	conn, err := pool.Acquire()
	must(err)
	for i := 0; i < numCounters; i++ {
		name := nthCounterName(i)
		_, err := conn.Exec("insert into gapless_sequence (name, last) values ($1, 0)", name)
		if err, ok := err.(pgx.PgError); ok && err.Code == "23505" {
			continue // Expected and ok.
		}
		must(err)
	}
	pool.Release(conn)

	log.Printf("progress will be shown every %s", logProgressInterval)

	// Launch workers.
	for i := 0; i < numCounters; i++ {
		name := nthCounterName(i)
		for j := 0; j < numWorkers; j++ {
			conn, err := pool.Acquire()
			must(err)
			go worker(name, j, conn)
		}
	}

	// This program runs until killed. Doesn't do graceful signal handling.
	for {
		time.Sleep(time.Minute)
	}
}

// Each counter should have a unique name.
func nthCounterName(n int) string { return fmt.Sprintf("C%02d", n) }

// How frequently workers will report progress.
const logProgressInterval = 5 * time.Second

// Loops forever, inserting into my_thing as quickly as possible.
func worker(name string, workerID int, c *pgx.Conn) {
	const insertSQL = `
insert into my_thing (seq_name, junk1, junk2, junk3, junk4)
values ($1, $2, $3, $4, $5)
`
	t := time.Now()
	numInserts := 0
	for {
		_, err := c.Exec(insertSQL, name, junk(), junk(), junk(), junk())
		must(err)
		numInserts++
		// Occasionally print our insert rate to the log.
		now := time.Now()
		if since := now.Sub(t); since >= logProgressInterval {
			s := fseconds(since)
			log.Printf("worker %s-%02d: rate over last %.1fs was %.1f/s",
				name, workerID, s, float64(numInserts)/s)
			t = now
			numInserts = 0
		}
	}
}

// Converts a duration to floating-point seconds.
func fseconds(d time.Duration) float64 { return float64(d) / float64(time.Second) }

// Returns 100 bytes of ASCII junk.
func junk() string {
	const n = 100
	a := make([]byte, n)
	for i := 0; i < n; i++ {
		a[i] = byte(32 + rand.Intn(127-32))
	}
	return string(a)
}

// We don't care about gracefully error handling. Just crash the program.
func must(err error) {
	if err != nil {
		// Use the %+v verb in case err supports it.
		log.Fatalf("%+v", err)
	}
}
