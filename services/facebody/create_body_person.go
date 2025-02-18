package facebody

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

// CreateBodyPerson invokes the facebody.CreateBodyPerson API synchronously
func (client *Client) CreateBodyPerson(request *CreateBodyPersonRequest) (response *CreateBodyPersonResponse, err error) {
	response = CreateCreateBodyPersonResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBodyPersonWithChan invokes the facebody.CreateBodyPerson API asynchronously
func (client *Client) CreateBodyPersonWithChan(request *CreateBodyPersonRequest) (<-chan *CreateBodyPersonResponse, <-chan error) {
	responseChan := make(chan *CreateBodyPersonResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBodyPerson(request)
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

// CreateBodyPersonWithCallback invokes the facebody.CreateBodyPerson API asynchronously
func (client *Client) CreateBodyPersonWithCallback(request *CreateBodyPersonRequest, callback func(response *CreateBodyPersonResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBodyPersonResponse
		var err error
		defer close(result)
		response, err = client.CreateBodyPerson(request)
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

// CreateBodyPersonRequest is the request struct for api CreateBodyPerson
type CreateBodyPersonRequest struct {
	*requests.RpcRequest
	FormatResultToJson requests.Boolean `position:"Query" name:"FormatResultToJson"`
	OssFile            string           `position:"Query" name:"OssFile"`
	RequestProxyBy     string           `position:"Query" name:"RequestProxyBy"`
	DbId               requests.Integer `position:"Body" name:"DbId"`
	Name               string           `position:"Body" name:"Name"`
}

// CreateBodyPersonResponse is the response struct for api CreateBodyPerson
type CreateBodyPersonResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateBodyPersonRequest creates a request to invoke CreateBodyPerson API
func CreateCreateBodyPersonRequest() (request *CreateBodyPersonRequest) {
	request = &CreateBodyPersonRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "CreateBodyPerson", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBodyPersonResponse creates a response to parse from CreateBodyPerson response
func CreateCreateBodyPersonResponse() (response *CreateBodyPersonResponse) {
	response = &CreateBodyPersonResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
