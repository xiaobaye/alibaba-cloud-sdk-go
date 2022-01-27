package das

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

// DeleteCloudBenchTask invokes the das.DeleteCloudBenchTask API synchronously
func (client *Client) DeleteCloudBenchTask(request *DeleteCloudBenchTaskRequest) (response *DeleteCloudBenchTaskResponse, err error) {
	response = CreateDeleteCloudBenchTaskResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCloudBenchTaskWithChan invokes the das.DeleteCloudBenchTask API asynchronously
func (client *Client) DeleteCloudBenchTaskWithChan(request *DeleteCloudBenchTaskRequest) (<-chan *DeleteCloudBenchTaskResponse, <-chan error) {
	responseChan := make(chan *DeleteCloudBenchTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCloudBenchTask(request)
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

// DeleteCloudBenchTaskWithCallback invokes the das.DeleteCloudBenchTask API asynchronously
func (client *Client) DeleteCloudBenchTaskWithCallback(request *DeleteCloudBenchTaskRequest, callback func(response *DeleteCloudBenchTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCloudBenchTaskResponse
		var err error
		defer close(result)
		response, err = client.DeleteCloudBenchTask(request)
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

// DeleteCloudBenchTaskRequest is the request struct for api DeleteCloudBenchTask
type DeleteCloudBenchTaskRequest struct {
	*requests.RpcRequest
	TaskId string `position:"Query" name:"TaskId"`
}

// DeleteCloudBenchTaskResponse is the response struct for api DeleteCloudBenchTask
type DeleteCloudBenchTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Success   string `json:"Success" xml:"Success"`
}

// CreateDeleteCloudBenchTaskRequest creates a request to invoke DeleteCloudBenchTask API
func CreateDeleteCloudBenchTaskRequest() (request *DeleteCloudBenchTaskRequest) {
	request = &DeleteCloudBenchTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "DeleteCloudBenchTask", "das", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCloudBenchTaskResponse creates a response to parse from DeleteCloudBenchTask response
func CreateDeleteCloudBenchTaskResponse() (response *DeleteCloudBenchTaskResponse) {
	response = &DeleteCloudBenchTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
