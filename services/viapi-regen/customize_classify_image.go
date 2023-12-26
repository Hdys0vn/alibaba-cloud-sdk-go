package viapi_regen

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

// CustomizeClassifyImage invokes the viapi_regen.CustomizeClassifyImage API synchronously
func (client *Client) CustomizeClassifyImage(request *CustomizeClassifyImageRequest) (response *CustomizeClassifyImageResponse, err error) {
	response = CreateCustomizeClassifyImageResponse()
	err = client.DoAction(request, response)
	return
}

// CustomizeClassifyImageWithChan invokes the viapi_regen.CustomizeClassifyImage API asynchronously
func (client *Client) CustomizeClassifyImageWithChan(request *CustomizeClassifyImageRequest) (<-chan *CustomizeClassifyImageResponse, <-chan error) {
	responseChan := make(chan *CustomizeClassifyImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CustomizeClassifyImage(request)
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

// CustomizeClassifyImageWithCallback invokes the viapi_regen.CustomizeClassifyImage API asynchronously
func (client *Client) CustomizeClassifyImageWithCallback(request *CustomizeClassifyImageRequest, callback func(response *CustomizeClassifyImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CustomizeClassifyImageResponse
		var err error
		defer close(result)
		response, err = client.CustomizeClassifyImage(request)
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

// CustomizeClassifyImageRequest is the request struct for api CustomizeClassifyImage
type CustomizeClassifyImageRequest struct {
	*requests.RpcRequest
	ImageUrl  string `position:"Body" name:"ImageUrl"`
	ServiceId string `position:"Body" name:"ServiceId"`
}

// CustomizeClassifyImageResponse is the response struct for api CustomizeClassifyImage
type CustomizeClassifyImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCustomizeClassifyImageRequest creates a request to invoke CustomizeClassifyImage API
func CreateCustomizeClassifyImageRequest() (request *CustomizeClassifyImageRequest) {
	request = &CustomizeClassifyImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("viapi-regen", "2021-11-19", "CustomizeClassifyImage", "selflearning", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCustomizeClassifyImageResponse creates a response to parse from CustomizeClassifyImage response
func CreateCustomizeClassifyImageResponse() (response *CustomizeClassifyImageResponse) {
	response = &CustomizeClassifyImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
