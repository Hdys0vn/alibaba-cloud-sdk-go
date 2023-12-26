package cdn

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

// SetWaitingRoomConfig invokes the cdn.SetWaitingRoomConfig API synchronously
func (client *Client) SetWaitingRoomConfig(request *SetWaitingRoomConfigRequest) (response *SetWaitingRoomConfigResponse, err error) {
	response = CreateSetWaitingRoomConfigResponse()
	err = client.DoAction(request, response)
	return
}

// SetWaitingRoomConfigWithChan invokes the cdn.SetWaitingRoomConfig API asynchronously
func (client *Client) SetWaitingRoomConfigWithChan(request *SetWaitingRoomConfigRequest) (<-chan *SetWaitingRoomConfigResponse, <-chan error) {
	responseChan := make(chan *SetWaitingRoomConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetWaitingRoomConfig(request)
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

// SetWaitingRoomConfigWithCallback invokes the cdn.SetWaitingRoomConfig API asynchronously
func (client *Client) SetWaitingRoomConfigWithCallback(request *SetWaitingRoomConfigRequest, callback func(response *SetWaitingRoomConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetWaitingRoomConfigResponse
		var err error
		defer close(result)
		response, err = client.SetWaitingRoomConfig(request)
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

// SetWaitingRoomConfigRequest is the request struct for api SetWaitingRoomConfig
type SetWaitingRoomConfigRequest struct {
	*requests.RpcRequest
	MaxTimeWait requests.Integer `position:"Query" name:"MaxTimeWait"`
	WaitUrl     string           `position:"Query" name:"WaitUrl"`
	DomainName  string           `position:"Query" name:"DomainName"`
	AllowPct    requests.Integer `position:"Query" name:"AllowPct"`
	GapTime     requests.Integer `position:"Query" name:"GapTime"`
	WaitUri     string           `position:"Query" name:"WaitUri"`
}

// SetWaitingRoomConfigResponse is the response struct for api SetWaitingRoomConfig
type SetWaitingRoomConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetWaitingRoomConfigRequest creates a request to invoke SetWaitingRoomConfig API
func CreateSetWaitingRoomConfigRequest() (request *SetWaitingRoomConfigRequest) {
	request = &SetWaitingRoomConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "SetWaitingRoomConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateSetWaitingRoomConfigResponse creates a response to parse from SetWaitingRoomConfig response
func CreateSetWaitingRoomConfigResponse() (response *SetWaitingRoomConfigResponse) {
	response = &SetWaitingRoomConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
