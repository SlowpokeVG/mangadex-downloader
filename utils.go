package main

import (
	"math"
	"os"

	"golang.org/x/net/html"
)

func IsFloat(n float64) bool {
	return math.Floor(n) == n
}

func GetAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

func DirExists(dir string) error {
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}
