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

// ModifyStrategyTarget invokes the aegis.ModifyStrategyTarget API synchronously
// api document: https://help.aliyun.com/api/aegis/modifystrategytarget.html
func (client *Client) ModifyStrategyTarget(request *ModifyStrategyTargetRequest) (response *ModifyStrategyTargetResponse, err error) {
	response = CreateModifyStrategyTargetResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyStrategyTargetWithChan invokes the aegis.ModifyStrategyTarget API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifystrategytarget.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyStrategyTargetWithChan(request *ModifyStrategyTargetRequest) (<-chan *ModifyStrategyTargetResponse, <-chan error) {
	responseChan := make(chan *ModifyStrategyTargetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyStrategyTarget(request)
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

// ModifyStrategyTargetWithCallback invokes the aegis.ModifyStrategyTarget API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifystrategytarget.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyStrategyTargetWithCallback(request *ModifyStrategyTargetRequest, callback func(response *ModifyStrategyTargetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyStrategyTargetResponse
		var err error
		defer close(result)
		response, err = client.ModifyStrategyTarget(request)
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

// ModifyStrategyTargetRequest is the request struct for api ModifyStrategyTarget
type ModifyStrategyTargetRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	Type     string `position:"Query" name:"Type"`
	Config   string `position:"Query" name:"Config"`
	Target   string `position:"Query" name:"Target"`
}

// ModifyStrategyTargetResponse is the response struct for api ModifyStrategyTarget
type ModifyStrategyTargetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyStrategyTargetRequest creates a request to invoke ModifyStrategyTarget API
func CreateModifyStrategyTargetRequest() (request *ModifyStrategyTargetRequest) {
	request = &ModifyStrategyTargetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyStrategyTarget", "vipaegis", "openAPI")
	return
}

// CreateModifyStrategyTargetResponse creates a response to parse from ModifyStrategyTarget response
func CreateModifyStrategyTargetResponse() (response *ModifyStrategyTargetResponse) {
	response = &ModifyStrategyTargetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
