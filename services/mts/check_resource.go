package mts

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

// CheckResource invokes the mts.CheckResource API synchronously
func (client *Client) CheckResource(request *CheckResourceRequest) (response *CheckResourceResponse, err error) {
	response = CreateCheckResourceResponse()
	err = client.DoAction(request, response)
	return
}

// CheckResourceWithChan invokes the mts.CheckResource API asynchronously
func (client *Client) CheckResourceWithChan(request *CheckResourceRequest) (<-chan *CheckResourceResponse, <-chan error) {
	responseChan := make(chan *CheckResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckResource(request)
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

// CheckResourceWithCallback invokes the mts.CheckResource API asynchronously
func (client *Client) CheckResourceWithCallback(request *CheckResourceRequest, callback func(response *CheckResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckResourceResponse
		var err error
		defer close(result)
		response, err = client.CheckResource(request)
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

// CheckResourceRequest is the request struct for api CheckResource
type CheckResourceRequest struct {
	*requests.RpcRequest
	Country        string           `position:"Query" name:"Country"`
	Hid            requests.Integer `position:"Query" name:"Hid"`
	Level          requests.Integer `position:"Query" name:"Level"`
	Invoker        string           `position:"Query" name:"Invoker"`
	Message        string           `position:"Query" name:"Message"`
	Url            string           `position:"Query" name:"Url"`
	Success        requests.Boolean `position:"Query" name:"Success"`
	Interrupt      requests.Boolean `position:"Query" name:"Interrupt"`
	GmtWakeup      string           `position:"Query" name:"GmtWakeup"`
	Pk             string           `position:"Query" name:"Pk"`
	Bid            string           `position:"Query" name:"Bid"`
	Prompt         string           `position:"Query" name:"Prompt"`
	TaskExtraData  string           `position:"Query" name:"TaskExtraData"`
	TaskIdentifier string           `position:"Query" name:"TaskIdentifier"`
}

// CheckResourceResponse is the response struct for api CheckResource
type CheckResourceResponse struct {
	*responses.BaseResponse
	GmtWakeup      string `json:"GmtWakeup" xml:"GmtWakeup"`
	Hid            int64  `json:"Hid" xml:"Hid"`
	Message        string `json:"Message" xml:"Message"`
	TaskIdentifier string `json:"TaskIdentifier" xml:"TaskIdentifier"`
	Success        bool   `json:"Success" xml:"Success"`
	Url            string `json:"Url" xml:"Url"`
	Interrupt      bool   `json:"Interrupt" xml:"Interrupt"`
	Invoker        string `json:"Invoker" xml:"Invoker"`
	TaskExtraData  string `json:"TaskExtraData" xml:"TaskExtraData"`
	Country        string `json:"Country" xml:"Country"`
	Prompt         string `json:"Prompt" xml:"Prompt"`
	Level          int64  `json:"Level" xml:"Level"`
	Pk             string `json:"Pk" xml:"Pk"`
	Bid            string `json:"Bid" xml:"Bid"`
}

// CreateCheckResourceRequest creates a request to invoke CheckResource API
func CreateCheckResourceRequest() (request *CheckResourceRequest) {
	request = &CheckResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "CheckResource", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckResourceResponse creates a response to parse from CheckResource response
func CreateCheckResourceResponse() (response *CheckResourceResponse) {
	response = &CheckResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
