package tile_rules

import (
	"path"
	"strconv"
	"strings"
)

// IDBRules IDB
type IDBRules struct {
	Path    string
	Crs     string
	Z, X, Y int64
}

func (i *IDBRules) Rules() string {
	key := ""
	switch i.Crs {
	case "3857":
		key = i.In3857()
	case "4326":
		key = i.In4326()
	default:
		key = "base"
	}
	length := len(key)
	k := length - 8
	dirLen := 5
	filename := strconv.Itoa(length) + "/"
	for i := dirLen; i < k; i = i + dirLen {
		filename = filename + key[i-dirLen:i] + "/"
	}
	filename = filename + key[0:k]
	if !strings.HasSuffix(i.Path, "/") {
		i.Path = i.Path + "/"
	}

	storagePath := "db://" + path.Join(i.Path, filename, key)
	return storagePath

}

func (i *IDBRules) In3857() string {
	z := i.Z
	y := i.Y
	x := i.X
	z = z - 1
	tilekey := ""
	if z != 0 {
		y := (1<<(z-1))*3 - y - 1
		for a := 0; a <= int(z); a++ {
			x0 := x % 2
			y0 := y % 2
			if x0 == 0 {
				if y0 == 0 {
					tilekey = "0" + tilekey
				} else {
					tilekey = "3" + tilekey
				}
			} else {
				if y0 == 0 {
					tilekey = "1" + tilekey
				} else {
					tilekey = "2" + tilekey
				}
			}
			x = x >> 1
			y = y >> 1
		}
		return "0" + tilekey
	} else {
		return "0"
	}
}

func (i *IDBRules) In4326() string {
	z := i.Z
	y := i.Y
	x := i.X

	tilekey := ""
	if z == 0 {
		return "0"
	}

	y = (1<<(z-1))*3 - y - 1
	for a := 0; a <= int(z); a++ {
		x0 := x % 2
		y0 := y % 2
		if x0 == 0 {
			if y0 == 0 {
				tilekey = "0" + tilekey
			} else {
				tilekey = "3" + tilekey
			}
		} else {
			if y0 == 0 {
				tilekey = "1" + tilekey
			} else {
				tilekey = "2" + tilekey
			}
		}
		x = x >> 1
		y = y >> 1
	}
	return "0" + tilekey
}
