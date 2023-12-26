package aliyuncvc

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

// ConferenceToLive invokes the aliyuncvc.ConferenceToLive API synchronously
func (client *Client) ConferenceToLive(request *ConferenceToLiveRequest) (response *ConferenceToLiveResponse, err error) {
	response = CreateConferenceToLiveResponse()
	err = client.DoAction(request, response)
	return
}

// ConferenceToLiveWithChan invokes the aliyuncvc.ConferenceToLive API asynchronously
func (client *Client) ConferenceToLiveWithChan(request *ConferenceToLiveRequest) (<-chan *ConferenceToLiveResponse, <-chan error) {
	responseChan := make(chan *ConferenceToLiveResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConferenceToLive(request)
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

// ConferenceToLiveWithCallback invokes the aliyuncvc.ConferenceToLive API asynchronously
func (client *Client) ConferenceToLiveWithCallback(request *ConferenceToLiveRequest, callback func(response *ConferenceToLiveResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConferenceToLiveResponse
		var err error
		defer close(result)
		response, err = client.ConferenceToLive(request)
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

// ConferenceToLiveRequest is the request struct for api ConferenceToLive
type ConferenceToLiveRequest struct {
	*requests.RpcRequest
	UserId           string           `position:"Body" name:"UserId"`
	OpenPasswordFlag requests.Boolean `position:"Body" name:"OpenPasswordFlag"`
	MeetingUUID      string           `position:"Body" name:"MeetingUUID"`
	Password         string           `position:"Body" name:"Password"`
	LiveName         string           `position:"Body" name:"LiveName"`
}

// ConferenceToLiveResponse is the response struct for api ConferenceToLive
type ConferenceToLiveResponse struct {
	*responses.BaseResponse
	ErrorCode int      `json:"ErrorCode" xml:"ErrorCode"`
	Message   string   `json:"Message" xml:"Message"`
	Success   bool     `json:"Success" xml:"Success"`
	RequestId string   `json:"RequestId" xml:"RequestId"`
	LiveInfo  LiveInfo `json:"LiveInfo" xml:"LiveInfo"`
}

// CreateConferenceToLiveRequest creates a request to invoke ConferenceToLive API
func CreateConferenceToLiveRequest() (request *ConferenceToLiveRequest) {
	request = &ConferenceToLiveRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "ConferenceToLive", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConferenceToLiveResponse creates a response to parse from ConferenceToLive response
func CreateConferenceToLiveResponse() (response *ConferenceToLiveResponse) {
	response = &ConferenceToLiveResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
