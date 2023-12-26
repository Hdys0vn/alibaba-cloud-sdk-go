package nas

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

// UpdateRecycleBinAttribute invokes the nas.UpdateRecycleBinAttribute API synchronously
func (client *Client) UpdateRecycleBinAttribute(request *UpdateRecycleBinAttributeRequest) (response *UpdateRecycleBinAttributeResponse, err error) {
	response = CreateUpdateRecycleBinAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRecycleBinAttributeWithChan invokes the nas.UpdateRecycleBinAttribute API asynchronously
func (client *Client) UpdateRecycleBinAttributeWithChan(request *UpdateRecycleBinAttributeRequest) (<-chan *UpdateRecycleBinAttributeResponse, <-chan error) {
	responseChan := make(chan *UpdateRecycleBinAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRecycleBinAttribute(request)
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

// UpdateRecycleBinAttributeWithCallback invokes the nas.UpdateRecycleBinAttribute API asynchronously
func (client *Client) UpdateRecycleBinAttributeWithCallback(request *UpdateRecycleBinAttributeRequest, callback func(response *UpdateRecycleBinAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRecycleBinAttributeResponse
		var err error
		defer close(result)
		response, err = client.UpdateRecycleBinAttribute(request)
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

// UpdateRecycleBinAttributeRequest is the request struct for api UpdateRecycleBinAttribute
type UpdateRecycleBinAttributeRequest struct {
	*requests.RpcRequest
	FileSystemId string           `position:"Query" name:"FileSystemId"`
	ReservedDays requests.Integer `position:"Query" name:"ReservedDays"`
}

// UpdateRecycleBinAttributeResponse is the response struct for api UpdateRecycleBinAttribute
type UpdateRecycleBinAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateRecycleBinAttributeRequest creates a request to invoke UpdateRecycleBinAttribute API
func CreateUpdateRecycleBinAttributeRequest() (request *UpdateRecycleBinAttributeRequest) {
	request = &UpdateRecycleBinAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "UpdateRecycleBinAttribute", "nas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateUpdateRecycleBinAttributeResponse creates a response to parse from UpdateRecycleBinAttribute response
func CreateUpdateRecycleBinAttributeResponse() (response *UpdateRecycleBinAttributeResponse) {
	response = &UpdateRecycleBinAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
