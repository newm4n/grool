package examples

import (
	"fmt"
	"github.com/newm4n/grool/builder"
	"github.com/newm4n/grool/context"
	engine2 "github.com/newm4n/grool/engine"
	"github.com/newm4n/grool/model"
	"github.com/newm4n/grool/pkg"
	"time"
)

const (
	ItemTypeLuxury = "LUXURY"
	ItemTypeNormal = "NORMAL"
)

type Purchase struct {
	PurchaseDate time.Time
	ItemType     string
	Price        float64

	Tax             float64
	PriceAfterTax   float64
	IgnoredPurchase bool
}

type CashFlow struct {
	TotalPurchases    float64
	PurchaseCount     int
	TotalTax          float64
	PurchasesAfterTax float64
}

func (cf *CashFlow) String() string {
	return fmt.Sprintf("Purchase count %d total amount %f. Total tax are %f thus total cash in %f", cf.PurchaseCount, cf.TotalPurchases, cf.TotalTax, cf.PurchasesAfterTax)
}

var (
	Purchases = []*Purchase{
		&Purchase{
			PurchaseDate: time.Date(2019, time.January, 4, 13, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.January, 17, 15, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.February, 12, 7, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.February, 24, 3, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 22, 22, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 24, 17, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 15, 14, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 25, 10, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.March, 19, 13, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.June, 6, 21, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
		&Purchase{
			PurchaseDate: time.Date(2019, time.June, 19, 22, 0, 0, 0, time.Local),
			ItemType:     ItemTypeLuxury,
			Price:        100000,
		},
	}
)

type CashFlowCalculator struct {
}

func (cf *CashFlowCalculator) CalculatePurchases() {
	cashFlow := &CashFlow{}

	kb := model.NewKnowledgeBase()
	rb := builder.NewRuleBuilder(kb)
	err := rb.BuildRuleFromResource(pkg.NewFileResource("CashFlowRule.grl"))
	if err != nil {
		panic(err)
	}

	engine := engine2.NewGroolEngine()

	for _, purchase := range Purchases {
		dctx := context.NewDataContext()
		dctx.Add("CashFlow", cashFlow)
		dctx.Add("Purchase", purchase)
		err = engine.Execute(dctx, kb)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(cashFlow.String())
}
