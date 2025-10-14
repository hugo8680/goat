package id

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

func Gen(key int64) (int64, error) {
	node, err := snowflake.NewNode(key)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}
	return node.Generate().Int64(), nil
}
