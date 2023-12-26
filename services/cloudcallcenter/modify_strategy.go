package cloudcallcenter

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

// ModifyStrategy invokes the cloudcallcenter.ModifyStrategy API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifystrategy.html
func (client *Client) ModifyStrategy(request *ModifyStrategyRequest) (response *ModifyStrategyResponse, err error) {
	response = CreateModifyStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyStrategyWithChan invokes the cloudcallcenter.ModifyStrategy API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifystrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyStrategyWithChan(request *ModifyStrategyRequest) (<-chan *ModifyStrategyResponse, <-chan error) {
	responseChan := make(chan *ModifyStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyStrategy(request)
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

// ModifyStrategyWithCallback invokes the cloudcallcenter.ModifyStrategy API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifystrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyStrategyWithCallback(request *ModifyStrategyRequest, callback func(response *ModifyStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyStrategyResponse
		var err error
		defer close(result)
		response, err = client.ModifyStrategy(request)
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

// ModifyStrategyRequest is the request struct for api ModifyStrategy
type ModifyStrategyRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	StrategyJson string `position:"Query" name:"StrategyJson"`
	StrategyId   string `position:"Query" name:"StrategyId"`
}

// ModifyStrategyResponse is the response struct for api ModifyStrategy
type ModifyStrategyResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Strategy       Strategy `json:"Strategy" xml:"Strategy"`
}

// CreateModifyStrategyRequest creates a request to invoke ModifyStrategy API
func CreateModifyStrategyRequest() (request *ModifyStrategyRequest) {
	request = &ModifyStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyStrategy", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyStrategyResponse creates a response to parse from ModifyStrategy response
func CreateModifyStrategyResponse() (response *ModifyStrategyResponse) {
	response = &ModifyStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}