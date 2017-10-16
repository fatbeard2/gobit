package main

import (
	"fmt"
	"path/filepath"
	"github.com/fatbeard2/gobit/torrentfile"
)

func main() {
	path, _ := filepath.Abs("./torrentfile/fixtures/book2.torrent")
	meta := torrentfile.TorrentMetaInfo {}
	torrentfile.GetTorrentMetaInfo(&meta, path)
	fmt.Println(meta.Info.Name)
	fmt.Println(meta.Info.PieceLength)
	fmt.Println(meta.Info.Files[0])
	fmt.Println(meta.Info.Files[1])
	fmt.Println(meta.Announce)
	fmt.Println(meta.AnnounceList)
}
