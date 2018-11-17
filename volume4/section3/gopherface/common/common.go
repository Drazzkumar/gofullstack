package common

import (
	"github.com/razzkumar/gofullstack/volume4/section3/gopherface/common/datastore"
	"go.isomorphicgo.org/go/isokit"
)

type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
