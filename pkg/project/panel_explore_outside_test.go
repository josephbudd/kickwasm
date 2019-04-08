package project_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/josephbudd/kickwasm/pkg/project"
	"github.com/josephbudd/kickwasm/pkg/slurp"
)

func TestOutsideExplore(t *testing.T) {
	var want = map[string]map[string]map[string][]string{
		"PriceList": {
			"": {
				"PriceListsButton": {
					"PriceListsButton",
				},
			},
			"PriceListPanel": {
				"ImportButton": {
					"PriceListsButton", "PriceListPanel", "ImportButton",
				},
				"EditButton": {
					"PriceListsButton", "PriceListPanel", "EditButton",
				},
				"ViewButton": {
					"PriceListsButton", "PriceListPanel", "ViewButton",
				},
			},
		},
		"Customer": {
			"": {
				"CustomersButton": {
					"CustomersButton",
				},
			},
			"CustomerPanel": {
				"AddButton": {
					"CustomersButton", "CustomerPanel", "AddButton",
				},
				"EditButton": {
					"CustomersButton", "CustomerPanel", "EditButton",
				},
			},
		},
		"PriceDrop": {
			"": {
				"PriceDropsButton": {
					"PriceDropsButton",
				},
			},
			"PriceDropPanel": {
				"ViewButton": {
					"PriceDropsButton", "PriceDropPanel", "ViewButton",
				},
				"AddButton": {
					"PriceDropsButton", "PriceDropPanel", "AddButton",
				},
				"VoidUnvoidButton": {
					"PriceDropsButton", "PriceDropPanel", "VoidUnvoidButton",
				},
			},
		},
		"PurchaseOrder": {
			"": {
				"PurchaseOrdersButton": {
					"PurchaseOrdersButton",
				},
			},
			"PurchaseOrderButtonPadPanel": {
				"AddButton": {
					"PurchaseOrdersButton", "PurchaseOrderButtonPadPanel", "AddButton",
				},
				"VoidUnvoidButton": {
					"PurchaseOrdersButton", "PurchaseOrderButtonPadPanel", "VoidUnvoidButton",
				},
			},
		},
	}
	// pdm tests
	pwd, _ := os.Getwd()
	pdmPath := filepath.Join(pwd, "testyaml", "pdm", "app.yaml")
	sl := slurp.NewSlurper()
	builder, err := sl.Gulp(pdmPath)
	if err != nil {
		t.Fatal(err)
	}
	testBuilder_GenerateServicePanelButtonFolderPathMap(t, builder, want)
}

func testBuilder_GenerateServicePanelButtonFolderPathMap(t *testing.T, builder *project.Builder, want map[string]map[string]map[string][]string) {
	if got := builder.GenerateServicePanelButtonFolderPathMap(); !reflect.DeepEqual(got, want) {
		t.Errorf("Builder.GenerateServicePanelButtonFolderPathMap() = %#v, want %v", got, want)
	}
}
