package metadata

import (
    "database/sql"
    "log"
    "sync"

    _ "github.com/mattn/go-sqlite3"
)

type MetadataDB struct {
    db   *sql.DB
    lock sync.Mutex
}

// OpenDB initializes the metadata SQLite DB at path.
func OpenDB(path string) (*MetadataDB, error) {
    db, err := sql.Open("sqlite3", path)
    if err != nil {
        return nil, err
    }

    createStmt := `
    CREATE TABLE IF NOT EXISTS metadata (
        id TEXT PRIMARY KEY,
        name TEXT,
        size INTEGER,
        type TEXT,
        uploader TEXT,
        timestamp TEXT,
        tags TEXT,
        hash TEXT
    );
    `
    if _, err := db.Exec(createStmt); err != nil {
        return nil, err
    }

    return &MetadataDB{db: db}, nil
}

// InsertOrUpdate adds new metadata or updates existing one.
func (m *MetadataDB) InsertOrUpdate(meta Metadata) error {
    m.lock.Lock()
    defer m.lock.Unlock()

    tagsStr := ""
    for _, tag := range meta.Tags {
        tagsStr += tag + ","
    }

    stmt := `
    INSERT OR REPLACE INTO metadata (id, name, size, type, uploader, timestamp, tags, hash)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `
    _, err := m.db.Exec(stmt,
        meta.ID,
        meta.Name,
        meta.Size,
        meta.Type,
        meta.Uploader,
        meta.Timestamp.Format("2006-01-02T15:04:05Z"),
        tagsStr,
        meta.Hash,
    )
    return err
}

// GetAll retrieves all known metadata entries.
func (m *MetadataDB) GetAll() ([]Metadata, error) {
    rows, err := m.db.Query(`SELECT id, name, size, type, uploader, timestamp, tags, hash FROM metadata`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var result []Metadata
    for rows.Next() {
        var meta Metadata
        var timestampStr string
        var tagsStr string

        err := rows.Scan(
            &meta.ID,
            &meta.Name,
            &meta.Size,
            &meta.Type,
            &meta.Uploader,
            &timestampStr,
            &tagsStr,
            &meta.Hash,
        )
        if err != nil {
            log.Println("Scan error:", err)
            continue
        }

        meta.Timestamp, _ = parseTime(timestampStr)
        meta.Tags = splitTags(tagsStr)

        result = append(result, meta)
    }

    return result, nil
}

func parseTime(t string) (resultTime sql.NullTime, _ error) {
    parsed, err := sql.ParseTime(t)
    if err != nil {
        return resultTime, err
    }
    return sql.NullTime{Time: parsed, Valid: true}, nil
}

func splitTags(tags string) []string {
    var result []string
    for _, tag := range split(tags, ",") {
        if tag != "" {
            result = append(result, tag)
        }
    }
    return result
}

func split(s, sep string) []string {
    var result []string
    current := ""
    for _, ch := range s {
        if string(ch) == sep {
            result = append(result, current)
            current = ""
        } else {
            current += string(ch)
        }
    }
    if current != "" {
        result = append(result, current)
    }
    return result
}

