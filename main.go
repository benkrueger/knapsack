package main
import (
	"fmt"
	"./util"
	"./common"
)
/*
Knapsack
Dynamic programming download optimizer/manager.
knapsack [-l] <config file> [-d](dry run) 
*/

func main(){
	parsechannel := make(chan *common.TorrentRecord,100)
	go util.ParseRSSFeed("http://127.0.0.1:9117/api/v2.0/indexers/animetorrents/results/torznab/api?apikey=9a368hodza783ve0tuvjpjzrq8txega3&t=search&cat=&q=",parsechannel)
	//go util.ParseRSSFeed("http://127.0.0.1:9117/api/v2.0/indexers/iptorrents/results/torznab/api?apikey=9a368hodza783ve0tuvjpjzrq8txega3&t=search&cat=&q=",torchannel)

	for t := range parsechannel {
		fmt.Println(t.Title)
	}

}
