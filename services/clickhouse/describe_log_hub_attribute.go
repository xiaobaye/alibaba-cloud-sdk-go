package clickhouse

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

// DescribeLogHubAttribute invokes the clickhouse.DescribeLogHubAttribute API synchronously
func (client *Client) DescribeLogHubAttribute(request *DescribeLogHubAttributeRequest) (response *DescribeLogHubAttributeResponse, err error) {
	response = CreateDescribeLogHubAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLogHubAttributeWithChan invokes the clickhouse.DescribeLogHubAttribute API asynchronously
func (client *Client) DescribeLogHubAttributeWithChan(request *DescribeLogHubAttributeRequest) (<-chan *DescribeLogHubAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeLogHubAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLogHubAttribute(request)
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

// DescribeLogHubAttributeWithCallback invokes the clickhouse.DescribeLogHubAttribute API asynchronously
func (client *Client) DescribeLogHubAttributeWithCallback(request *DescribeLogHubAttributeRequest, callback func(response *DescribeLogHubAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLogHubAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLogHubAttribute(request)
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

// DescribeLogHubAttributeRequest is the request struct for api DescribeLogHubAttribute
type DescribeLogHubAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DeliverName          string           `position:"Query" name:"DeliverName"`
	TaskId               string           `position:"Query" name:"TaskId"`
	ProjectName          string           `position:"Query" name:"ProjectName"`
	LogStoreName         string           `position:"Query" name:"LogStoreName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLogHubAttributeResponse is the response struct for api DescribeLogHubAttribute
type DescribeLogHubAttributeResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	LoghubInfo LoghubInfo `json:"LoghubInfo" xml:"LoghubInfo"`
}

// CreateDescribeLogHubAttributeRequest creates a request to invoke DescribeLogHubAttribute API
func CreateDescribeLogHubAttributeRequest() (request *DescribeLogHubAttributeRequest) {
	request = &DescribeLogHubAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeLogHubAttribute", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLogHubAttributeResponse creates a response to parse from DescribeLogHubAttribute response
func CreateDescribeLogHubAttributeResponse() (response *DescribeLogHubAttributeResponse) {
	response = &DescribeLogHubAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
