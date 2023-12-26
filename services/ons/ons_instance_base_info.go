package ons

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

// OnsInstanceBaseInfo invokes the ons.OnsInstanceBaseInfo API synchronously
func (client *Client) OnsInstanceBaseInfo(request *OnsInstanceBaseInfoRequest) (response *OnsInstanceBaseInfoResponse, err error) {
	response = CreateOnsInstanceBaseInfoResponse()
	err = client.DoAction(request, response)
	return
}

// OnsInstanceBaseInfoWithChan invokes the ons.OnsInstanceBaseInfo API asynchronously
func (client *Client) OnsInstanceBaseInfoWithChan(request *OnsInstanceBaseInfoRequest) (<-chan *OnsInstanceBaseInfoResponse, <-chan error) {
	responseChan := make(chan *OnsInstanceBaseInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsInstanceBaseInfo(request)
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

// OnsInstanceBaseInfoWithCallback invokes the ons.OnsInstanceBaseInfo API asynchronously
func (client *Client) OnsInstanceBaseInfoWithCallback(request *OnsInstanceBaseInfoRequest, callback func(response *OnsInstanceBaseInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsInstanceBaseInfoResponse
		var err error
		defer close(result)
		response, err = client.OnsInstanceBaseInfo(request)
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

// OnsInstanceBaseInfoRequest is the request struct for api OnsInstanceBaseInfo
type OnsInstanceBaseInfoRequest struct {
	*requests.RpcRequest
	UserId     string `position:"Query" name:"UserId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// OnsInstanceBaseInfoResponse is the response struct for api OnsInstanceBaseInfo
type OnsInstanceBaseInfoResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	HelpUrl          string           `json:"HelpUrl" xml:"HelpUrl"`
	InstanceBaseInfo InstanceBaseInfo `json:"InstanceBaseInfo" xml:"InstanceBaseInfo"`
}

// CreateOnsInstanceBaseInfoRequest creates a request to invoke OnsInstanceBaseInfo API
func CreateOnsInstanceBaseInfoRequest() (request *OnsInstanceBaseInfoRequest) {
	request = &OnsInstanceBaseInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsInstanceBaseInfo", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsInstanceBaseInfoResponse creates a response to parse from OnsInstanceBaseInfo response
func CreateOnsInstanceBaseInfoResponse() (response *OnsInstanceBaseInfoResponse) {
	response = &OnsInstanceBaseInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}