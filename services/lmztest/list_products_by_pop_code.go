package lmztest

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

// ListProductsByPopCode invokes the lmztest.ListProductsByPopCode API synchronously
func (client *Client) ListProductsByPopCode(request *ListProductsByPopCodeRequest) (response *ListProductsByPopCodeResponse, err error) {
	response = CreateListProductsByPopCodeResponse()
	err = client.DoAction(request, response)
	return
}

// ListProductsByPopCodeWithChan invokes the lmztest.ListProductsByPopCode API asynchronously
func (client *Client) ListProductsByPopCodeWithChan(request *ListProductsByPopCodeRequest) (<-chan *ListProductsByPopCodeResponse, <-chan error) {
	responseChan := make(chan *ListProductsByPopCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProductsByPopCode(request)
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

// ListProductsByPopCodeWithCallback invokes the lmztest.ListProductsByPopCode API asynchronously
func (client *Client) ListProductsByPopCodeWithCallback(request *ListProductsByPopCodeRequest, callback func(response *ListProductsByPopCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProductsByPopCodeResponse
		var err error
		defer close(result)
		response, err = client.ListProductsByPopCode(request)
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

// ListProductsByPopCodeRequest is the request struct for api ListProductsByPopCode
type ListProductsByPopCodeRequest struct {
	*requests.RpcRequest
	AuthKey        string `position:"Query" name:"AuthKey"`
	ExtraParams    string `position:"Body" name:"ExtraParams"`
	ServerClientIp string `position:"Query" name:"ServerClientIp"`
	RequestId      string `position:"Query" name:"RequestId"`
	ClientIp       string `position:"Query" name:"ClientIp"`
	IdentityDTO    string `position:"Body" name:"IdentityDTO"`
	Env            string `position:"Query" name:"Env"`
	PopCode        string `position:"Query" name:"PopCode"`
}

// ListProductsByPopCodeResponse is the response struct for api ListProductsByPopCode
type ListProductsByPopCodeResponse struct {
	*responses.BaseResponse
	ErrorMessage       string       `json:"ErrorMessage" xml:"ErrorMessage"`
	TraceId            string       `json:"TraceId" xml:"TraceId"`
	ErrorDetailMessage string       `json:"ErrorDetailMessage" xml:"ErrorDetailMessage"`
	Success            bool         `json:"Success" xml:"Success"`
	ErrorCode          string       `json:"ErrorCode" xml:"ErrorCode"`
	Result             []ResultItem `json:"Result" xml:"Result"`
}

// CreateListProductsByPopCodeRequest creates a request to invoke ListProductsByPopCode API
func CreateListProductsByPopCodeRequest() (request *ListProductsByPopCodeRequest) {
	request = &ListProductsByPopCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LmzTest", "2010-10-11", "ListProductsByPopCode", "", "")
	request.Method = requests.POST
	return
}

// CreateListProductsByPopCodeResponse creates a response to parse from ListProductsByPopCode response
func CreateListProductsByPopCodeResponse() (response *ListProductsByPopCodeResponse) {
	response = &ListProductsByPopCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
