package openanalytics_open

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

// SetAllowIP invokes the openanalytics_open.SetAllowIP API synchronously
func (client *Client) SetAllowIP(request *SetAllowIPRequest) (response *SetAllowIPResponse, err error) {
	response = CreateSetAllowIPResponse()
	err = client.DoAction(request, response)
	return
}

// SetAllowIPWithChan invokes the openanalytics_open.SetAllowIP API asynchronously
func (client *Client) SetAllowIPWithChan(request *SetAllowIPRequest) (<-chan *SetAllowIPResponse, <-chan error) {
	responseChan := make(chan *SetAllowIPResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetAllowIP(request)
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

// SetAllowIPWithCallback invokes the openanalytics_open.SetAllowIP API asynchronously
func (client *Client) SetAllowIPWithCallback(request *SetAllowIPRequest, callback func(response *SetAllowIPResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetAllowIPResponse
		var err error
		defer close(result)
		response, err = client.SetAllowIP(request)
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

// SetAllowIPRequest is the request struct for api SetAllowIP
type SetAllowIPRequest struct {
	*requests.RpcRequest
	NetworkType string           `position:"Body" name:"NetworkType"`
	Product     string           `position:"Body" name:"Product"`
	AllowIP     string           `position:"Body" name:"AllowIP"`
	Append      requests.Boolean `position:"Body" name:"Append"`
}

// SetAllowIPResponse is the response struct for api SetAllowIP
type SetAllowIPResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	RegionId  string `json:"RegionId" xml:"RegionId"`
}

// CreateSetAllowIPRequest creates a request to invoke SetAllowIP API
func CreateSetAllowIPRequest() (request *SetAllowIPRequest) {
	request = &SetAllowIPRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "SetAllowIP", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetAllowIPResponse creates a response to parse from SetAllowIP response
func CreateSetAllowIPResponse() (response *SetAllowIPResponse) {
	response = &SetAllowIPResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
