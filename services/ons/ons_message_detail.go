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

// OnsMessageDetail invokes the ons.OnsMessageDetail API synchronously
func (client *Client) OnsMessageDetail(request *OnsMessageDetailRequest) (response *OnsMessageDetailResponse, err error) {
	response = CreateOnsMessageDetailResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMessageDetailWithChan invokes the ons.OnsMessageDetail API asynchronously
func (client *Client) OnsMessageDetailWithChan(request *OnsMessageDetailRequest) (<-chan *OnsMessageDetailResponse, <-chan error) {
	responseChan := make(chan *OnsMessageDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMessageDetail(request)
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

// OnsMessageDetailWithCallback invokes the ons.OnsMessageDetail API asynchronously
func (client *Client) OnsMessageDetailWithCallback(request *OnsMessageDetailRequest, callback func(response *OnsMessageDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMessageDetailResponse
		var err error
		defer close(result)
		response, err = client.OnsMessageDetail(request)
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

// OnsMessageDetailRequest is the request struct for api OnsMessageDetail
type OnsMessageDetailRequest struct {
	*requests.RpcRequest
	MsgId      string `position:"Query" name:"MsgId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
}

// OnsMessageDetailResponse is the response struct for api OnsMessageDetail
type OnsMessageDetailResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	HelpUrl   string                 `json:"HelpUrl" xml:"HelpUrl"`
	Data      DataInOnsMessageDetail `json:"Data" xml:"Data"`
}

// CreateOnsMessageDetailRequest creates a request to invoke OnsMessageDetail API
func CreateOnsMessageDetailRequest() (request *OnsMessageDetailRequest) {
	request = &OnsMessageDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMessageDetail", "ons", "openAPI")
	request.Method = requests.GET
	return
}

// CreateOnsMessageDetailResponse creates a response to parse from OnsMessageDetail response
func CreateOnsMessageDetailResponse() (response *OnsMessageDetailResponse) {
	response = &OnsMessageDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}