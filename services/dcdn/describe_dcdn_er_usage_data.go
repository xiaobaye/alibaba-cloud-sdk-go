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

// DescribeDcdnErUsageData invokes the dcdn.DescribeDcdnErUsageData API synchronously
func (client *Client) DescribeDcdnErUsageData(request *DescribeDcdnErUsageDataRequest) (response *DescribeDcdnErUsageDataResponse, err error) {
	response = CreateDescribeDcdnErUsageDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnErUsageDataWithChan invokes the dcdn.DescribeDcdnErUsageData API asynchronously
func (client *Client) DescribeDcdnErUsageDataWithChan(request *DescribeDcdnErUsageDataRequest) (<-chan *DescribeDcdnErUsageDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnErUsageDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnErUsageData(request)
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

// DescribeDcdnErUsageDataWithCallback invokes the dcdn.DescribeDcdnErUsageData API asynchronously
func (client *Client) DescribeDcdnErUsageDataWithCallback(request *DescribeDcdnErUsageDataRequest, callback func(response *DescribeDcdnErUsageDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnErUsageDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnErUsageData(request)
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

// DescribeDcdnErUsageDataRequest is the request struct for api DescribeDcdnErUsageData
type DescribeDcdnErUsageDataRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Spec      string           `position:"Query" name:"Spec"`
	RoutineID string           `position:"Query" name:"RoutineID"`
	SplitBy   string           `position:"Query" name:"SplitBy"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnErUsageDataResponse is the response struct for api DescribeDcdnErUsageData
type DescribeDcdnErUsageDataResponse struct {
	*responses.BaseResponse
	EndTime   string    `json:"EndTime" xml:"EndTime"`
	StartTime string    `json:"StartTime" xml:"StartTime"`
	RequestId string    `json:"RequestId" xml:"RequestId"`
	ErAccData ErAccData `json:"ErAccData" xml:"ErAccData"`
}

// CreateDescribeDcdnErUsageDataRequest creates a request to invoke DescribeDcdnErUsageData API
func CreateDescribeDcdnErUsageDataRequest() (request *DescribeDcdnErUsageDataRequest) {
	request = &DescribeDcdnErUsageDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnErUsageData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnErUsageDataResponse creates a response to parse from DescribeDcdnErUsageData response
func CreateDescribeDcdnErUsageDataResponse() (response *DescribeDcdnErUsageDataResponse) {
	response = &DescribeDcdnErUsageDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
