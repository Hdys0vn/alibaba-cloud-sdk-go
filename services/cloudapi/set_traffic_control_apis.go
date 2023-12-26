package cloudapi

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

// SetTrafficControlApis invokes the cloudapi.SetTrafficControlApis API synchronously
func (client *Client) SetTrafficControlApis(request *SetTrafficControlApisRequest) (response *SetTrafficControlApisResponse, err error) {
	response = CreateSetTrafficControlApisResponse()
	err = client.DoAction(request, response)
	return
}

// SetTrafficControlApisWithChan invokes the cloudapi.SetTrafficControlApis API asynchronously
func (client *Client) SetTrafficControlApisWithChan(request *SetTrafficControlApisRequest) (<-chan *SetTrafficControlApisResponse, <-chan error) {
	responseChan := make(chan *SetTrafficControlApisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetTrafficControlApis(request)
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

// SetTrafficControlApisWithCallback invokes the cloudapi.SetTrafficControlApis API asynchronously
func (client *Client) SetTrafficControlApisWithCallback(request *SetTrafficControlApisRequest, callback func(response *SetTrafficControlApisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetTrafficControlApisResponse
		var err error
		defer close(result)
		response, err = client.SetTrafficControlApis(request)
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

// SetTrafficControlApisRequest is the request struct for api SetTrafficControlApis
type SetTrafficControlApisRequest struct {
	*requests.RpcRequest
	TrafficControlId string `position:"Query" name:"TrafficControlId"`
	StageName        string `position:"Query" name:"StageName"`
	GroupId          string `position:"Query" name:"GroupId"`
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	ApiIds           string `position:"Query" name:"ApiIds"`
}

// SetTrafficControlApisResponse is the response struct for api SetTrafficControlApis
type SetTrafficControlApisResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetTrafficControlApisRequest creates a request to invoke SetTrafficControlApis API
func CreateSetTrafficControlApisRequest() (request *SetTrafficControlApisRequest) {
	request = &SetTrafficControlApisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "SetTrafficControlApis", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetTrafficControlApisResponse creates a response to parse from SetTrafficControlApis response
func CreateSetTrafficControlApisResponse() (response *SetTrafficControlApisResponse) {
	response = &SetTrafficControlApisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
