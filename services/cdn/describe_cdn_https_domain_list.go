package cdn

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeCdnHttpsDomainList invokes the cdn.DescribeCdnHttpsDomainList API synchronously
func (client *Client) DescribeCdnHttpsDomainList(request *DescribeCdnHttpsDomainListRequest) (response *DescribeCdnHttpsDomainListResponse, err error) {
	response = CreateDescribeCdnHttpsDomainListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCdnHttpsDomainListWithChan invokes the cdn.DescribeCdnHttpsDomainList API asynchronously
func (client *Client) DescribeCdnHttpsDomainListWithChan(request *DescribeCdnHttpsDomainListRequest) (<-chan *DescribeCdnHttpsDomainListResponse, <-chan error) {
	responseChan := make(chan *DescribeCdnHttpsDomainListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCdnHttpsDomainList(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeCdnHttpsDomainListWithCallback invokes the cdn.DescribeCdnHttpsDomainList API asynchronously
func (client *Client) DescribeCdnHttpsDomainListWithCallback(request *DescribeCdnHttpsDomainListRequest, callback func(response *DescribeCdnHttpsDomainListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCdnHttpsDomainListResponse
		var err error
		defer close(result)
		response, err = client.DescribeCdnHttpsDomainList(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeCdnHttpsDomainListRequest is the request struct for api DescribeCdnHttpsDomainList
type DescribeCdnHttpsDomainListRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Keyword    string           `position:"Query" name:"Keyword"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCdnHttpsDomainListResponse is the response struct for api DescribeCdnHttpsDomainList
type DescribeCdnHttpsDomainListResponse struct {
	*responses.BaseResponse
	TotalCount int                                   `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                                `json:"RequestId" xml:"RequestId"`
	CertInfos  CertInfosInDescribeCdnHttpsDomainList `json:"CertInfos" xml:"CertInfos"`
}

// CreateDescribeCdnHttpsDomainListRequest creates a request to invoke DescribeCdnHttpsDomainList API
func CreateDescribeCdnHttpsDomainListRequest() (request *DescribeCdnHttpsDomainListRequest) {
	request = &DescribeCdnHttpsDomainListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeCdnHttpsDomainList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCdnHttpsDomainListResponse creates a response to parse from DescribeCdnHttpsDomainList response
func CreateDescribeCdnHttpsDomainListResponse() (response *DescribeCdnHttpsDomainListResponse) {
	response = &DescribeCdnHttpsDomainListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
