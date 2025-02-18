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

// DescribeUserDcdnStatus invokes the dcdn.DescribeUserDcdnStatus API synchronously
func (client *Client) DescribeUserDcdnStatus(request *DescribeUserDcdnStatusRequest) (response *DescribeUserDcdnStatusResponse, err error) {
	response = CreateDescribeUserDcdnStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserDcdnStatusWithChan invokes the dcdn.DescribeUserDcdnStatus API asynchronously
func (client *Client) DescribeUserDcdnStatusWithChan(request *DescribeUserDcdnStatusRequest) (<-chan *DescribeUserDcdnStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeUserDcdnStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserDcdnStatus(request)
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

// DescribeUserDcdnStatusWithCallback invokes the dcdn.DescribeUserDcdnStatus API asynchronously
func (client *Client) DescribeUserDcdnStatusWithCallback(request *DescribeUserDcdnStatusRequest, callback func(response *DescribeUserDcdnStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserDcdnStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserDcdnStatus(request)
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

// DescribeUserDcdnStatusRequest is the request struct for api DescribeUserDcdnStatus
type DescribeUserDcdnStatusRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeUserDcdnStatusResponse is the response struct for api DescribeUserDcdnStatus
type DescribeUserDcdnStatusResponse struct {
	*responses.BaseResponse
	InDebt        bool   `json:"InDebt" xml:"InDebt"`
	OnService     bool   `json:"OnService" xml:"OnService"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
	InDebtOverdue bool   `json:"InDebtOverdue" xml:"InDebtOverdue"`
	Enabled       bool   `json:"Enabled" xml:"Enabled"`
}

// CreateDescribeUserDcdnStatusRequest creates a request to invoke DescribeUserDcdnStatus API
func CreateDescribeUserDcdnStatusRequest() (request *DescribeUserDcdnStatusRequest) {
	request = &DescribeUserDcdnStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeUserDcdnStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeUserDcdnStatusResponse creates a response to parse from DescribeUserDcdnStatus response
func CreateDescribeUserDcdnStatusResponse() (response *DescribeUserDcdnStatusResponse) {
	response = &DescribeUserDcdnStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
