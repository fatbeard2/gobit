package torrentfile

import (
	"github.com/jackpal/bencode-go"
	"log"
	"os"
)

func GetTorrentMetaInfo(meta *TorrentMetaInfo, path string) (error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return bencode.Unmarshal(file, meta)
}

type TorrentMetaInfo struct {
	Info singleOrMultipleFilesInfo
	Announce string
	AnnounceList [][]string `bencode:"announce-list"`
	Pieces string
}

type singleOrMultipleFilesInfo struct {
	Name string
	Length int
	Md5sum string
	Files []FileInfo
	Pieces string
	PieceLength int `bencode:"piece length"`
}

type FileInfo struct {
	Name string
	Length int
	Md5sum string
	Path []string
}
