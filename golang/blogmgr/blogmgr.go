package blogmgr

import (
	bloglog "blog/log"
	"blog/sqlutils"
	"database/sql"
	"time"

	"golang.org/x/net/context"
)

type BlogMgr struct {
	db   *sql.DB
	conn *sql.Conn
	ctx  context.Context
}

type BlogKey struct {
	Id       int       `json:"id"`
	CateId   int       `json:"cate_id"`
	Kind     int       `json:"kind"`
	Status   int       `json:"status"`
	Title    string    `json:"title"`
	Path     string    `json:"path"`
	Summary  string    `json:"summary"`
	Markdown string    `json:"markdown"`
	RichText string    `json:"richtext"`
	Allow    int       `json:"allow"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Creator  int       `json:"creator"`
	Updater  int       `json:"updater"`
}

func NewBlogMgr(ctx context.Context, logger *bloglog.Logger) (*BlogMgr, error) {
	if !bInit {
		err := InitAPI(logger)
		if err != nil {
			(*logger).ErrErr("initapi failed", err)
			return nil, err
		}
	}
	conn, err := db.Conn(ctx)
	if err != nil {
		(*logger).ErrErr("db.Conn failed", err)
		return nil, err
	}
	return &BlogMgr{
		db:   db,
		conn: conn,
		ctx:  ctx,
	}, nil
}

func (mgr *BlogMgr) Destruct() {
	mgr.conn.Close()
}

type QueryBlogsReturn struct {
	Title   string    `json:"title" sql:"title"`
	Summary string    `json:"summary" sql:"summary"`
	Created time.Time `json:"created" sql:"created"`
}

func (mgr *BlogMgr) QueryBlogs(key *BlogKey) ([]*QueryBlogsReturn, error) {
	var strQuerySql = `select title,summary,created from post where creator=$1 order by created`
	params := []interface{}{(*key).Creator}
	rows, err := mgr.conn.QueryContext(mgr.ctx, strQuerySql, params...)
	if err != nil {
		logger.ErrErr("QueryContext failed", err)
		return nil, err
	}
	defer rows.Close()

	records := sqlutils.GetRow(rows)
	var results []*QueryBlogsReturn
	for i := 0; i < len(records); i++ {
		var r QueryBlogsReturn
		sqlutils.DataToStructBySqlTag(records[i], &r)
		results = append(results, &r)
	}
	return results, nil
}

func (mgr *BlogMgr) AddBlog(key *BlogKey) error {
	var blog = Blog{
		CateId:   key.CateId,
		Kind:     key.Kind,
		Status:   key.Status,
		Title:    key.Title,
		Path:     key.Path,
		Summary:  key.Summary,
		Markdown: key.Markdown,
		RichText: key.RichText,
		Allow:    key.Allow,
		Created:  key.Created,
		Updated:  key.Updated,
		Creator:  key.Creator,
		Updater:  key.Updater,
	}

	tx, err := mgr.conn.BeginTx(mgr.ctx, nil)
	if err != nil {
		logger.ErrErr("BeginTx failed", err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = blog.insert(mgr.conn, mgr.ctx)
	if err != nil {
		logger.ErrErr("blog.insert failed", err)
		return err
	}
	return nil
}

func (mgr *BlogMgr) Delete(id int) error {
	var strDeleteSql = `delete from post where id=$1`
	params := []interface{}{id}

	_, err := mgr.conn.ExecContext(mgr.ctx, strDeleteSql, params...)
	if err != nil {
		logger.ErrErr("ExecContext Failed"+strDeleteSql, err)
		return err
	}
	return nil
}
