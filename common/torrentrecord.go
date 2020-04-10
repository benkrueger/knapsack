package common

type TorrentRecord struct {
	Length int
	Peers int
	Seeders int
	Leeches int
	Link string
	Title string
}