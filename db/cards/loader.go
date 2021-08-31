package cards

import (
	"github.com/zachtaylor/7elements/card"
	"taylz.io/db"
)

type Loader struct{ conn *db.DB }

func NewLoader(conn *db.DB) *Loader { return &Loader{conn} }

func (loader Loader) isLoader() card.Loader { return loader }

func (loader Loader) GetAll() card.Prototypes { return GetAll(loader.conn) }
