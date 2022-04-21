package vs

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

// DescribeStreams invokes the vs.DescribeStreams API synchronously
func (client *Client) DescribeStreams(request *DescribeStreamsRequest) (response *DescribeStreamsResponse, err error) {
	response = CreateDescribeStreamsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStreamsWithChan invokes the vs.DescribeStreams API asynchronously
func (client *Client) DescribeStreamsWithChan(request *DescribeStreamsRequest) (<-chan *DescribeStreamsResponse, <-chan error) {
	responseChan := make(chan *DescribeStreamsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStreams(request)
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

// DescribeStreamsWithCallback invokes the vs.DescribeStreams API asynchronously
func (client *Client) DescribeStreamsWithCallback(request *DescribeStreamsRequest, callback func(response *DescribeStreamsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStreamsResponse
		var err error
		defer close(result)
		response, err = client.DescribeStreams(request)
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

// DescribeStreamsRequest is the request struct for api DescribeStreams
type DescribeStreamsRequest struct {
	*requests.RpcRequest
	SortDirection string           `position:"Query" name:"SortDirection"`
	PageNum       requests.Integer `position:"Query" name:"PageNum"`
	ParentId      string           `position:"Query" name:"ParentId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	Id            string           `position:"Query" name:"Id"`
	ShowLog       string           `position:"Query" name:"ShowLog"`
	App           string           `position:"Query" name:"App"`
	GroupId       string           `position:"Query" name:"GroupId"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	DeviceId      string           `position:"Query" name:"DeviceId"`
	Domain        string           `position:"Query" name:"Domain"`
	Name          string           `position:"Query" name:"Name"`
	SortBy        string           `position:"Query" name:"SortBy"`
}

// DescribeStreamsResponse is the response struct for api DescribeStreams
type DescribeStreamsResponse struct {
	*responses.BaseResponse
	PageNum    int64    `json:"PageNum" xml:"PageNum"`
	PageSize   int64    `json:"PageSize" xml:"PageSize"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount int64    `json:"TotalCount" xml:"TotalCount"`
	PageCount  int64    `json:"PageCount" xml:"PageCount"`
	Streams    []Stream `json:"Streams" xml:"Streams"`
}

// CreateDescribeStreamsRequest creates a request to invoke DescribeStreams API
func CreateDescribeStreamsRequest() (request *DescribeStreamsRequest) {
	request = &DescribeStreamsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "DescribeStreams", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeStreamsResponse creates a response to parse from DescribeStreams response
func CreateDescribeStreamsResponse() (response *DescribeStreamsResponse) {
	response = &DescribeStreamsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
