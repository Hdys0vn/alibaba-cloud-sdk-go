package faas

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

// UpdateImageAttribute invokes the faas.UpdateImageAttribute API synchronously
// api document: https://help.aliyun.com/api/faas/updateimageattribute.html
func (client *Client) UpdateImageAttribute(request *UpdateImageAttributeRequest) (response *UpdateImageAttributeResponse, err error) {
	response = CreateUpdateImageAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateImageAttributeWithChan invokes the faas.UpdateImageAttribute API asynchronously
// api document: https://help.aliyun.com/api/faas/updateimageattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateImageAttributeWithChan(request *UpdateImageAttributeRequest) (<-chan *UpdateImageAttributeResponse, <-chan error) {
	responseChan := make(chan *UpdateImageAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateImageAttribute(request)
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

// UpdateImageAttributeWithCallback invokes the faas.UpdateImageAttribute API asynchronously
// api document: https://help.aliyun.com/api/faas/updateimageattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateImageAttributeWithCallback(request *UpdateImageAttributeRequest, callback func(response *UpdateImageAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateImageAttributeResponse
		var err error
		defer close(result)
		response, err = client.UpdateImageAttribute(request)
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

// UpdateImageAttributeRequest is the request struct for api UpdateImageAttribute
type UpdateImageAttributeRequest struct {
	*requests.RpcRequest
	Name              string `position:"Query" name:"Name"`
	Description       string `position:"Query" name:"Description"`
	FpgaImageUniqueId string `position:"Query" name:"FpgaImageUniqueId"`
	Tags              string `position:"Query" name:"Tags"`
}

// UpdateImageAttributeResponse is the response struct for api UpdateImageAttribute
type UpdateImageAttributeResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	Name              string `json:"Name" xml:"Name"`
	Description       string `json:"Description" xml:"Description"`
	Tags              string `json:"Tags" xml:"Tags"`
	FpgaImageUniqueId string `json:"FpgaImageUniqueId" xml:"FpgaImageUniqueId"`
}

// CreateUpdateImageAttributeRequest creates a request to invoke UpdateImageAttribute API
func CreateUpdateImageAttributeRequest() (request *UpdateImageAttributeRequest) {
	request = &UpdateImageAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("faas", "2020-02-17", "UpdateImageAttribute", "faas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateImageAttributeResponse creates a response to parse from UpdateImageAttribute response
func CreateUpdateImageAttributeResponse() (response *UpdateImageAttributeResponse) {
	response = &UpdateImageAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
