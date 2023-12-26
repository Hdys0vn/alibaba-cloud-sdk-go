package live

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

// GetCustomTemplate invokes the live.GetCustomTemplate API synchronously
func (client *Client) GetCustomTemplate(request *GetCustomTemplateRequest) (response *GetCustomTemplateResponse, err error) {
	response = CreateGetCustomTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetCustomTemplateWithChan invokes the live.GetCustomTemplate API asynchronously
func (client *Client) GetCustomTemplateWithChan(request *GetCustomTemplateRequest) (<-chan *GetCustomTemplateResponse, <-chan error) {
	responseChan := make(chan *GetCustomTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCustomTemplate(request)
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

// GetCustomTemplateWithCallback invokes the live.GetCustomTemplate API asynchronously
func (client *Client) GetCustomTemplateWithCallback(request *GetCustomTemplateRequest, callback func(response *GetCustomTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCustomTemplateResponse
		var err error
		defer close(result)
		response, err = client.GetCustomTemplate(request)
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

// GetCustomTemplateRequest is the request struct for api GetCustomTemplate
type GetCustomTemplateRequest struct {
	*requests.RpcRequest
	Template string           `position:"Query" name:"Template"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
}

// GetCustomTemplateResponse is the response struct for api GetCustomTemplate
type GetCustomTemplateResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	CustomTemplate string `json:"CustomTemplate" xml:"CustomTemplate"`
	Template       string `json:"Template" xml:"Template"`
}

// CreateGetCustomTemplateRequest creates a request to invoke GetCustomTemplate API
func CreateGetCustomTemplateRequest() (request *GetCustomTemplateRequest) {
	request = &GetCustomTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "GetCustomTemplate", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetCustomTemplateResponse creates a response to parse from GetCustomTemplate response
func CreateGetCustomTemplateResponse() (response *GetCustomTemplateResponse) {
	response = &GetCustomTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
