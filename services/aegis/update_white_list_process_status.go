package aegis

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

// UpdateWhiteListProcessStatus invokes the aegis.UpdateWhiteListProcessStatus API synchronously
// api document: https://help.aliyun.com/api/aegis/updatewhitelistprocessstatus.html
func (client *Client) UpdateWhiteListProcessStatus(request *UpdateWhiteListProcessStatusRequest) (response *UpdateWhiteListProcessStatusResponse, err error) {
	response = CreateUpdateWhiteListProcessStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWhiteListProcessStatusWithChan invokes the aegis.UpdateWhiteListProcessStatus API asynchronously
// api document: https://help.aliyun.com/api/aegis/updatewhitelistprocessstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateWhiteListProcessStatusWithChan(request *UpdateWhiteListProcessStatusRequest) (<-chan *UpdateWhiteListProcessStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateWhiteListProcessStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWhiteListProcessStatus(request)
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

// UpdateWhiteListProcessStatusWithCallback invokes the aegis.UpdateWhiteListProcessStatus API asynchronously
// api document: https://help.aliyun.com/api/aegis/updatewhitelistprocessstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateWhiteListProcessStatusWithCallback(request *UpdateWhiteListProcessStatusRequest, callback func(response *UpdateWhiteListProcessStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWhiteListProcessStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateWhiteListProcessStatus(request)
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

// UpdateWhiteListProcessStatusRequest is the request struct for api UpdateWhiteListProcessStatus
type UpdateWhiteListProcessStatusRequest struct {
	*requests.RpcRequest
	ProcessIds string           `position:"Query" name:"ProcessIds"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
	StrategyId requests.Integer `position:"Query" name:"StrategyId"`
	Lang       string           `position:"Query" name:"Lang"`
	Status     requests.Integer `position:"Query" name:"Status"`
}

// UpdateWhiteListProcessStatusResponse is the response struct for api UpdateWhiteListProcessStatus
type UpdateWhiteListProcessStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateWhiteListProcessStatusRequest creates a request to invoke UpdateWhiteListProcessStatus API
func CreateUpdateWhiteListProcessStatusRequest() (request *UpdateWhiteListProcessStatusRequest) {
	request = &UpdateWhiteListProcessStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "UpdateWhiteListProcessStatus", "vipaegis", "openAPI")
	return
}

// CreateUpdateWhiteListProcessStatusResponse creates a response to parse from UpdateWhiteListProcessStatus response
func CreateUpdateWhiteListProcessStatusResponse() (response *UpdateWhiteListProcessStatusResponse) {
	response = &UpdateWhiteListProcessStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
