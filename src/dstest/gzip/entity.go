package gzip

import (
	"dstest/util"

	"google.golang.org/appengine/datastore"
)

type Gzip struct {
	Value     string `json:"value" datastore:"-"`
	GzipValue []byte `json:"-" datastore:"Value"`
}

func (g *Gzip) Load(p []datastore.Property) error {
	datastore.LoadStruct(g, p)
	for _, s := range p {
		if s.Name == "Value" {
			g.loadGzipValue(s)
		}
	}
	return nil
}

func (g *Gzip) loadGzipValue(p datastore.Property) error {
	switch v := p.Value.(type) {
	case string:
		g.Value = v
	case datastore.ByteString:
		g.Value, _ = util.UnGzipBytes(v)
	case []byte:
		g.Value, _ = util.UnGzipBytes(v)
	}
	return nil
}

func (g *Gzip) Save() (p []datastore.Property, err error) {
	if len(g.Value) <= 0 {
		return
	}

	gval := util.GzipString(g.Value)
	g.GzipValue = gval
	p, err = datastore.SaveStruct(g)
	return
}
