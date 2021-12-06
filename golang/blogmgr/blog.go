package blogmgr

import (
	"database/sql"
	"time"

	"golang.org/x/net/context"
)

type Blog struct {
	Id       int       `json:"id" sql:"id"`
	CateId   int       `json:"cate_id" sql:"cate_id"`
	Kind     int       `json:"kind" sql:"kind"`
	Status   int       `json:"status" sql:"status"`
	Title    string    `json:"title" sql:"title"`
	Path     string    `json:"path" sql:"path"`
	Summary  string    `json:"summary" sql:"summary"`
	Markdown string    `json:"markdown" sql:"markdown"`
	RichText string    `json:"richtext" sql:"richtext"`
	Allow    int       `json:"allow" sql:"allow"`
	Created  time.Time `json:"created" sql:"created"`
	Updated  time.Time `json:"updated" sql:"updated"`
	Creator  int       `json:"creator" sql:"creator"`
	Updater  int       `json:"updater" sql:"updater"`
}

func (b *Blog) insert(conn *sql.Conn, ctx context.Context) error {
	var strInsertSql = `insert into post (cate_id,kind,status,title,path,summary,markdown,richtext,allow,created,updated,creator,updater) 
	values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	params := []interface{}{b.CateId, b.Kind, b.Status, b.Title}
	_, err := conn.ExecContext(ctx, strInsertSql, params...)
	if err != nil {
		logger.ErrErr("ExecContext Failed"+strInsertSql, err)
		return err
	}
	return nil
}
