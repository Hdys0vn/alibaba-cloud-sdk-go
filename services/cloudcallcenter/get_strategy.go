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

// GetStrategy invokes the cloudcallcenter.GetStrategy API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getstrategy.html
func (client *Client) GetStrategy(request *GetStrategyRequest) (response *GetStrategyResponse, err error) {
	response = CreateGetStrategyResponse()
	err = client.DoAction(request, response)
	return
}

// GetStrategyWithChan invokes the cloudcallcenter.GetStrategy API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getstrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetStrategyWithChan(request *GetStrategyRequest) (<-chan *GetStrategyResponse, <-chan error) {
	responseChan := make(chan *GetStrategyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetStrategy(request)
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

// GetStrategyWithCallback invokes the cloudcallcenter.GetStrategy API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getstrategy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetStrategyWithCallback(request *GetStrategyRequest, callback func(response *GetStrategyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetStrategyResponse
		var err error
		defer close(result)
		response, err = client.GetStrategy(request)
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

// GetStrategyRequest is the request struct for api GetStrategy
type GetStrategyRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	StrategyId string `position:"Query" name:"StrategyId"`
}

// GetStrategyResponse is the response struct for api GetStrategy
type GetStrategyResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Strategy       Strategy `json:"Strategy" xml:"Strategy"`
}

// CreateGetStrategyRequest creates a request to invoke GetStrategy API
func CreateGetStrategyRequest() (request *GetStrategyRequest) {
	request = &GetStrategyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetStrategy", "", "")
	request.Method = requests.POST
	return
}

// CreateGetStrategyResponse creates a response to parse from GetStrategy response
func CreateGetStrategyResponse() (response *GetStrategyResponse) {
	response = &GetStrategyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
