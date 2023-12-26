package ga

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

// UpdateIpSet invokes the ga.UpdateIpSet API synchronously
func (client *Client) UpdateIpSet(request *UpdateIpSetRequest) (response *UpdateIpSetResponse, err error) {
	response = CreateUpdateIpSetResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateIpSetWithChan invokes the ga.UpdateIpSet API asynchronously
func (client *Client) UpdateIpSetWithChan(request *UpdateIpSetRequest) (<-chan *UpdateIpSetResponse, <-chan error) {
	responseChan := make(chan *UpdateIpSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateIpSet(request)
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

// UpdateIpSetWithCallback invokes the ga.UpdateIpSet API asynchronously
func (client *Client) UpdateIpSetWithCallback(request *UpdateIpSetRequest, callback func(response *UpdateIpSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateIpSetResponse
		var err error
		defer close(result)
		response, err = client.UpdateIpSet(request)
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

// UpdateIpSetRequest is the request struct for api UpdateIpSet
type UpdateIpSetRequest struct {
	*requests.RpcRequest
	ClientToken string           `position:"Query" name:"ClientToken"`
	Bandwidth   requests.Integer `position:"Query" name:"Bandwidth"`
	IpSetId     string           `position:"Query" name:"IpSetId"`
}

// UpdateIpSetResponse is the response struct for api UpdateIpSet
type UpdateIpSetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateIpSetRequest creates a request to invoke UpdateIpSet API
func CreateUpdateIpSetRequest() (request *UpdateIpSetRequest) {
	request = &UpdateIpSetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ga", "2019-11-20", "UpdateIpSet", "gaplus", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateIpSetResponse creates a response to parse from UpdateIpSet response
func CreateUpdateIpSetResponse() (response *UpdateIpSetResponse) {
	response = &UpdateIpSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
