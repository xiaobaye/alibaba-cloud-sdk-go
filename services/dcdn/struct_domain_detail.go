package dcdn

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

// DomainDetail is a nested struct in dcdn response
type DomainDetail struct {
	Cname           string                            `json:"Cname" xml:"Cname"`
	CertName        string                            `json:"CertName" xml:"CertName"`
	DomainStatus    string                            `json:"DomainStatus" xml:"DomainStatus"`
	DomainName      string                            `json:"DomainName" xml:"DomainName"`
	GmtModified     string                            `json:"GmtModified" xml:"GmtModified"`
	SSLPub          string                            `json:"SSLPub" xml:"SSLPub"`
	SSLProtocol     string                            `json:"SSLProtocol" xml:"SSLProtocol"`
	GmtCreated      string                            `json:"GmtCreated" xml:"GmtCreated"`
	ResourceGroupId string                            `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Scope           string                            `json:"Scope" xml:"Scope"`
	Description     string                            `json:"Description" xml:"Description"`
	TenantID        string                            `json:"TenantID" xml:"TenantID"`
	Sources         SourcesInDescribeDcdnDomainDetail `json:"Sources" xml:"Sources"`
}
