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

// AddCasterEpisodeGroup invokes the live.AddCasterEpisodeGroup API synchronously
func (client *Client) AddCasterEpisodeGroup(request *AddCasterEpisodeGroupRequest) (response *AddCasterEpisodeGroupResponse, err error) {
	response = CreateAddCasterEpisodeGroupResponse()
	err = client.DoAction(request, response)
	return
}

// AddCasterEpisodeGroupWithChan invokes the live.AddCasterEpisodeGroup API asynchronously
func (client *Client) AddCasterEpisodeGroupWithChan(request *AddCasterEpisodeGroupRequest) (<-chan *AddCasterEpisodeGroupResponse, <-chan error) {
	responseChan := make(chan *AddCasterEpisodeGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCasterEpisodeGroup(request)
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

// AddCasterEpisodeGroupWithCallback invokes the live.AddCasterEpisodeGroup API asynchronously
func (client *Client) AddCasterEpisodeGroupWithCallback(request *AddCasterEpisodeGroupRequest, callback func(response *AddCasterEpisodeGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCasterEpisodeGroupResponse
		var err error
		defer close(result)
		response, err = client.AddCasterEpisodeGroup(request)
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

// AddCasterEpisodeGroupRequest is the request struct for api AddCasterEpisodeGroup
type AddCasterEpisodeGroupRequest struct {
	*requests.RpcRequest
	ClientToken   string                       `position:"Query" name:"ClientToken"`
	StartTime     string                       `position:"Query" name:"StartTime"`
	SideOutputUrl string                       `position:"Query" name:"SideOutputUrl"`
	Item          *[]AddCasterEpisodeGroupItem `position:"Query" name:"Item"  type:"Repeated"`
	DomainName    string                       `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer             `position:"Query" name:"OwnerId"`
	RepeatNum     requests.Integer             `position:"Query" name:"RepeatNum"`
	CallbackUrl   string                       `position:"Query" name:"CallbackUrl"`
}

// AddCasterEpisodeGroupItem is a repeated param struct in AddCasterEpisodeGroupRequest
type AddCasterEpisodeGroupItem struct {
	ItemName string `name:"ItemName"`
	VodUrl   string `name:"VodUrl"`
}

// AddCasterEpisodeGroupResponse is the response struct for api AddCasterEpisodeGroup
type AddCasterEpisodeGroupResponse struct {
	*responses.BaseResponse
	ProgramId string                         `json:"ProgramId" xml:"ProgramId"`
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	ItemIds   ItemIdsInAddCasterEpisodeGroup `json:"ItemIds" xml:"ItemIds"`
}

// CreateAddCasterEpisodeGroupRequest creates a request to invoke AddCasterEpisodeGroup API
func CreateAddCasterEpisodeGroupRequest() (request *AddCasterEpisodeGroupRequest) {
	request = &AddCasterEpisodeGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddCasterEpisodeGroup", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddCasterEpisodeGroupResponse creates a response to parse from AddCasterEpisodeGroup response
func CreateAddCasterEpisodeGroupResponse() (response *AddCasterEpisodeGroupResponse) {
	response = &AddCasterEpisodeGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
