package cloud_siem

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

// PostCustomizeRuleTest invokes the cloud_siem.PostCustomizeRuleTest API synchronously
func (client *Client) PostCustomizeRuleTest(request *PostCustomizeRuleTestRequest) (response *PostCustomizeRuleTestResponse, err error) {
	response = CreatePostCustomizeRuleTestResponse()
	err = client.DoAction(request, response)
	return
}

// PostCustomizeRuleTestWithChan invokes the cloud_siem.PostCustomizeRuleTest API asynchronously
func (client *Client) PostCustomizeRuleTestWithChan(request *PostCustomizeRuleTestRequest) (<-chan *PostCustomizeRuleTestResponse, <-chan error) {
	responseChan := make(chan *PostCustomizeRuleTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PostCustomizeRuleTest(request)
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

// PostCustomizeRuleTestWithCallback invokes the cloud_siem.PostCustomizeRuleTest API asynchronously
func (client *Client) PostCustomizeRuleTestWithCallback(request *PostCustomizeRuleTestRequest, callback func(response *PostCustomizeRuleTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PostCustomizeRuleTestResponse
		var err error
		defer close(result)
		response, err = client.PostCustomizeRuleTest(request)
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

// PostCustomizeRuleTestRequest is the request struct for api PostCustomizeRuleTest
type PostCustomizeRuleTestRequest struct {
	*requests.RpcRequest
	Id            requests.Integer `position:"Body" name:"Id"`
	SimulatedData string           `position:"Body" name:"SimulatedData"`
	TestType      string           `position:"Body" name:"TestType"`
}

// PostCustomizeRuleTestResponse is the response struct for api PostCustomizeRuleTest
type PostCustomizeRuleTestResponse struct {
	*responses.BaseResponse
}

// CreatePostCustomizeRuleTestRequest creates a request to invoke PostCustomizeRuleTest API
func CreatePostCustomizeRuleTestRequest() (request *PostCustomizeRuleTestRequest) {
	request = &PostCustomizeRuleTestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "PostCustomizeRuleTest", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePostCustomizeRuleTestResponse creates a response to parse from PostCustomizeRuleTest response
func CreatePostCustomizeRuleTestResponse() (response *PostCustomizeRuleTestResponse) {
	response = &PostCustomizeRuleTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}