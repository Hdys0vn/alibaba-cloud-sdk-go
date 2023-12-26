package sls

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

// GetAlert invokes the sls.GetAlert API synchronously
func (client *Client) GetAlert(request *GetAlertRequest) (response *GetAlertResponse, err error) {
	response = CreateGetAlertResponse()
	err = client.DoAction(request, response)
	return
}

// GetAlertWithChan invokes the sls.GetAlert API asynchronously
func (client *Client) GetAlertWithChan(request *GetAlertRequest) (<-chan *GetAlertResponse, <-chan error) {
	responseChan := make(chan *GetAlertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAlert(request)
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

// GetAlertWithCallback invokes the sls.GetAlert API asynchronously
func (client *Client) GetAlertWithCallback(request *GetAlertRequest, callback func(response *GetAlertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAlertResponse
		var err error
		defer close(result)
		response, err = client.GetAlert(request)
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

// GetAlertRequest is the request struct for api GetAlert
type GetAlertRequest struct {
	*requests.RpcRequest
	App         string `position:"Body" name:"App"`
	ProjectName string `position:"Body" name:"ProjectName"`
	Endpoint    string `position:"Body" name:"Endpoint"`
	AlertId     string `position:"Body" name:"AlertId"`
}

// GetAlertResponse is the response struct for api GetAlert
type GetAlertResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetAlertRequest creates a request to invoke GetAlert API
func CreateGetAlertRequest() (request *GetAlertRequest) {
	request = &GetAlertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2019-10-23", "GetAlert", "sls", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAlertResponse creates a response to parse from GetAlert response
func CreateGetAlertResponse() (response *GetAlertResponse) {
	response = &GetAlertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
