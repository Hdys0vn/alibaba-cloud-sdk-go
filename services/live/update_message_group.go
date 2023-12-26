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
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateMessageGroup invokes the live.UpdateMessageGroup API synchronously
func (client *Client) UpdateMessageGroup(request *UpdateMessageGroupRequest) (response *UpdateMessageGroupResponse, err error) {
	response = CreateUpdateMessageGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMessageGroupWithChan invokes the live.UpdateMessageGroup API asynchronously
func (client *Client) UpdateMessageGroupWithChan(request *UpdateMessageGroupRequest) (<-chan *UpdateMessageGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateMessageGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMessageGroup(request)
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

// UpdateMessageGroupWithCallback invokes the live.UpdateMessageGroup API asynchronously
func (client *Client) UpdateMessageGroupWithCallback(request *UpdateMessageGroupRequest, callback func(response *UpdateMessageGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMessageGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateMessageGroup(request)
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

// UpdateMessageGroupRequest is the request struct for api UpdateMessageGroup
type UpdateMessageGroupRequest struct {
	*requests.RpcRequest
	Extension map[string]string `position:"Body" name:"Extension"  type:"Map"`
	GroupId   string            `position:"Body" name:"GroupId"`
	AppId     string            `position:"Body" name:"AppId"`
}

// UpdateMessageGroupResponse is the response struct for api UpdateMessageGroup
type UpdateMessageGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdateMessageGroupRequest creates a request to invoke UpdateMessageGroup API
func CreateUpdateMessageGroupRequest() (request *UpdateMessageGroupRequest) {
	request = &UpdateMessageGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateMessageGroup", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateMessageGroupResponse creates a response to parse from UpdateMessageGroup response
func CreateUpdateMessageGroupResponse() (response *UpdateMessageGroupResponse) {
	response = &UpdateMessageGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}