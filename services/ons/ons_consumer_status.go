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

// OnsConsumerStatus invokes the ons.OnsConsumerStatus API synchronously
func (client *Client) OnsConsumerStatus(request *OnsConsumerStatusRequest) (response *OnsConsumerStatusResponse, err error) {
	response = CreateOnsConsumerStatusResponse()
	err = client.DoAction(request, response)
	return
}

// OnsConsumerStatusWithChan invokes the ons.OnsConsumerStatus API asynchronously
func (client *Client) OnsConsumerStatusWithChan(request *OnsConsumerStatusRequest) (<-chan *OnsConsumerStatusResponse, <-chan error) {
	responseChan := make(chan *OnsConsumerStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsConsumerStatus(request)
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

// OnsConsumerStatusWithCallback invokes the ons.OnsConsumerStatus API asynchronously
func (client *Client) OnsConsumerStatusWithCallback(request *OnsConsumerStatusRequest, callback func(response *OnsConsumerStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsConsumerStatusResponse
		var err error
		defer close(result)
		response, err = client.OnsConsumerStatus(request)
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

// OnsConsumerStatusRequest is the request struct for api OnsConsumerStatus
type OnsConsumerStatusRequest struct {
	*requests.RpcRequest
	GroupId    string           `position:"Query" name:"GroupId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	NeedJstack requests.Boolean `position:"Query" name:"NeedJstack"`
	Detail     requests.Boolean `position:"Query" name:"Detail"`
}

// OnsConsumerStatusResponse is the response struct for api OnsConsumerStatus
type OnsConsumerStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateOnsConsumerStatusRequest creates a request to invoke OnsConsumerStatus API
func CreateOnsConsumerStatusRequest() (request *OnsConsumerStatusRequest) {
	request = &OnsConsumerStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsConsumerStatus", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsConsumerStatusResponse creates a response to parse from OnsConsumerStatus response
func CreateOnsConsumerStatusResponse() (response *OnsConsumerStatusResponse) {
	response = &OnsConsumerStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
