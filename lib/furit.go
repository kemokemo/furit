package furit

import (
	"io"
	"log"
)

// ログファイルのClose処理など、Fatalな扱いにするほどでないClose処理用。
// Special thanks to https://zenn.dev/sess/articles/3ef56eef7fbf92
func shortClose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("cannot close : %v", err)
	}
}
