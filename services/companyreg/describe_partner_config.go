package companyreg

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

// DescribePartnerConfig invokes the companyreg.DescribePartnerConfig API synchronously
func (client *Client) DescribePartnerConfig(request *DescribePartnerConfigRequest) (response *DescribePartnerConfigResponse, err error) {
	response = CreateDescribePartnerConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePartnerConfigWithChan invokes the companyreg.DescribePartnerConfig API asynchronously
func (client *Client) DescribePartnerConfigWithChan(request *DescribePartnerConfigRequest) (<-chan *DescribePartnerConfigResponse, <-chan error) {
	responseChan := make(chan *DescribePartnerConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePartnerConfig(request)
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

// DescribePartnerConfigWithCallback invokes the companyreg.DescribePartnerConfig API asynchronously
func (client *Client) DescribePartnerConfigWithCallback(request *DescribePartnerConfigRequest, callback func(response *DescribePartnerConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePartnerConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribePartnerConfig(request)
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

// DescribePartnerConfigRequest is the request struct for api DescribePartnerConfig
type DescribePartnerConfigRequest struct {
	*requests.RpcRequest
	PartnerCode string `position:"Query" name:"PartnerCode"`
	BizType     string `position:"Query" name:"BizType"`
}

// DescribePartnerConfigResponse is the response struct for api DescribePartnerConfig
type DescribePartnerConfigResponse struct {
	*responses.BaseResponse
	PartnerName string `json:"PartnerName" xml:"PartnerName"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	PartnerCode string `json:"PartnerCode" xml:"PartnerCode"`
	Contact     string `json:"Contact" xml:"Contact"`
}

// CreateDescribePartnerConfigRequest creates a request to invoke DescribePartnerConfig API
func CreateDescribePartnerConfigRequest() (request *DescribePartnerConfigRequest) {
	request = &DescribePartnerConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-03-06", "DescribePartnerConfig", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePartnerConfigResponse creates a response to parse from DescribePartnerConfig response
func CreateDescribePartnerConfigResponse() (response *DescribePartnerConfigResponse) {
	response = &DescribePartnerConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
