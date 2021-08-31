package apihttp

// import (
// 	"fmt"

// 	"github.com/zachtaylor/7elements"
// 	"ztaylor.me/http"
// 	"ztaylor.me/js"
// )

// func PacksHandler(r *http.Quest) error {
// 	packs, err := vii.PackService.GetAll()
// 	if err != nil {
// 		return err
// 	}

// 	j := js.Object{}
// 	for packid, pack := range packs {
// 		j[fmt.Sprintf("%d", packid)] = pack.Json()
// 	}
// 	r.WriteJson(j)
// 	return nil
// }
