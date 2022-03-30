package live

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

// StartLiveIndex invokes the live.StartLiveIndex API synchronously
func (client *Client) StartLiveIndex(request *StartLiveIndexRequest) (response *StartLiveIndexResponse, err error) {
	response = CreateStartLiveIndexResponse()
	err = client.DoAction(request, response)
	return
}

// StartLiveIndexWithChan invokes the live.StartLiveIndex API asynchronously
func (client *Client) StartLiveIndexWithChan(request *StartLiveIndexRequest) (<-chan *StartLiveIndexResponse, <-chan error) {
	responseChan := make(chan *StartLiveIndexResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartLiveIndex(request)
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

// StartLiveIndexWithCallback invokes the live.StartLiveIndex API asynchronously
func (client *Client) StartLiveIndexWithCallback(request *StartLiveIndexRequest, callback func(response *StartLiveIndexResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartLiveIndexResponse
		var err error
		defer close(result)
		response, err = client.StartLiveIndex(request)
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

// StartLiveIndexRequest is the request struct for api StartLiveIndex
type StartLiveIndexRequest struct {
	*requests.RpcRequest
	TokenId     string           `position:"Query" name:"TokenId"`
	OssEndpoint string           `position:"Query" name:"OssEndpoint"`
	AppName     string           `position:"Query" name:"AppName"`
	OssRamRole  string           `position:"Query" name:"OssRamRole"`
	StreamName  string           `position:"Query" name:"StreamName"`
	OssUserId   string           `position:"Query" name:"OssUserId"`
	OssBucket   string           `position:"Query" name:"OssBucket"`
	DomainName  string           `position:"Query" name:"DomainName"`
	InputUrl    string           `position:"Query" name:"InputUrl"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	Interval    requests.Integer `position:"Query" name:"Interval"`
}

// StartLiveIndexResponse is the response struct for api StartLiveIndex
type StartLiveIndexResponse struct {
	*responses.BaseResponse
	TaskId    string `json:"TaskId" xml:"TaskId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStartLiveIndexRequest creates a request to invoke StartLiveIndex API
func CreateStartLiveIndexRequest() (request *StartLiveIndexRequest) {
	request = &StartLiveIndexRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "StartLiveIndex", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartLiveIndexResponse creates a response to parse from StartLiveIndex response
func CreateStartLiveIndexResponse() (response *StartLiveIndexResponse) {
	response = &StartLiveIndexResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
