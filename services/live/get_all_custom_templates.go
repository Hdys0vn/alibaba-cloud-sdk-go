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

// GetAllCustomTemplates invokes the live.GetAllCustomTemplates API synchronously
func (client *Client) GetAllCustomTemplates(request *GetAllCustomTemplatesRequest) (response *GetAllCustomTemplatesResponse, err error) {
	response = CreateGetAllCustomTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// GetAllCustomTemplatesWithChan invokes the live.GetAllCustomTemplates API asynchronously
func (client *Client) GetAllCustomTemplatesWithChan(request *GetAllCustomTemplatesRequest) (<-chan *GetAllCustomTemplatesResponse, <-chan error) {
	responseChan := make(chan *GetAllCustomTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAllCustomTemplates(request)
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

// GetAllCustomTemplatesWithCallback invokes the live.GetAllCustomTemplates API asynchronously
func (client *Client) GetAllCustomTemplatesWithCallback(request *GetAllCustomTemplatesRequest, callback func(response *GetAllCustomTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAllCustomTemplatesResponse
		var err error
		defer close(result)
		response, err = client.GetAllCustomTemplates(request)
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

// GetAllCustomTemplatesRequest is the request struct for api GetAllCustomTemplates
type GetAllCustomTemplatesRequest struct {
	*requests.RpcRequest
	UserId  string           `position:"Query" name:"UserId"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// GetAllCustomTemplatesResponse is the response struct for api GetAllCustomTemplates
type GetAllCustomTemplatesResponse struct {
	*responses.BaseResponse
	CustomTemplates string `json:"CustomTemplates" xml:"CustomTemplates"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
}

// CreateGetAllCustomTemplatesRequest creates a request to invoke GetAllCustomTemplates API
func CreateGetAllCustomTemplatesRequest() (request *GetAllCustomTemplatesRequest) {
	request = &GetAllCustomTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "GetAllCustomTemplates", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAllCustomTemplatesResponse creates a response to parse from GetAllCustomTemplates response
func CreateGetAllCustomTemplatesResponse() (response *GetAllCustomTemplatesResponse) {
	response = &GetAllCustomTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
