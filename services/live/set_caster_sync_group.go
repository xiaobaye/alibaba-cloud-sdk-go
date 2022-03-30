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

// SetCasterSyncGroup invokes the live.SetCasterSyncGroup API synchronously
func (client *Client) SetCasterSyncGroup(request *SetCasterSyncGroupRequest) (response *SetCasterSyncGroupResponse, err error) {
	response = CreateSetCasterSyncGroupResponse()
	err = client.DoAction(request, response)
	return
}

// SetCasterSyncGroupWithChan invokes the live.SetCasterSyncGroup API asynchronously
func (client *Client) SetCasterSyncGroupWithChan(request *SetCasterSyncGroupRequest) (<-chan *SetCasterSyncGroupResponse, <-chan error) {
	responseChan := make(chan *SetCasterSyncGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetCasterSyncGroup(request)
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

// SetCasterSyncGroupWithCallback invokes the live.SetCasterSyncGroup API asynchronously
func (client *Client) SetCasterSyncGroupWithCallback(request *SetCasterSyncGroupRequest, callback func(response *SetCasterSyncGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetCasterSyncGroupResponse
		var err error
		defer close(result)
		response, err = client.SetCasterSyncGroup(request)
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

// SetCasterSyncGroupRequest is the request struct for api SetCasterSyncGroup
type SetCasterSyncGroupRequest struct {
	*requests.RpcRequest
	CasterId  string                         `position:"Query" name:"CasterId"`
	OwnerId   requests.Integer               `position:"Query" name:"OwnerId"`
	SyncGroup *[]SetCasterSyncGroupSyncGroup `position:"Query" name:"SyncGroup"  type:"Repeated"`
}

// SetCasterSyncGroupSyncGroup is a repeated param struct in SetCasterSyncGroupRequest
type SetCasterSyncGroupSyncGroup struct {
	HostResourceId     string    `name:"HostResourceId"`
	ResourceIds        *[]string `name:"ResourceIds" type:"Repeated"`
	SyncOffsets        *[]string `name:"SyncOffsets" type:"Repeated"`
	SyncDelayThreshold string    `name:"SyncDelayThreshold"`
	Mode               string    `name:"Mode"`
}

// SetCasterSyncGroupResponse is the response struct for api SetCasterSyncGroup
type SetCasterSyncGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetCasterSyncGroupRequest creates a request to invoke SetCasterSyncGroup API
func CreateSetCasterSyncGroupRequest() (request *SetCasterSyncGroupRequest) {
	request = &SetCasterSyncGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "SetCasterSyncGroup", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetCasterSyncGroupResponse creates a response to parse from SetCasterSyncGroup response
func CreateSetCasterSyncGroupResponse() (response *SetCasterSyncGroupResponse) {
	response = &SetCasterSyncGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
