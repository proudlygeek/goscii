//
// MongoDB manager implementation of
// proudlygeek/goscii/art/manager.ArtManager
//
package manager

import (
	"labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

func (this *MongoWriter) Write(p []byte) (i int, err error) {
    for _, b := range p {
        this.Buffer = append(this.Buffer, b)
    }

    return len(this.Buffer), err
}

func (this *MongoArtManager) Load(uri string) []byte {

    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    c := session.DB("test").C("urls")

    result := &Art{}

    err = c.Find(bson.M{"_id": this.Encoder.DecodeURI(uri)}).One(&result)
    if err != nil {
        panic(err)
    }

    return result.Content
}

func (this *MongoArtManager) Save(writer *MongoWriter) string {

    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    c := session.DB("test").C("counters")

    // If counter document doesn't exist create it
    if n, err := c.Count(); n == 0 {
        err = c.Insert(bson.M{ "_id": "urlsId", "c": 0})
        if err != nil {
            panic(err)
        }
    }

    change := mgo.Change{
        Update: bson.M{"$inc": bson.M{"c": 1}},
        ReturnNew: true,
    }
    
    doc := &Doc{}

    _, err = c.Find(bson.M{"_id": "urlsId"}).Apply(change, &doc)
    if err != nil {
        panic(err)
    }

    c = session.DB("test").C("urls")
    err = c.Insert(bson.M{"_id": doc.C, "content": writer.Buffer})
    if err != nil {
        panic(err)
    }

    // fmt.Printf("PIC URL IS %s (ID is %d)\n", encoder.EncodeURI(doc.C), encoder.DecodeURI(encoder.EncodeURI(doc.C)))
    return this.Encoder.EncodeURI(doc.C)   
}
