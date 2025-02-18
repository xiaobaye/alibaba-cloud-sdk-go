package alikafka

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

// CreateConsumerGroup invokes the alikafka.CreateConsumerGroup API synchronously
func (client *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
	response = CreateCreateConsumerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateConsumerGroupWithChan invokes the alikafka.CreateConsumerGroup API asynchronously
func (client *Client) CreateConsumerGroupWithChan(request *CreateConsumerGroupRequest) (<-chan *CreateConsumerGroupResponse, <-chan error) {
	responseChan := make(chan *CreateConsumerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateConsumerGroup(request)
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

// CreateConsumerGroupWithCallback invokes the alikafka.CreateConsumerGroup API asynchronously
func (client *Client) CreateConsumerGroupWithCallback(request *CreateConsumerGroupRequest, callback func(response *CreateConsumerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateConsumerGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateConsumerGroup(request)
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

// CreateConsumerGroupRequest is the request struct for api CreateConsumerGroup
type CreateConsumerGroupRequest struct {
	*requests.RpcRequest
	ConsumerId string `position:"Query" name:"ConsumerId"`
	Remark     string `position:"Query" name:"Remark"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// CreateConsumerGroupResponse is the response struct for api CreateConsumerGroup
type CreateConsumerGroupResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateConsumerGroupRequest creates a request to invoke CreateConsumerGroup API
func CreateCreateConsumerGroupRequest() (request *CreateConsumerGroupRequest) {
	request = &CreateConsumerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "CreateConsumerGroup", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateConsumerGroupResponse creates a response to parse from CreateConsumerGroup response
func CreateCreateConsumerGroupResponse() (response *CreateConsumerGroupResponse) {
	response = &CreateConsumerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
