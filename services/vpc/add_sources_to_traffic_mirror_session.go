package vpc

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

// AddSourcesToTrafficMirrorSession invokes the vpc.AddSourcesToTrafficMirrorSession API synchronously
func (client *Client) AddSourcesToTrafficMirrorSession(request *AddSourcesToTrafficMirrorSessionRequest) (response *AddSourcesToTrafficMirrorSessionResponse, err error) {
	response = CreateAddSourcesToTrafficMirrorSessionResponse()
	err = client.DoAction(request, response)
	return
}

// AddSourcesToTrafficMirrorSessionWithChan invokes the vpc.AddSourcesToTrafficMirrorSession API asynchronously
func (client *Client) AddSourcesToTrafficMirrorSessionWithChan(request *AddSourcesToTrafficMirrorSessionRequest) (<-chan *AddSourcesToTrafficMirrorSessionResponse, <-chan error) {
	responseChan := make(chan *AddSourcesToTrafficMirrorSessionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddSourcesToTrafficMirrorSession(request)
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

// AddSourcesToTrafficMirrorSessionWithCallback invokes the vpc.AddSourcesToTrafficMirrorSession API asynchronously
func (client *Client) AddSourcesToTrafficMirrorSessionWithCallback(request *AddSourcesToTrafficMirrorSessionRequest, callback func(response *AddSourcesToTrafficMirrorSessionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddSourcesToTrafficMirrorSessionResponse
		var err error
		defer close(result)
		response, err = client.AddSourcesToTrafficMirrorSession(request)
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

// AddSourcesToTrafficMirrorSessionRequest is the request struct for api AddSourcesToTrafficMirrorSession
type AddSourcesToTrafficMirrorSessionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken            string           `position:"Query" name:"ClientToken"`
	TrafficMirrorSourceIds *[]string        `position:"Query" name:"TrafficMirrorSourceIds"  type:"Repeated"`
	DryRun                 requests.Boolean `position:"Query" name:"DryRun"`
	TrafficMirrorSessionId string           `position:"Query" name:"TrafficMirrorSessionId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
}

// AddSourcesToTrafficMirrorSessionResponse is the response struct for api AddSourcesToTrafficMirrorSession
type AddSourcesToTrafficMirrorSessionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddSourcesToTrafficMirrorSessionRequest creates a request to invoke AddSourcesToTrafficMirrorSession API
func CreateAddSourcesToTrafficMirrorSessionRequest() (request *AddSourcesToTrafficMirrorSessionRequest) {
	request = &AddSourcesToTrafficMirrorSessionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "AddSourcesToTrafficMirrorSession", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddSourcesToTrafficMirrorSessionResponse creates a response to parse from AddSourcesToTrafficMirrorSession response
func CreateAddSourcesToTrafficMirrorSessionResponse() (response *AddSourcesToTrafficMirrorSessionResponse) {
	response = &AddSourcesToTrafficMirrorSessionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
