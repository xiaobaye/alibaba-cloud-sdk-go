package cbn

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

// TransitRouter is a nested struct in cbn response
type TransitRouter struct {
	CreationTime             string `json:"CreationTime" xml:"CreationTime"`
	Type                     string `json:"Type" xml:"Type"`
	Status                   string `json:"Status" xml:"Status"`
	TransitRouterId          string `json:"TransitRouterId" xml:"TransitRouterId"`
	TransitRouterDescription string `json:"TransitRouterDescription" xml:"TransitRouterDescription"`
	TransitRouterName        string `json:"TransitRouterName" xml:"TransitRouterName"`
	CenId                    string `json:"CenId" xml:"CenId"`
	AliUid                   int64  `json:"AliUid" xml:"AliUid"`
	RegionId                 string `json:"RegionId" xml:"RegionId"`
	ServiceMode              string `json:"ServiceMode" xml:"ServiceMode"`
	SupportMulticast         bool   `json:"SupportMulticast" xml:"SupportMulticast"`
}
