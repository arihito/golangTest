package fruit

import "log"

type Product struct {
	Name string
}

type Data struct {
	Product []*Product
}

func init() {
	log.Println("include fruit.")
}

func GetList() Data {
	p1 := Product{Name: "りんご"}
	p2 := Product{Name: "なし"}
	p3 := Product{Name: "ばなな"}
	return Data{Product: []*Product{&p1,&p2,&p3}}
}