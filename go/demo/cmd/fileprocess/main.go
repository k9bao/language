package main

import (
	"../../fileG"
)

func main() {
	fileG.WalkDir("D:/work/knowledgebao/knowledgebao.github.io/_posts", 0, fileG.AddTagData)
}
