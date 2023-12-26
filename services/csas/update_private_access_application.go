package csas

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

// UpdatePrivateAccessApplication invokes the csas.UpdatePrivateAccessApplication API synchronously
func (client *Client) UpdatePrivateAccessApplication(request *UpdatePrivateAccessApplicationRequest) (response *UpdatePrivateAccessApplicationResponse, err error) {
	response = CreateUpdatePrivateAccessApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePrivateAccessApplicationWithChan invokes the csas.UpdatePrivateAccessApplication API asynchronously
func (client *Client) UpdatePrivateAccessApplicationWithChan(request *UpdatePrivateAccessApplicationRequest) (<-chan *UpdatePrivateAccessApplicationResponse, <-chan error) {
	responseChan := make(chan *UpdatePrivateAccessApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePrivateAccessApplication(request)
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

// UpdatePrivateAccessApplicationWithCallback invokes the csas.UpdatePrivateAccessApplication API asynchronously
func (client *Client) UpdatePrivateAccessApplicationWithCallback(request *UpdatePrivateAccessApplicationRequest, callback func(response *UpdatePrivateAccessApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePrivateAccessApplicationResponse
		var err error
		defer close(result)
		response, err = client.UpdatePrivateAccessApplication(request)
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

// UpdatePrivateAccessApplicationRequest is the request struct for api UpdatePrivateAccessApplication
type UpdatePrivateAccessApplicationRequest struct {
	*requests.RpcRequest
	Addresses     *[]string                                   `position:"Body" name:"Addresses"  type:"Repeated"`
	Description   string                                      `position:"Body" name:"Description"`
	Protocol      string                                      `position:"Body" name:"Protocol"`
	SourceIp      string                                      `position:"Query" name:"SourceIp"`
	ApplicationId string                                      `position:"Body" name:"ApplicationId"`
	TagIds        *[]string                                   `position:"Body" name:"TagIds"  type:"Repeated"`
	PortRanges    *[]UpdatePrivateAccessApplicationPortRanges `position:"Body" name:"PortRanges"  type:"Repeated"`
	ModifyType    string                                      `position:"Body" name:"ModifyType"`
	Status        string                                      `position:"Body" name:"Status"`
}

// UpdatePrivateAccessApplicationPortRanges is a repeated param struct in UpdatePrivateAccessApplicationRequest
type UpdatePrivateAccessApplicationPortRanges struct {
	End   string `name:"End"`
	Begin string `name:"Begin"`
}

// UpdatePrivateAccessApplicationResponse is the response struct for api UpdatePrivateAccessApplication
type UpdatePrivateAccessApplicationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdatePrivateAccessApplicationRequest creates a request to invoke UpdatePrivateAccessApplication API
func CreateUpdatePrivateAccessApplicationRequest() (request *UpdatePrivateAccessApplicationRequest) {
	request = &UpdatePrivateAccessApplicationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "UpdatePrivateAccessApplication", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdatePrivateAccessApplicationResponse creates a response to parse from UpdatePrivateAccessApplication response
func CreateUpdatePrivateAccessApplicationResponse() (response *UpdatePrivateAccessApplicationResponse) {
	response = &UpdatePrivateAccessApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
