package dbs

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

// CreateGetDBListFromAgentTask invokes the dbs.CreateGetDBListFromAgentTask API synchronously
func (client *Client) CreateGetDBListFromAgentTask(request *CreateGetDBListFromAgentTaskRequest) (response *CreateGetDBListFromAgentTaskResponse, err error) {
	response = CreateCreateGetDBListFromAgentTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateGetDBListFromAgentTaskWithChan invokes the dbs.CreateGetDBListFromAgentTask API asynchronously
func (client *Client) CreateGetDBListFromAgentTaskWithChan(request *CreateGetDBListFromAgentTaskRequest) (<-chan *CreateGetDBListFromAgentTaskResponse, <-chan error) {
	responseChan := make(chan *CreateGetDBListFromAgentTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateGetDBListFromAgentTask(request)
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

// CreateGetDBListFromAgentTaskWithCallback invokes the dbs.CreateGetDBListFromAgentTask API asynchronously
func (client *Client) CreateGetDBListFromAgentTaskWithCallback(request *CreateGetDBListFromAgentTaskRequest, callback func(response *CreateGetDBListFromAgentTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateGetDBListFromAgentTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateGetDBListFromAgentTask(request)
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

// CreateGetDBListFromAgentTaskRequest is the request struct for api CreateGetDBListFromAgentTask
type CreateGetDBListFromAgentTaskRequest struct {
	*requests.RpcRequest
	SourceEndpointRegion string           `position:"Query" name:"SourceEndpointRegion"`
	BackupGatewayId      requests.Integer `position:"Query" name:"BackupGatewayId"`
	DatabaseType         string           `position:"Query" name:"DatabaseType"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	SourceEndpointPort   requests.Integer `position:"Query" name:"SourceEndpointPort"`
	SourceEndpointIP     string           `position:"Query" name:"SourceEndpointIP"`
}

// CreateGetDBListFromAgentTaskResponse is the response struct for api CreateGetDBListFromAgentTask
type CreateGetDBListFromAgentTaskResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	TaskId         int64  `json:"TaskId" xml:"TaskId"`
}

// CreateCreateGetDBListFromAgentTaskRequest creates a request to invoke CreateGetDBListFromAgentTask API
func CreateCreateGetDBListFromAgentTaskRequest() (request *CreateGetDBListFromAgentTaskRequest) {
	request = &CreateGetDBListFromAgentTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "CreateGetDBListFromAgentTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateGetDBListFromAgentTaskResponse creates a response to parse from CreateGetDBListFromAgentTask response
func CreateCreateGetDBListFromAgentTaskResponse() (response *CreateGetDBListFromAgentTaskResponse) {
	response = &CreateGetDBListFromAgentTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
