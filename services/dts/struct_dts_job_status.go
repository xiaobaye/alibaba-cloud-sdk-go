package dts

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

// DtsJobStatus is a nested struct in dts response
type DtsJobStatus struct {
	DtsJobDirection               string                          `json:"DtsJobDirection" xml:"DtsJobDirection"`
	DedicatedClusterId            string                          `json:"DedicatedClusterId" xml:"DedicatedClusterId"`
	MemUsage                      string                          `json:"MemUsage" xml:"MemUsage"`
	ErrorMessage                  string                          `json:"ErrorMessage" xml:"ErrorMessage"`
	OriginType                    string                          `json:"OriginType" xml:"OriginType"`
	CreateTime                    string                          `json:"CreateTime" xml:"CreateTime"`
	EndTimestamp                  string                          `json:"EndTimestamp" xml:"EndTimestamp"`
	BeginTimestamp                string                          `json:"BeginTimestamp" xml:"BeginTimestamp"`
	ConsumptionClient             string                          `json:"ConsumptionClient" xml:"ConsumptionClient"`
	Reserved                      string                          `json:"Reserved" xml:"Reserved"`
	DbObject                      string                          `json:"DbObject" xml:"DbObject"`
	DtsJobId                      string                          `json:"DtsJobId" xml:"DtsJobId"`
	Delay                         int64                           `json:"Delay" xml:"Delay"`
	ExpireTime                    string                          `json:"ExpireTime" xml:"ExpireTime"`
	DtsJobName                    string                          `json:"DtsJobName" xml:"DtsJobName"`
	JobType                       string                          `json:"JobType" xml:"JobType"`
	DuUsage                       int64                           `json:"DuUsage" xml:"DuUsage"`
	Checkpoint                    string                          `json:"Checkpoint" xml:"Checkpoint"`
	PayType                       string                          `json:"PayType" xml:"PayType"`
	DtsJobClass                   string                          `json:"DtsJobClass" xml:"DtsJobClass"`
	CpuUsage                      string                          `json:"CpuUsage" xml:"CpuUsage"`
	Status                        string                          `json:"Status" xml:"Status"`
	ConsumptionCheckpoint         string                          `json:"ConsumptionCheckpoint" xml:"ConsumptionCheckpoint"`
	AppName                       string                          `json:"AppName" xml:"AppName"`
	DtsInstanceID                 string                          `json:"DtsInstanceID" xml:"DtsInstanceID"`
	DataEtlStatus                 DataEtlStatus                   `json:"DataEtlStatus" xml:"DataEtlStatus"`
	MigrationMode                 MigrationMode                   `json:"MigrationMode" xml:"MigrationMode"`
	ReverseJob                    ReverseJob                      `json:"ReverseJob" xml:"ReverseJob"`
	DataSynchronizationStatus     DataSynchronizationStatus       `json:"DataSynchronizationStatus" xml:"DataSynchronizationStatus"`
	SourceEndpoint                SourceEndpoint                  `json:"SourceEndpoint" xml:"SourceEndpoint"`
	PrecheckStatus                PrecheckStatusInDescribeDtsJobs `json:"PrecheckStatus" xml:"PrecheckStatus"`
	DataInitializationStatus      DataInitializationStatus        `json:"DataInitializationStatus" xml:"DataInitializationStatus"`
	Performance                   Performance                     `json:"Performance" xml:"Performance"`
	DestinationEndpoint           DestinationEndpoint             `json:"DestinationEndpoint" xml:"DestinationEndpoint"`
	StructureInitializationStatus StructureInitializationStatus   `json:"StructureInitializationStatus" xml:"StructureInitializationStatus"`
	RetryState                    RetryState                      `json:"RetryState" xml:"RetryState"`
	TagList                       []DtsTag                        `json:"TagList" xml:"TagList"`
	ErrorDetails                  []ErrorDetail                   `json:"ErrorDetails" xml:"ErrorDetails"`
}
