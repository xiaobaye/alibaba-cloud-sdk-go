package ecs

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

// InvokeCommand invokes the ecs.InvokeCommand API synchronously
func (client *Client) InvokeCommand(request *InvokeCommandRequest) (response *InvokeCommandResponse, err error) {
	response = CreateInvokeCommandResponse()
	err = client.DoAction(request, response)
	return
}

// InvokeCommandWithChan invokes the ecs.InvokeCommand API asynchronously
func (client *Client) InvokeCommandWithChan(request *InvokeCommandRequest) (<-chan *InvokeCommandResponse, <-chan error) {
	responseChan := make(chan *InvokeCommandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InvokeCommand(request)
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

// InvokeCommandWithCallback invokes the ecs.InvokeCommand API asynchronously
func (client *Client) InvokeCommandWithCallback(request *InvokeCommandRequest, callback func(response *InvokeCommandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InvokeCommandResponse
		var err error
		defer close(result)
		response, err = client.InvokeCommand(request)
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

// InvokeCommandRequest is the request struct for api InvokeCommand
type InvokeCommandRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer       `position:"Query" name:"ResourceOwnerId"`
	CommandId            string                 `position:"Query" name:"CommandId"`
	Frequency            string                 `position:"Query" name:"Frequency"`
	RepeatMode           string                 `position:"Query" name:"RepeatMode"`
	WindowsPasswordName  string                 `position:"Query" name:"WindowsPasswordName"`
	Timed                requests.Boolean       `position:"Query" name:"Timed"`
	ResourceOwnerAccount string                 `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                 `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer       `position:"Query" name:"OwnerId"`
	InstanceId           *[]string              `position:"Query" name:"InstanceId"  type:"Repeated"`
	ContainerId          string                 `position:"Query" name:"ContainerId"`
	Parameters           map[string]interface{} `position:"Query" name:"Parameters"`
	Username             string                 `position:"Query" name:"Username"`
}

// InvokeCommandResponse is the response struct for api InvokeCommand
type InvokeCommandResponse struct {
	*responses.BaseResponse
	InvokeId  string `json:"InvokeId" xml:"InvokeId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateInvokeCommandRequest creates a request to invoke InvokeCommand API
func CreateInvokeCommandRequest() (request *InvokeCommandRequest) {
	request = &InvokeCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "InvokeCommand", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInvokeCommandResponse creates a response to parse from InvokeCommand response
func CreateInvokeCommandResponse() (response *InvokeCommandResponse) {
	response = &InvokeCommandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
