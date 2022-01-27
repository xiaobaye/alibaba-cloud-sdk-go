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

// GetRequestDiagnosisPage invokes the das.GetRequestDiagnosisPage API synchronously
func (client *Client) GetRequestDiagnosisPage(request *GetRequestDiagnosisPageRequest) (response *GetRequestDiagnosisPageResponse, err error) {
	response = CreateGetRequestDiagnosisPageResponse()
	err = client.DoAction(request, response)
	return
}

// GetRequestDiagnosisPageWithChan invokes the das.GetRequestDiagnosisPage API asynchronously
func (client *Client) GetRequestDiagnosisPageWithChan(request *GetRequestDiagnosisPageRequest) (<-chan *GetRequestDiagnosisPageResponse, <-chan error) {
	responseChan := make(chan *GetRequestDiagnosisPageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRequestDiagnosisPage(request)
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

// GetRequestDiagnosisPageWithCallback invokes the das.GetRequestDiagnosisPage API asynchronously
func (client *Client) GetRequestDiagnosisPageWithCallback(request *GetRequestDiagnosisPageRequest, callback func(response *GetRequestDiagnosisPageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRequestDiagnosisPageResponse
		var err error
		defer close(result)
		response, err = client.GetRequestDiagnosisPage(request)
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

// GetRequestDiagnosisPageRequest is the request struct for api GetRequestDiagnosisPage
type GetRequestDiagnosisPageRequest struct {
	*requests.RpcRequest
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	InstanceId string `position:"Query" name:"InstanceId"`
	PageNo     string `position:"Query" name:"PageNo"`
	PageSize   string `position:"Query" name:"PageSize"`
	NodeId     string `position:"Query" name:"NodeId"`
}

// GetRequestDiagnosisPageResponse is the response struct for api GetRequestDiagnosisPage
type GetRequestDiagnosisPageResponse struct {
	*responses.BaseResponse
	Message   string                        `json:"Message" xml:"Message"`
	RequestId string                        `json:"RequestId" xml:"RequestId"`
	Code      string                        `json:"Code" xml:"Code"`
	Success   string                        `json:"Success" xml:"Success"`
	Data      DataInGetRequestDiagnosisPage `json:"Data" xml:"Data"`
}

// CreateGetRequestDiagnosisPageRequest creates a request to invoke GetRequestDiagnosisPage API
func CreateGetRequestDiagnosisPageRequest() (request *GetRequestDiagnosisPageRequest) {
	request = &GetRequestDiagnosisPageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetRequestDiagnosisPage", "das", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetRequestDiagnosisPageResponse creates a response to parse from GetRequestDiagnosisPage response
func CreateGetRequestDiagnosisPageResponse() (response *GetRequestDiagnosisPageResponse) {
	response = &GetRequestDiagnosisPageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
