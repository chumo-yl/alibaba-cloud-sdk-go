package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AlarmInListActiveAlertRule is a nested struct in cms response
type AlarmInListActiveAlertRule struct {
	Uuid               string `json:"Uuid" xml:"Uuid"`
	Name               string `json:"Name" xml:"Name"`
	Namespace          string `json:"Namespace" xml:"Namespace"`
	MetricName         string `json:"MetricName" xml:"MetricName"`
	Period             string `json:"Period" xml:"Period"`
	Statistics         string `json:"Statistics" xml:"Statistics"`
	ComparisonOperator string `json:"ComparisonOperator" xml:"ComparisonOperator"`
	Threshold          string `json:"Threshold" xml:"Threshold"`
	EvaluationCount    string `json:"EvaluationCount" xml:"EvaluationCount"`
	StartTime          string `json:"StartTime" xml:"StartTime"`
	EndTime            string `json:"EndTime" xml:"EndTime"`
	SilenceTime        string `json:"SilenceTime" xml:"SilenceTime"`
	NotifyType         string `json:"NotifyType" xml:"NotifyType"`
	Enable             string `json:"Enable" xml:"Enable"`
	State              string `json:"State" xml:"State"`
	ContactGroups      string `json:"ContactGroups" xml:"ContactGroups"`
	Webhook            string `json:"Webhook" xml:"Webhook"`
	RuleName           string `json:"RuleName" xml:"RuleName"`
}
