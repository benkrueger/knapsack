package util
import (
	"fmt"
	"../common"
	"strconv"
	"github.com/mmcdole/gofeed"
	"github.com/mmcdole/gofeed/rss"
)

var records chan *common.TorrentRecord


type FeedTranslator struct {
    defaultTranslator *gofeed.DefaultRSSTranslator
}

func NewFeedTranslator() *FeedTranslator {
  t := &FeedTranslator{}
  
  // We create a DefaultRSSTranslator internally so we can wrap its Translate
  // call since we only want to modify the precedence for a single field.
  t.defaultTranslator = &gofeed.DefaultRSSTranslator{}
  return t
}
func (ct* FeedTranslator) Translate(feed interface{}) (*gofeed.Feed, error) {
	rss, found := feed.(*rss.Feed)

	if !found {
		return nil, fmt.Errorf("Feed did not match expected type of *rss.Feed")
	}

  f, err := ct.defaultTranslator.Translate(rss)
  if err != nil {
    return nil, err
  }

  for _,i := range f.Items {

	t := new(common.TorrentRecord)

	t.Title = i.Title
	t.Link = i.Link
	for _,e := range i.Enclosures {
		tl,_ := strconv.Atoi(e.Length)
		t.Length = tl
	}
	for _,a := range i.Extensions["torznab"]["attr"]{
		if a.Attrs["name"] == "seeders"{
			//get seeds
			ts,_ := strconv.Atoi(a.Attrs["value"])
			t.Seeders = ts
		}
		if a.Attrs["name"] == "peers"{
			//get peers
			tp,_ := strconv.Atoi(a.Attrs["value"])
			t.Peers = tp 
		}
	}
	if(t.Peers - t.Seeders > 0){
		records <-t
	}

	}

	return f,err
}

func ParseRSSFeed(url string,ch chan<- *common.TorrentRecord){
	fp := gofeed.NewParser()
	fp.RSSTranslator = NewFeedTranslator()
	go fp.ParseURL(url)

}

func ParseTZFeed(url string, apikey string)(error){
	fp := gofeed.NewParser()
	fp.RSSTranslator = NewFeedTranslator()
	_ , err := fp.ParseURL(fmt.Sprintf("%s?apikey=%s",url,apikey))
	return err
}