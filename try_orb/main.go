package main

import (
	"fmt"
	"strconv"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/maptile"
)

func generateQuadkeyStr(t maptile.Tile) string {
	s := strconv.FormatInt(int64(t.Quadkey()), 4)

	zeros := "000000000000000000000000000000"
	return zeros[:((int(t.Z)+1)-len(s))/2] + s
}

func main() {
	lat := 35.685323
	lng := 139.752768
	point := orb.Point{lng, lat}

	mapTile := maptile.At(point, 17)
	fmt.Println(generateQuadkeyStr(mapTile))
}