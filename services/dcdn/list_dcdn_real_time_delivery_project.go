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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListDcdnRealTimeDeliveryProject invokes the dcdn.ListDcdnRealTimeDeliveryProject API synchronously
func (client *Client) ListDcdnRealTimeDeliveryProject(request *ListDcdnRealTimeDeliveryProjectRequest) (response *ListDcdnRealTimeDeliveryProjectResponse, err error) {
	response = CreateListDcdnRealTimeDeliveryProjectResponse()
	err = client.DoAction(request, response)
	return
}

// ListDcdnRealTimeDeliveryProjectWithChan invokes the dcdn.ListDcdnRealTimeDeliveryProject API asynchronously
func (client *Client) ListDcdnRealTimeDeliveryProjectWithChan(request *ListDcdnRealTimeDeliveryProjectRequest) (<-chan *ListDcdnRealTimeDeliveryProjectResponse, <-chan error) {
	responseChan := make(chan *ListDcdnRealTimeDeliveryProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDcdnRealTimeDeliveryProject(request)
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

// ListDcdnRealTimeDeliveryProjectWithCallback invokes the dcdn.ListDcdnRealTimeDeliveryProject API asynchronously
func (client *Client) ListDcdnRealTimeDeliveryProjectWithCallback(request *ListDcdnRealTimeDeliveryProjectRequest, callback func(response *ListDcdnRealTimeDeliveryProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDcdnRealTimeDeliveryProjectResponse
		var err error
		defer close(result)
		response, err = client.ListDcdnRealTimeDeliveryProject(request)
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

// ListDcdnRealTimeDeliveryProjectRequest is the request struct for api ListDcdnRealTimeDeliveryProject
type ListDcdnRealTimeDeliveryProjectRequest struct {
	*requests.RpcRequest
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	BusinessType string           `position:"Query" name:"BusinessType"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	DomainName   string           `position:"Query" name:"DomainName"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
}

// ListDcdnRealTimeDeliveryProjectResponse is the response struct for api ListDcdnRealTimeDeliveryProject
type ListDcdnRealTimeDeliveryProjectResponse struct {
	*responses.BaseResponse
	TotalCount int                                      `json:"TotalCount" xml:"TotalCount"`
	RequestId  string                                   `json:"RequestId" xml:"RequestId"`
	Content    ContentInListDcdnRealTimeDeliveryProject `json:"Content" xml:"Content"`
}

// CreateListDcdnRealTimeDeliveryProjectRequest creates a request to invoke ListDcdnRealTimeDeliveryProject API
func CreateListDcdnRealTimeDeliveryProjectRequest() (request *ListDcdnRealTimeDeliveryProjectRequest) {
	request = &ListDcdnRealTimeDeliveryProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "ListDcdnRealTimeDeliveryProject", "", "")
	request.Method = requests.POST
	return
}

// CreateListDcdnRealTimeDeliveryProjectResponse creates a response to parse from ListDcdnRealTimeDeliveryProject response
func CreateListDcdnRealTimeDeliveryProjectResponse() (response *ListDcdnRealTimeDeliveryProjectResponse) {
	response = &ListDcdnRealTimeDeliveryProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
