package sgw

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

// AddTagsToGateway invokes the sgw.AddTagsToGateway API synchronously
func (client *Client) AddTagsToGateway(request *AddTagsToGatewayRequest) (response *AddTagsToGatewayResponse, err error) {
	response = CreateAddTagsToGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// AddTagsToGatewayWithChan invokes the sgw.AddTagsToGateway API asynchronously
func (client *Client) AddTagsToGatewayWithChan(request *AddTagsToGatewayRequest) (<-chan *AddTagsToGatewayResponse, <-chan error) {
	responseChan := make(chan *AddTagsToGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTagsToGateway(request)
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

// AddTagsToGatewayWithCallback invokes the sgw.AddTagsToGateway API asynchronously
func (client *Client) AddTagsToGatewayWithCallback(request *AddTagsToGatewayRequest, callback func(response *AddTagsToGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTagsToGatewayResponse
		var err error
		defer close(result)
		response, err = client.AddTagsToGateway(request)
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

// AddTagsToGatewayRequest is the request struct for api AddTagsToGateway
type AddTagsToGatewayRequest struct {
	*requests.RpcRequest
	Tags          string `position:"Query" name:"Tags"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// AddTagsToGatewayResponse is the response struct for api AddTagsToGateway
type AddTagsToGatewayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateAddTagsToGatewayRequest creates a request to invoke AddTagsToGateway API
func CreateAddTagsToGatewayRequest() (request *AddTagsToGatewayRequest) {
	request = &AddTagsToGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "AddTagsToGateway", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddTagsToGatewayResponse creates a response to parse from AddTagsToGateway response
func CreateAddTagsToGatewayResponse() (response *AddTagsToGatewayResponse) {
	response = &AddTagsToGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}