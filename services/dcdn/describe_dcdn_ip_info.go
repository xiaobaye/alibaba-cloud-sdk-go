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

// DescribeDcdnIpInfo invokes the dcdn.DescribeDcdnIpInfo API synchronously
func (client *Client) DescribeDcdnIpInfo(request *DescribeDcdnIpInfoRequest) (response *DescribeDcdnIpInfoResponse, err error) {
	response = CreateDescribeDcdnIpInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnIpInfoWithChan invokes the dcdn.DescribeDcdnIpInfo API asynchronously
func (client *Client) DescribeDcdnIpInfoWithChan(request *DescribeDcdnIpInfoRequest) (<-chan *DescribeDcdnIpInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnIpInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnIpInfo(request)
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

// DescribeDcdnIpInfoWithCallback invokes the dcdn.DescribeDcdnIpInfo API asynchronously
func (client *Client) DescribeDcdnIpInfoWithCallback(request *DescribeDcdnIpInfoRequest, callback func(response *DescribeDcdnIpInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnIpInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnIpInfo(request)
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

// DescribeDcdnIpInfoRequest is the request struct for api DescribeDcdnIpInfo
type DescribeDcdnIpInfoRequest struct {
	*requests.RpcRequest
	IP            string           `position:"Query" name:"IP"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DescribeDcdnIpInfoResponse is the response struct for api DescribeDcdnIpInfo
type DescribeDcdnIpInfoResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	RegionEname string `json:"RegionEname" xml:"RegionEname"`
	Region      string `json:"Region" xml:"Region"`
	IspEname    string `json:"IspEname" xml:"IspEname"`
	DcdnIp      string `json:"DcdnIp" xml:"DcdnIp"`
	ISP         string `json:"ISP" xml:"ISP"`
}

// CreateDescribeDcdnIpInfoRequest creates a request to invoke DescribeDcdnIpInfo API
func CreateDescribeDcdnIpInfoRequest() (request *DescribeDcdnIpInfoRequest) {
	request = &DescribeDcdnIpInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnIpInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnIpInfoResponse creates a response to parse from DescribeDcdnIpInfo response
func CreateDescribeDcdnIpInfoResponse() (response *DescribeDcdnIpInfoResponse) {
	response = &DescribeDcdnIpInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
