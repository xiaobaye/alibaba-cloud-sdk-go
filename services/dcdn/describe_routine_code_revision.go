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

// DescribeRoutineCodeRevision invokes the dcdn.DescribeRoutineCodeRevision API synchronously
func (client *Client) DescribeRoutineCodeRevision(request *DescribeRoutineCodeRevisionRequest) (response *DescribeRoutineCodeRevisionResponse, err error) {
	response = CreateDescribeRoutineCodeRevisionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRoutineCodeRevisionWithChan invokes the dcdn.DescribeRoutineCodeRevision API asynchronously
func (client *Client) DescribeRoutineCodeRevisionWithChan(request *DescribeRoutineCodeRevisionRequest) (<-chan *DescribeRoutineCodeRevisionResponse, <-chan error) {
	responseChan := make(chan *DescribeRoutineCodeRevisionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRoutineCodeRevision(request)
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

// DescribeRoutineCodeRevisionWithCallback invokes the dcdn.DescribeRoutineCodeRevision API asynchronously
func (client *Client) DescribeRoutineCodeRevisionWithCallback(request *DescribeRoutineCodeRevisionRequest, callback func(response *DescribeRoutineCodeRevisionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRoutineCodeRevisionResponse
		var err error
		defer close(result)
		response, err = client.DescribeRoutineCodeRevision(request)
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

// DescribeRoutineCodeRevisionRequest is the request struct for api DescribeRoutineCodeRevision
type DescribeRoutineCodeRevisionRequest struct {
	*requests.RpcRequest
	SelectCodeRevision string           `position:"Body" name:"SelectCodeRevision"`
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
	Name               string           `position:"Body" name:"Name"`
}

// DescribeRoutineCodeRevisionResponse is the response struct for api DescribeRoutineCodeRevision
type DescribeRoutineCodeRevisionResponse struct {
	*responses.BaseResponse
	Content   map[string]interface{} `json:"Content" xml:"Content"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeRoutineCodeRevisionRequest creates a request to invoke DescribeRoutineCodeRevision API
func CreateDescribeRoutineCodeRevisionRequest() (request *DescribeRoutineCodeRevisionRequest) {
	request = &DescribeRoutineCodeRevisionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeRoutineCodeRevision", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRoutineCodeRevisionResponse creates a response to parse from DescribeRoutineCodeRevision response
func CreateDescribeRoutineCodeRevisionResponse() (response *DescribeRoutineCodeRevisionResponse) {
	response = &DescribeRoutineCodeRevisionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
