// Create a new dashboard with a change widget using formulas and functions slo query

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	// there is a valid "slo" in the system
	SloData0ID := os.Getenv("SLO_DATA_0_ID")

	body := datadogV1.Dashboard{
		Title: "Example-Dashboard",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					ChangeWidgetDefinition: &datadogV1.ChangeWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Time:       &datadogV1.WidgetTime{},
						Type:       datadogV1.CHANGEWIDGETDEFINITIONTYPE_CHANGE,
						Requests: []datadogV1.ChangeWidgetRequest{
							{
								Formulas: []datadogV1.WidgetFormula{
									{
										Formula: "hour_before(query1)",
									},
									{
										Formula: "query1",
									},
								},
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionSLOQueryDefinition: &datadogV1.FormulaAndFunctionSLOQueryDefinition{
											Name:                   datadog.PtrString("query1"),
											DataSource:             datadogV1.FORMULAANDFUNCTIONSLODATASOURCE_SLO,
											SloId:                  SloData0ID,
											Measure:                datadogV1.FORMULAANDFUNCTIONSLOMEASURE_SLO_STATUS,
											GroupMode:              datadogV1.FORMULAANDFUNCTIONSLOGROUPMODE_OVERALL.Ptr(),
											SloQueryType:           datadogV1.FORMULAANDFUNCTIONSLOQUERYTYPE_METRIC.Ptr(),
											AdditionalQueryFilters: datadog.PtrString("*"),
										}},
								},
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
								OrderBy:        datadogV1.WIDGETORDERBY_CHANGE.Ptr(),
								ChangeType:     datadogV1.WIDGETCHANGETYPE_ABSOLUTE.Ptr(),
								IncreaseGood:   datadog.PtrBool(true),
								OrderDir:       datadogV1.WIDGETSORT_ASCENDING.Ptr(),
							},
						},
					}},
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  4,
					Height: 2,
				},
			},
		},
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
