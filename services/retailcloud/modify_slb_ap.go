package retailcloud

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

// ModifySlbAP invokes the retailcloud.ModifySlbAP API synchronously
func (client *Client) ModifySlbAP(request *ModifySlbAPRequest) (response *ModifySlbAPResponse, err error) {
	response = CreateModifySlbAPResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySlbAPWithChan invokes the retailcloud.ModifySlbAP API asynchronously
func (client *Client) ModifySlbAPWithChan(request *ModifySlbAPRequest) (<-chan *ModifySlbAPResponse, <-chan error) {
	responseChan := make(chan *ModifySlbAPResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySlbAP(request)
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

// ModifySlbAPWithCallback invokes the retailcloud.ModifySlbAP API asynchronously
func (client *Client) ModifySlbAPWithCallback(request *ModifySlbAPRequest, callback func(response *ModifySlbAPResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySlbAPResponse
		var err error
		defer close(result)
		response, err = client.ModifySlbAP(request)
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

// ModifySlbAPRequest is the request struct for api ModifySlbAP
type ModifySlbAPRequest struct {
	*requests.RpcRequest
	SslCertId          string           `position:"Query" name:"SslCertId"`
	EstablishedTimeout requests.Integer `position:"Query" name:"EstablishedTimeout"`
	RealServerPort     requests.Integer `position:"Query" name:"RealServerPort"`
	StickySession      requests.Integer `position:"Query" name:"StickySession"`
	CookieTimeout      requests.Integer `position:"Query" name:"CookieTimeout"`
	Name               string           `position:"Query" name:"Name"`
	SlbAPId            requests.Integer `position:"Query" name:"SlbAPId"`
}

// ModifySlbAPResponse is the response struct for api ModifySlbAP
type ModifySlbAPResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifySlbAPRequest creates a request to invoke ModifySlbAP API
func CreateModifySlbAPRequest() (request *ModifySlbAPRequest) {
	request = &ModifySlbAPRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ModifySlbAP", "", "")
	request.Method = requests.POST
	return
}

// CreateModifySlbAPResponse creates a response to parse from ModifySlbAP response
func CreateModifySlbAPResponse() (response *ModifySlbAPResponse) {
	response = &ModifySlbAPResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
