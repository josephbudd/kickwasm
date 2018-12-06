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
		"PriceList": map[string]map[string][]string{
			"": map[string][]string{
				"PriceListsButton": []string{
					"PriceListsButton",
				},
			},
			"PriceListPanel": map[string][]string{
				"ImportButton": []string{
					"PriceListsButton", "PriceListPanel", "ImportButton",
				},
				"EditButton": []string{
					"PriceListsButton", "PriceListPanel", "EditButton",
				},
				"ViewButton": []string{
					"PriceListsButton", "PriceListPanel", "ViewButton",
				},
			},
		},
		"Customer": map[string]map[string][]string{
			"": map[string][]string{
				"CustomersButton": []string{
					"CustomersButton",
				},
			},
			"CustomerPanel": map[string][]string{
				"AddButton": []string{
					"CustomersButton", "CustomerPanel", "AddButton",
				},
				"EditButton": []string{
					"CustomersButton", "CustomerPanel", "EditButton",
				},
			},
		},
		"PriceDrop": map[string]map[string][]string{
			"": map[string][]string{
				"PriceDropsButton": []string{
					"PriceDropsButton",
				},
			},
			"PriceDropPanel": map[string][]string{
				"ViewButton": []string{
					"PriceDropsButton", "PriceDropPanel", "ViewButton",
				},
				"AddButton": []string{
					"PriceDropsButton", "PriceDropPanel", "AddButton",
				},
				"VoidUnvoidButton": []string{
					"PriceDropsButton", "PriceDropPanel", "VoidUnvoidButton",
				},
			},
		},
		"PurchaseOrder": map[string]map[string][]string{
			"": map[string][]string{
				"PurchaseOrdersButton": []string{
					"PurchaseOrdersButton",
				},
			},
			"PurchaseOrderButtonPadPanel": map[string][]string{
				"AddButton": []string{
					"PurchaseOrdersButton", "PurchaseOrderButtonPadPanel", "AddButton",
				},
				"VoidUnvoidButton": []string{
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
