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

// CreateMeetingInternational invokes the aliyuncvc.CreateMeetingInternational API synchronously
func (client *Client) CreateMeetingInternational(request *CreateMeetingInternationalRequest) (response *CreateMeetingInternationalResponse, err error) {
	response = CreateCreateMeetingInternationalResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMeetingInternationalWithChan invokes the aliyuncvc.CreateMeetingInternational API asynchronously
func (client *Client) CreateMeetingInternationalWithChan(request *CreateMeetingInternationalRequest) (<-chan *CreateMeetingInternationalResponse, <-chan error) {
	responseChan := make(chan *CreateMeetingInternationalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMeetingInternational(request)
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

// CreateMeetingInternationalWithCallback invokes the aliyuncvc.CreateMeetingInternational API asynchronously
func (client *Client) CreateMeetingInternationalWithCallback(request *CreateMeetingInternationalRequest, callback func(response *CreateMeetingInternationalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMeetingInternationalResponse
		var err error
		defer close(result)
		response, err = client.CreateMeetingInternational(request)
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

// CreateMeetingInternationalRequest is the request struct for api CreateMeetingInternational
type CreateMeetingInternationalRequest struct {
	*requests.RpcRequest
	MeetingName      string `position:"Body" name:"MeetingName"`
	UserId           string `position:"Body" name:"UserId"`
	OpenPasswordFlag string `position:"Body" name:"OpenPasswordFlag"`
	Password         string `position:"Body" name:"Password"`
}

// CreateMeetingInternationalResponse is the response struct for api CreateMeetingInternational
type CreateMeetingInternationalResponse struct {
	*responses.BaseResponse
	ErrorCode   int         `json:"ErrorCode" xml:"ErrorCode"`
	Message     string      `json:"Message" xml:"Message"`
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	MeetingInfo MeetingInfo `json:"MeetingInfo" xml:"MeetingInfo"`
}

// CreateCreateMeetingInternationalRequest creates a request to invoke CreateMeetingInternational API
func CreateCreateMeetingInternationalRequest() (request *CreateMeetingInternationalRequest) {
	request = &CreateMeetingInternationalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "CreateMeetingInternational", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMeetingInternationalResponse creates a response to parse from CreateMeetingInternational response
func CreateCreateMeetingInternationalResponse() (response *CreateMeetingInternationalResponse) {
	response = &CreateMeetingInternationalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
