package amp

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

// DingAPITcTc invokes the amp.DingAPITcTc API synchronously
func (client *Client) DingAPITcTc(request *DingAPITcTcRequest) (response *DingAPITcTcResponse, err error) {
	response = CreateDingAPITcTcResponse()
	err = client.DoAction(request, response)
	return
}

// DingAPITcTcWithChan invokes the amp.DingAPITcTc API asynchronously
func (client *Client) DingAPITcTcWithChan(request *DingAPITcTcRequest) (<-chan *DingAPITcTcResponse, <-chan error) {
	responseChan := make(chan *DingAPITcTcResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DingAPITcTc(request)
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

// DingAPITcTcWithCallback invokes the amp.DingAPITcTc API asynchronously
func (client *Client) DingAPITcTcWithCallback(request *DingAPITcTcRequest, callback func(response *DingAPITcTcResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DingAPITcTcResponse
		var err error
		defer close(result)
		response, err = client.DingAPITcTc(request)
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

// DingAPITcTcRequest is the request struct for api DingAPITcTc
type DingAPITcTcRequest struct {
	*requests.RoaRequest
}

// DingAPITcTcResponse is the response struct for api DingAPITcTc
type DingAPITcTcResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Res       Res    `json:"Res" xml:"Res"`
}

// CreateDingAPITcTcRequest creates a request to invoke DingAPITcTc API
func CreateDingAPITcTcRequest() (request *DingAPITcTcRequest) {
	request = &DingAPITcTcRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("amp", "2020-07-08", "DingAPITcTc", "/ding/tctc", "ServiceCode", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDingAPITcTcResponse creates a response to parse from DingAPITcTc response
func CreateDingAPITcTcResponse() (response *DingAPITcTcResponse) {
	response = &DingAPITcTcResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
