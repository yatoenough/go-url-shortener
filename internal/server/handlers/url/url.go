package url

import "database/sql"

type UrlsHandler struct {
	db *sql.DB
}

func New(db *sql.DB) *UrlsHandler {
	return &UrlsHandler{
		db: db,
	}
}

func (uh *UrlsHandler) Shorten(){

}

func (uh *UrlsHandler) Redirect(){
	
}