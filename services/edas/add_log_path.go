package edas

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

// AddLogPath invokes the edas.AddLogPath API synchronously
func (client *Client) AddLogPath(request *AddLogPathRequest) (response *AddLogPathResponse, err error) {
	response = CreateAddLogPathResponse()
	err = client.DoAction(request, response)
	return
}

// AddLogPathWithChan invokes the edas.AddLogPath API asynchronously
func (client *Client) AddLogPathWithChan(request *AddLogPathRequest) (<-chan *AddLogPathResponse, <-chan error) {
	responseChan := make(chan *AddLogPathResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLogPath(request)
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

// AddLogPathWithCallback invokes the edas.AddLogPath API asynchronously
func (client *Client) AddLogPathWithCallback(request *AddLogPathRequest, callback func(response *AddLogPathResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLogPathResponse
		var err error
		defer close(result)
		response, err = client.AddLogPath(request)
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

// AddLogPathRequest is the request struct for api AddLogPath
type AddLogPathRequest struct {
	*requests.RoaRequest
	Path  string `position:"Body" name:"Path"`
	AppId string `position:"Body" name:"AppId"`
}

// AddLogPathResponse is the response struct for api AddLogPath
type AddLogPathResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLogPathRequest creates a request to invoke AddLogPath API
func CreateAddLogPathRequest() (request *AddLogPathRequest) {
	request = &AddLogPathRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "AddLogPath", "/pop/v5/log/popListLogDirs", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddLogPathResponse creates a response to parse from AddLogPath response
func CreateAddLogPathResponse() (response *AddLogPathResponse) {
	response = &AddLogPathResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
