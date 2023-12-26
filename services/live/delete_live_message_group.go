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

// DeleteLiveMessageGroup invokes the live.DeleteLiveMessageGroup API synchronously
func (client *Client) DeleteLiveMessageGroup(request *DeleteLiveMessageGroupRequest) (response *DeleteLiveMessageGroupResponse, err error) {
	response = CreateDeleteLiveMessageGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveMessageGroupWithChan invokes the live.DeleteLiveMessageGroup API asynchronously
func (client *Client) DeleteLiveMessageGroupWithChan(request *DeleteLiveMessageGroupRequest) (<-chan *DeleteLiveMessageGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveMessageGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveMessageGroup(request)
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

// DeleteLiveMessageGroupWithCallback invokes the live.DeleteLiveMessageGroup API asynchronously
func (client *Client) DeleteLiveMessageGroupWithCallback(request *DeleteLiveMessageGroupRequest, callback func(response *DeleteLiveMessageGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveMessageGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveMessageGroup(request)
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

// DeleteLiveMessageGroupRequest is the request struct for api DeleteLiveMessageGroup
type DeleteLiveMessageGroupRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	DataCenter string `position:"Query" name:"DataCenter"`
	AppId      string `position:"Query" name:"AppId"`
	OperatorId string `position:"Query" name:"OperatorId"`
}

// DeleteLiveMessageGroupResponse is the response struct for api DeleteLiveMessageGroup
type DeleteLiveMessageGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	GroupId   string `json:"GroupId" xml:"GroupId"`
}

// CreateDeleteLiveMessageGroupRequest creates a request to invoke DeleteLiveMessageGroup API
func CreateDeleteLiveMessageGroupRequest() (request *DeleteLiveMessageGroupRequest) {
	request = &DeleteLiveMessageGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveMessageGroup", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLiveMessageGroupResponse creates a response to parse from DeleteLiveMessageGroup response
func CreateDeleteLiveMessageGroupResponse() (response *DeleteLiveMessageGroupResponse) {
	response = &DeleteLiveMessageGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
