package harvest

import (
	"fmt"
	"time"
)

type ProjectResponse struct {
	Project *Project `json:"project"`
}
type ProjectAnalysisResponse struct {
	Tasks []ProjectAnalysisTask `json:"tasks"`
}

type ProjectAnalysisTask struct {
	ID                    int64   `json:"id"`
	Name                  string  `json:"name"`
	IsBillable            bool    `json:"billable"`
	IsArchived            bool    `json:"archived"`
	TotalHours            float64 `json:"total_hours"`
	BillableHours         float64 `json:"billable_hours"`
	BilledRate            float64 `json:"billed_rate"`
	BillableAmount        float64 `json:"billable_amount"`
	InternalCost          float64 `json:"internal_cost"`
	DetailedReportUrl     string  `json:"detailed_report_url"`
	Budget                float64 `json:"budget"`
	BudgetSpent           float64 `json:"budget_spent"`
	BudgetLeft            float64 `json:"budget_left"`
	IsOverBudget          bool    `json:"over_budget"`
	OverBudgetPercentage  float64 `json:"over_budget_percentage"`
	BudgetSpentPercentage float64 `json:"budget_spent_percentage"`
}

type Project struct {
	ID                               int64     `json:"id"`
	ClientID                         int64     `json:"client_id"`
	Name                             string    `json:"name"`
	Code                             string    `json:"code"`
	Active                           bool      `json:"active"`
	Billable                         bool      `json:"billable"`
	BillBy                           string    `json:"bill_by"`
	HourlyRate                       *float64  `json:"hourly_rate"`
	BudgetBy                         string    `json:"budget_by"`
	Budget                           *float64  `json:"budget"`
	NotifyWhenOverBudget             bool      `json:"notify_when_over_budget"`
	OverBudgetNotificationPercentage float64   `json:"over_budget_notification_percentage"`
	OverBudgetNotifiedAt             *Date     `json:"over_budget_notified_at"`
	ShowBudgetToAll                  bool      `json:"show_budget_to_all"`
	CreatedAt                        time.Time `json:"created_at"`
	UpdatedAt                        time.Time `json:"updated_at"`
	StartsOn                         Date      `json:"starts_on"`
	EndsOn                           Date      `json:"ends_on"`
	Estimate                         *float64  `json:"estimate"`
	EstimateBy                       string    `json:"estimate_by"`
	HintEarliestRecordAt             Date      `json:"hint_earliest_record_at"`
	HintLatestRecordAt               Date      `json:"hint_latest_record_at"`
	Notes                            string    `json:"notes"`
	CostBudget                       *float64  `json:"cost_budget"`
	CostBudgetIncludeExpenses        bool      `json:"cost_budget_include_expenses"`
}

func (a *API) GetProject(projectID int64, args Arguments) (project *Project, err error) {
	projectResponse := ProjectResponse{}
	path := fmt.Sprintf("/projects/%v", projectID)
	err = a.Get(path, args, &projectResponse)
	return projectResponse.Project, err
}

func (a *API) GetProjectAnalysis(projectID int64, args Arguments) ([]ProjectAnalysisTask, error) {
	projectAnalysisResponse := ProjectAnalysisResponse{}
	path := fmt.Sprintf("/projects/%v/analysis.json", projectID)
	err := a.Get(path, args, &projectAnalysisResponse)
	return projectAnalysisResponse.Tasks, err
}

func (a *API) GetProjects(args Arguments) (projects []*Project, err error) {
	projects = make([]*Project, 0)
	projectsResponse := make([]*ProjectResponse, 0)
	path := fmt.Sprintf("/projects")
	err = a.Get(path, args, &projectsResponse)
	for _, pr := range projectsResponse {
		projects = append(projects, pr.Project)
	}
	return projects, err
}
