package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func getFileNameWithoutExt(path string) string {
	// Fixed with a nice method given by mattn-san
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}

func incId(id string) string {
	// byte --> int
	idIntList := []int{}
	for i := 0; i < len(id); i++ {
		idInt, _ := strconv.Atoi(string(id[i]))
		idIntList = append(idIntList, idInt)
	}

	var carry bool
	idIntList[len(idIntList)-1]++
	for i := len(idIntList) - 1; i >= 0; i-- {
		if carry {
			idIntList[i]++
		}
		if idIntList[i] > 9 {
			idIntList[i] -= 10
			carry = true
		}
	}

	// int --> string
	var ret string
	for _, idInt := range idIntList {
		ret += strconv.Itoa(idInt)
	}
	return ret
}

/**
 * fileName: インクリメントしたいベースファイルの名前(foo.bar)
 * dir     : 検索するディレクトリ
 * n       : 番号の桁数
 */
func GetNewFilePath(baseFileName string, dir string, n int) string {
	ex := filepath.Ext(baseFileName)
	baseFileNameWithoutExt := baseFileName[:len(baseFileName)-len(ex)]
	base := filepath.Join(dir, baseFileNameWithoutExt)
	files, _ := filepath.Glob(base + "*")

	latestId := strings.Repeat("0", n-1) + "1" // 初期値
	if len(files) != 0 {
		for _, f := range files {
			id := f[len(f)-4-len(latestId) : len(f)-4]
			if latestId < id {
				latestId = id
			}
		}
		latestId = incId(latestId)
	}
	return base + latestId + ex
}
