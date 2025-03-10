package product

import (
	"fmt"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
)

var products lib.Map[Product]

type Product struct {
	*model.Product
	//Values map[string]float64
}

func New(model *model.Product) *Product {
	return &Product{
		Product: model,
		//Values: map[string]float64{},
	}
}

func Ensure(id string) (*Product, error) {
	dev := products.Load(id)
	if dev == nil {
		err := Load(id)
		if err != nil {
			return nil, err
		}
		dev = products.Load(id)
	}
	return dev, nil
}

func Get(id string) *Product {
	return products.Load(id)
}

func Load(id string) error {
	var p model.Product
	get, err := db.Engine.ID(id).Get(&p)
	if err != nil {
		return err
	}
	if !get {
		return fmt.Errorf("product %s not found", id)
	}

	return From(&p)
}

func From(model *model.Product) error {
	p := New(model)
	
	products.Store(model.Id, p)

	return nil
}

func LoadAll() error {
	//开机加载所有产品，好像没有必要???
	var ps []*model.Product
	err := db.Engine.Find(&ps)
	if err != nil {
		return err
	}

	for _, p := range ps {
		err = From(p)
		if err != nil {
			log.Error(err)
			//return err
		}
	}

	return nil
}
