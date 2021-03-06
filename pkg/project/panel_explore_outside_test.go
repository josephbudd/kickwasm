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
		"CustomersButton": map[string]map[string][]string{
			"": map[string][]string{
				"CustomersButton": []string{"CustomersButton"},
			},
			"CustomerPanel": map[string][]string{
				"CustomerAddButton":  []string{"CustomersButton", "CustomerPanel", "CustomerAddButton"},
				"CustomerEditButton": []string{"CustomersButton", "CustomerPanel", "CustomerEditButton"},
			},
		},
		"PriceDropsButton": map[string]map[string][]string{
			"": map[string][]string{
				"PriceDropsButton": []string{"PriceDropsButton"},
			},
			"PriceDropPanel": map[string][]string{
				"PriceDropAddButton":        []string{"PriceDropsButton", "PriceDropPanel", "PriceDropAddButton"},
				"PriceDropViewButton":       []string{"PriceDropsButton", "PriceDropPanel", "PriceDropViewButton"},
				"PriceDropVoidUnvoidButton": []string{"PriceDropsButton", "PriceDropPanel", "PriceDropVoidUnvoidButton"},
			},
		},
		"PriceListsButton": map[string]map[string][]string{
			"": map[string][]string{
				"PriceListsButton": []string{"PriceListsButton"},
			},
			"PriceListPanel": map[string][]string{
				"PriceListEditButton":   []string{"PriceListsButton", "PriceListPanel", "PriceListEditButton"},
				"PriceListImportButton": []string{"PriceListsButton", "PriceListPanel", "PriceListImportButton"},
				"PriceListViewButton":   []string{"PriceListsButton", "PriceListPanel", "PriceListViewButton"},
			},
		},
		"PurchaseOrdersButton": map[string]map[string][]string{
			"": map[string][]string{
				"PurchaseOrdersButton": []string{"PurchaseOrdersButton"},
			},
			"PurchaseOrderButtonPadPanel": map[string][]string{
				"PurchaseOrderAddButton":        []string{"PurchaseOrdersButton", "PurchaseOrderButtonPadPanel", "PurchaseOrderAddButton"},
				"PurchaseOrderVoidUnvoidButton": []string{"PurchaseOrdersButton", "PurchaseOrderButtonPadPanel", "PurchaseOrderVoidUnvoidButton"},
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
	testBuilder_GenerateHomePanelButtonFolderPathMap(t, builder, want)
}

func testBuilder_GenerateHomePanelButtonFolderPathMap(t *testing.T, builder *project.Builder, want map[string]map[string]map[string][]string) {
	if got := builder.GenerateHomePanelButtonFolderPathMap(); !reflect.DeepEqual(got, want) {
		t.Errorf("Builder.GenerateHomePanelButtonFolderPathMap() = %#v, want %v", got, want)
	}
}
