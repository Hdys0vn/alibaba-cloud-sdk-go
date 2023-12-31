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

// ModifyMeetingPasswordInternational invokes the aliyuncvc.ModifyMeetingPasswordInternational API synchronously
func (client *Client) ModifyMeetingPasswordInternational(request *ModifyMeetingPasswordInternationalRequest) (response *ModifyMeetingPasswordInternationalResponse, err error) {
	response = CreateModifyMeetingPasswordInternationalResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMeetingPasswordInternationalWithChan invokes the aliyuncvc.ModifyMeetingPasswordInternational API asynchronously
func (client *Client) ModifyMeetingPasswordInternationalWithChan(request *ModifyMeetingPasswordInternationalRequest) (<-chan *ModifyMeetingPasswordInternationalResponse, <-chan error) {
	responseChan := make(chan *ModifyMeetingPasswordInternationalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMeetingPasswordInternational(request)
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

// ModifyMeetingPasswordInternationalWithCallback invokes the aliyuncvc.ModifyMeetingPasswordInternational API asynchronously
func (client *Client) ModifyMeetingPasswordInternationalWithCallback(request *ModifyMeetingPasswordInternationalRequest, callback func(response *ModifyMeetingPasswordInternationalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMeetingPasswordInternationalResponse
		var err error
		defer close(result)
		response, err = client.ModifyMeetingPasswordInternational(request)
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

// ModifyMeetingPasswordInternationalRequest is the request struct for api ModifyMeetingPasswordInternational
type ModifyMeetingPasswordInternationalRequest struct {
	*requests.RpcRequest
	UserId           string           `position:"Body" name:"UserId"`
	OpenPasswordFlag requests.Boolean `position:"Body" name:"OpenPasswordFlag"`
	MeetingUUID      string           `position:"Body" name:"MeetingUUID"`
	Password         string           `position:"Body" name:"Password"`
}

// ModifyMeetingPasswordInternationalResponse is the response struct for api ModifyMeetingPasswordInternational
type ModifyMeetingPasswordInternationalResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyMeetingPasswordInternationalRequest creates a request to invoke ModifyMeetingPasswordInternational API
func CreateModifyMeetingPasswordInternationalRequest() (request *ModifyMeetingPasswordInternationalRequest) {
	request = &ModifyMeetingPasswordInternationalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "ModifyMeetingPasswordInternational", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyMeetingPasswordInternationalResponse creates a response to parse from ModifyMeetingPasswordInternational response
func CreateModifyMeetingPasswordInternationalResponse() (response *ModifyMeetingPasswordInternationalResponse) {
	response = &ModifyMeetingPasswordInternationalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
