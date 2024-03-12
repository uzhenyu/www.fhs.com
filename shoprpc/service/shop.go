package service

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/httplib"
	"gorm.io/gorm"
	"strconv"
	"zg5/z311/message/user"
	"zg5/z311/shoprpc/model"
)

func Create(info *user.ShopInfo) (*user.ShopInfo, error) {
	newShop := model.NewShop()
	res, err := newShop.Create(pbToMysql(info))
	if err != nil {
		return nil, err
	}
	return mysqlToPb(res)
}

func Update(Id, num int64) error {
	newShop := model.NewShop()
	err := newShop.Update(Id, num)
	if err != nil {
		return err
	}
	return nil
}

func Select() (infos []*user.ShopInfo, err error) {
	newShop := model.NewShop()
	res, err := newShop.Select()
	if err != nil {
		return nil, err
	}
	for _, v := range res {
		info, _ := mysqlToPb(v)
		infos = append(infos, info)
	}
	return infos, nil
}

func Delete(id int64) error {
	newShop := model.NewShop()
	err := newShop.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func pbToMysql(info *user.ShopInfo) *model.Shop {
	return &model.Shop{
		Model: gorm.Model{
			ID: uint(info.ID),
		},
		Name:  info.Name,
		Price: float64(info.Price),
		Num:   info.Num,
	}
}

func mysqlToPb(shop *model.Shop) (*user.ShopInfo, error) {
	return &user.ShopInfo{
		ID:    int64(shop.ID),
		Name:  shop.Name,
		Price: float32(shop.Price),
		Num:   shop.Num,
	}, nil
}

type SearchT struct {
	Query struct {
		Match struct {
			Name string `json:"name"`
		} `json:"match"`
	} `json:"query"`
}

type SearchTT struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			Id     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Id    int     `json:"id"`
				Name  string  `json:"name"`
				Price float64 `json:"price"`
				Num   int     `json:"num"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func Search(name string) (SearchTT, error) {
	res := httplib.Post("http://127.0.0.1:9201/shop/_search")
	var s SearchT
	s.Query.Match.Name = name
	res.JSONBody(s)
	r, err := res.String()
	if err != nil {
		return SearchTT{}, err
	}
	var ss SearchTT
	json.Unmarshal([]byte(r), &ss)
	return ss, nil
}

type AddT struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Num   int     `json:"num"`
}

func Add(id, num int64, price float64, name string) error {
	res := httplib.Post("http://127.0.0.1:9201/shop/_doc/" + strconv.Itoa(int(id)))
	var a AddT
	a.Id = int(id)
	a.Name = name
	a.Num = int(num)
	a.Price = price
	res.JSONBody(a)
	_, err := res.String()
	if err != nil {
		return err
	}
	return nil
}
