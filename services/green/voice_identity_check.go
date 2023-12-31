package green

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

// VoiceIdentityCheck invokes the green.VoiceIdentityCheck API synchronously
func (client *Client) VoiceIdentityCheck(request *VoiceIdentityCheckRequest) (response *VoiceIdentityCheckResponse, err error) {
	response = CreateVoiceIdentityCheckResponse()
	err = client.DoAction(request, response)
	return
}

// VoiceIdentityCheckWithChan invokes the green.VoiceIdentityCheck API asynchronously
func (client *Client) VoiceIdentityCheckWithChan(request *VoiceIdentityCheckRequest) (<-chan *VoiceIdentityCheckResponse, <-chan error) {
	responseChan := make(chan *VoiceIdentityCheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VoiceIdentityCheck(request)
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

// VoiceIdentityCheckWithCallback invokes the green.VoiceIdentityCheck API asynchronously
func (client *Client) VoiceIdentityCheckWithCallback(request *VoiceIdentityCheckRequest, callback func(response *VoiceIdentityCheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VoiceIdentityCheckResponse
		var err error
		defer close(result)
		response, err = client.VoiceIdentityCheck(request)
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

// VoiceIdentityCheckRequest is the request struct for api VoiceIdentityCheck
type VoiceIdentityCheckRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// VoiceIdentityCheckResponse is the response struct for api VoiceIdentityCheck
type VoiceIdentityCheckResponse struct {
	*responses.BaseResponse
}

// CreateVoiceIdentityCheckRequest creates a request to invoke VoiceIdentityCheck API
func CreateVoiceIdentityCheckRequest() (request *VoiceIdentityCheckRequest) {
	request = &VoiceIdentityCheckRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "VoiceIdentityCheck", "/green/voice/auth/check", "", "")
	request.Method = requests.POST
	return
}

// CreateVoiceIdentityCheckResponse creates a response to parse from VoiceIdentityCheck response
func CreateVoiceIdentityCheckResponse() (response *VoiceIdentityCheckResponse) {
	response = &VoiceIdentityCheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
