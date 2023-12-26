package eflo

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

// AssignPrivateIpAddress invokes the eflo.AssignPrivateIpAddress API synchronously
func (client *Client) AssignPrivateIpAddress(request *AssignPrivateIpAddressRequest) (response *AssignPrivateIpAddressResponse, err error) {
	response = CreateAssignPrivateIpAddressResponse()
	err = client.DoAction(request, response)
	return
}

// AssignPrivateIpAddressWithChan invokes the eflo.AssignPrivateIpAddress API asynchronously
func (client *Client) AssignPrivateIpAddressWithChan(request *AssignPrivateIpAddressRequest) (<-chan *AssignPrivateIpAddressResponse, <-chan error) {
	responseChan := make(chan *AssignPrivateIpAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssignPrivateIpAddress(request)
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

// AssignPrivateIpAddressWithCallback invokes the eflo.AssignPrivateIpAddress API asynchronously
func (client *Client) AssignPrivateIpAddressWithCallback(request *AssignPrivateIpAddressRequest, callback func(response *AssignPrivateIpAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssignPrivateIpAddressResponse
		var err error
		defer close(result)
		response, err = client.AssignPrivateIpAddress(request)
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

// AssignPrivateIpAddressRequest is the request struct for api AssignPrivateIpAddress
type AssignPrivateIpAddressRequest struct {
	*requests.RpcRequest
	ClientToken        string           `position:"Body" name:"ClientToken"`
	Description        string           `position:"Body" name:"Description"`
	AssignMac          requests.Boolean `position:"Body" name:"AssignMac"`
	SubnetId           string           `position:"Body" name:"SubnetId"`
	SkipConfig         requests.Boolean `position:"Body" name:"SkipConfig"`
	PrivateIpAddress   string           `position:"Body" name:"PrivateIpAddress"`
	NetworkInterfaceId string           `position:"Body" name:"NetworkInterfaceId"`
}

// AssignPrivateIpAddressResponse is the response struct for api AssignPrivateIpAddress
type AssignPrivateIpAddressResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateAssignPrivateIpAddressRequest creates a request to invoke AssignPrivateIpAddress API
func CreateAssignPrivateIpAddressRequest() (request *AssignPrivateIpAddressRequest) {
	request = &AssignPrivateIpAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "AssignPrivateIpAddress", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssignPrivateIpAddressResponse creates a response to parse from AssignPrivateIpAddress response
func CreateAssignPrivateIpAddressResponse() (response *AssignPrivateIpAddressResponse) {
	response = &AssignPrivateIpAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
