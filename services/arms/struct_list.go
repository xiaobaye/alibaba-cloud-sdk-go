package arms

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

// List is a nested struct in arms response
type List struct {
	TaskId        string  `json:"TaskId" xml:"TaskId"`
	TaskName      string  `json:"TaskName" xml:"TaskName"`
	TaskType      int64   `json:"TaskType" xml:"TaskType"`
	CreateTime    string  `json:"CreateTime" xml:"CreateTime"`
	MonitorNumber int64   `json:"MonitorNumber" xml:"MonitorNumber"`
	Usable        float64 `json:"Usable" xml:"Usable"`
	TaskTypeName  string  `json:"TaskTypeName" xml:"TaskTypeName"`
	TaskStatus    string  `json:"TaskStatus" xml:"TaskStatus"`
	Url           string  `json:"Url" xml:"Url"`
}
