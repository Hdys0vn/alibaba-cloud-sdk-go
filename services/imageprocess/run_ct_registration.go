package imageprocess

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

// RunCTRegistration invokes the imageprocess.RunCTRegistration API synchronously
func (client *Client) RunCTRegistration(request *RunCTRegistrationRequest) (response *RunCTRegistrationResponse, err error) {
	response = CreateRunCTRegistrationResponse()
	err = client.DoAction(request, response)
	return
}

// RunCTRegistrationWithChan invokes the imageprocess.RunCTRegistration API asynchronously
func (client *Client) RunCTRegistrationWithChan(request *RunCTRegistrationRequest) (<-chan *RunCTRegistrationResponse, <-chan error) {
	responseChan := make(chan *RunCTRegistrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RunCTRegistration(request)
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

// RunCTRegistrationWithCallback invokes the imageprocess.RunCTRegistration API asynchronously
func (client *Client) RunCTRegistrationWithCallback(request *RunCTRegistrationRequest, callback func(response *RunCTRegistrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RunCTRegistrationResponse
		var err error
		defer close(result)
		response, err = client.RunCTRegistration(request)
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

// RunCTRegistrationRequest is the request struct for api RunCTRegistration
type RunCTRegistrationRequest struct {
	*requests.RpcRequest
	DataSourceType string                            `position:"Body" name:"DataSourceType"`
	OrgName        string                            `position:"Body" name:"OrgName"`
	ReferenceList  *[]RunCTRegistrationReferenceList `position:"Body" name:"ReferenceList"  type:"Repeated"`
	DataFormat     string                            `position:"Body" name:"DataFormat"`
	OrgId          string                            `position:"Body" name:"OrgId"`
	Async          requests.Boolean                  `position:"Body" name:"Async"`
	FloatingList   *[]RunCTRegistrationFloatingList  `position:"Body" name:"FloatingList"  type:"Repeated"`
}

// RunCTRegistrationReferenceList is a repeated param struct in RunCTRegistrationRequest
type RunCTRegistrationReferenceList struct {
	ReferenceURL string `name:"ReferenceURL"`
}

// RunCTRegistrationFloatingList is a repeated param struct in RunCTRegistrationRequest
type RunCTRegistrationFloatingList struct {
	FloatingURL string `name:"FloatingURL"`
}

// RunCTRegistrationResponse is the response struct for api RunCTRegistration
type RunCTRegistrationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRunCTRegistrationRequest creates a request to invoke RunCTRegistration API
func CreateRunCTRegistrationRequest() (request *RunCTRegistrationRequest) {
	request = &RunCTRegistrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageprocess", "2020-03-20", "RunCTRegistration", "imageprocess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRunCTRegistrationResponse creates a response to parse from RunCTRegistration response
func CreateRunCTRegistrationResponse() (response *RunCTRegistrationResponse) {
	response = &RunCTRegistrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}