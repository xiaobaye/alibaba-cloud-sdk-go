package das

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

// EnableAutoThrottleList is a nested struct in das response
type EnableAutoThrottleList struct {
	ActiveSessions         int64  `json:"ActiveSessions" xml:"ActiveSessions"`
	AutoKillSession        bool   `json:"AutoKillSession" xml:"AutoKillSession"`
	Visible                bool   `json:"Visible" xml:"Visible"`
	CpuUsage               string `json:"CpuUsage" xml:"CpuUsage"`
	InstanceId             string `json:"InstanceId" xml:"InstanceId"`
	UserId                 string `json:"UserId" xml:"UserId"`
	AllowThrottleStartTime string `json:"AllowThrottleStartTime" xml:"AllowThrottleStartTime"`
	AllowThrottleEndTime   string `json:"AllowThrottleEndTime" xml:"AllowThrottleEndTime"`
	CpuSessionRelation     string `json:"CpuSessionRelation" xml:"CpuSessionRelation"`
	AbnormalDuration       string `json:"AbnormalDuration" xml:"AbnormalDuration"`
	MaxThrottleTime        string `json:"MaxThrottleTime" xml:"MaxThrottleTime"`
}
