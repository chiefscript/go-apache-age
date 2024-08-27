package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func NewDB(host, port, user, password, dbname string) (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}

// CreateGraph creates a new graph in the specified database
func (db *DB) CreateGraph(graphName string) error {
	query := fmt.Sprintf("SELECT create_graph('%s');", graphName)
	_, err := db.Conn.Exec(query)
	if err != nil {
		return err
	}
	log.Printf("Graph %s created successfully", graphName)
	return nil
}

// AddVertex adds a new vertex to the specified graph
func (db *DB) AddVertex(graphName, label, properties string) error {
	query := fmt.Sprintf("SELECT * FROM cypher('%s', $$ CREATE (n:%s %s) RETURN n $$) as (n agtype);", graphName, label, properties)
	_, err := db.Conn.Exec(query)
	if err != nil {
		return err
	}
	log.Printf("Vertex with label %s added to graph %s", label, graphName)
	return nil
}

// AddEdge adds a new edge between vertices in the specified graph
func (db *DB) AddEdge(graphName, label, properties, fromVertexId, toVertexId string) error {
	query := fmt.Sprintf("SELECT * FROM cypher('%s', $$ MATCH (a), (b) WHERE id(a) = %s AND id(b) = %s CREATE (a)-[r:%s %s]->(b) RETURN r $$) as (r agtype);", graphName, fromVertexId, toVertexId, label, properties)
	_, err := db.Conn.Exec(query)
	if err != nil {
		return err
	}
	log.Printf("Edge with label %s added between vertices %s and %s in graph %s", label, fromVertexId, toVertexId, graphName)
	return nil
}
