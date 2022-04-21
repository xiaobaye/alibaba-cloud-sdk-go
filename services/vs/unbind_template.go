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

// UnbindTemplate invokes the vs.UnbindTemplate API synchronously
func (client *Client) UnbindTemplate(request *UnbindTemplateRequest) (response *UnbindTemplateResponse, err error) {
	response = CreateUnbindTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindTemplateWithChan invokes the vs.UnbindTemplate API asynchronously
func (client *Client) UnbindTemplateWithChan(request *UnbindTemplateRequest) (<-chan *UnbindTemplateResponse, <-chan error) {
	responseChan := make(chan *UnbindTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindTemplate(request)
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

// UnbindTemplateWithCallback invokes the vs.UnbindTemplate API asynchronously
func (client *Client) UnbindTemplateWithCallback(request *UnbindTemplateRequest, callback func(response *UnbindTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindTemplateResponse
		var err error
		defer close(result)
		response, err = client.UnbindTemplate(request)
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

// UnbindTemplateRequest is the request struct for api UnbindTemplate
type UnbindTemplateRequest struct {
	*requests.RpcRequest
	TemplateType string           `position:"Query" name:"TemplateType"`
	InstanceType string           `position:"Query" name:"InstanceType"`
	ShowLog      string           `position:"Query" name:"ShowLog"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId   string           `position:"Query" name:"TemplateId"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
}

// UnbindTemplateResponse is the response struct for api UnbindTemplate
type UnbindTemplateResponse struct {
	*responses.BaseResponse
	TemplateType string `json:"TemplateType" xml:"TemplateType"`
	InstanceId   string `json:"InstanceId" xml:"InstanceId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	InstanceType string `json:"InstanceType" xml:"InstanceType"`
	TemplateId   string `json:"TemplateId" xml:"TemplateId"`
}

// CreateUnbindTemplateRequest creates a request to invoke UnbindTemplate API
func CreateUnbindTemplateRequest() (request *UnbindTemplateRequest) {
	request = &UnbindTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "UnbindTemplate", "", "")
	request.Method = requests.POST
	return
}

// CreateUnbindTemplateResponse creates a response to parse from UnbindTemplate response
func CreateUnbindTemplateResponse() (response *UnbindTemplateResponse) {
	response = &UnbindTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
