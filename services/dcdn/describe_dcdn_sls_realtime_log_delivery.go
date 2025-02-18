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

// DescribeDcdnSLSRealtimeLogDelivery invokes the dcdn.DescribeDcdnSLSRealtimeLogDelivery API synchronously
func (client *Client) DescribeDcdnSLSRealtimeLogDelivery(request *DescribeDcdnSLSRealtimeLogDeliveryRequest) (response *DescribeDcdnSLSRealtimeLogDeliveryResponse, err error) {
	response = CreateDescribeDcdnSLSRealtimeLogDeliveryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnSLSRealtimeLogDeliveryWithChan invokes the dcdn.DescribeDcdnSLSRealtimeLogDelivery API asynchronously
func (client *Client) DescribeDcdnSLSRealtimeLogDeliveryWithChan(request *DescribeDcdnSLSRealtimeLogDeliveryRequest) (<-chan *DescribeDcdnSLSRealtimeLogDeliveryResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnSLSRealtimeLogDeliveryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnSLSRealtimeLogDelivery(request)
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

// DescribeDcdnSLSRealtimeLogDeliveryWithCallback invokes the dcdn.DescribeDcdnSLSRealtimeLogDelivery API asynchronously
func (client *Client) DescribeDcdnSLSRealtimeLogDeliveryWithCallback(request *DescribeDcdnSLSRealtimeLogDeliveryRequest, callback func(response *DescribeDcdnSLSRealtimeLogDeliveryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnSLSRealtimeLogDeliveryResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnSLSRealtimeLogDelivery(request)
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

// DescribeDcdnSLSRealtimeLogDeliveryRequest is the request struct for api DescribeDcdnSLSRealtimeLogDelivery
type DescribeDcdnSLSRealtimeLogDeliveryRequest struct {
	*requests.RpcRequest
	ProjectName string           `position:"Query" name:"ProjectName"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnSLSRealtimeLogDeliveryResponse is the response struct for api DescribeDcdnSLSRealtimeLogDelivery
type DescribeDcdnSLSRealtimeLogDeliveryResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateDescribeDcdnSLSRealtimeLogDeliveryRequest creates a request to invoke DescribeDcdnSLSRealtimeLogDelivery API
func CreateDescribeDcdnSLSRealtimeLogDeliveryRequest() (request *DescribeDcdnSLSRealtimeLogDeliveryRequest) {
	request = &DescribeDcdnSLSRealtimeLogDeliveryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnSLSRealtimeLogDelivery", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnSLSRealtimeLogDeliveryResponse creates a response to parse from DescribeDcdnSLSRealtimeLogDelivery response
func CreateDescribeDcdnSLSRealtimeLogDeliveryResponse() (response *DescribeDcdnSLSRealtimeLogDeliveryResponse) {
	response = &DescribeDcdnSLSRealtimeLogDeliveryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
