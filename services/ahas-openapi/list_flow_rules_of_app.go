package ahas_openapi

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

// ListFlowRulesOfApp invokes the ahas_openapi.ListFlowRulesOfApp API synchronously
func (client *Client) ListFlowRulesOfApp(request *ListFlowRulesOfAppRequest) (response *ListFlowRulesOfAppResponse, err error) {
	response = CreateListFlowRulesOfAppResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlowRulesOfAppWithChan invokes the ahas_openapi.ListFlowRulesOfApp API asynchronously
func (client *Client) ListFlowRulesOfAppWithChan(request *ListFlowRulesOfAppRequest) (<-chan *ListFlowRulesOfAppResponse, <-chan error) {
	responseChan := make(chan *ListFlowRulesOfAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlowRulesOfApp(request)
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

// ListFlowRulesOfAppWithCallback invokes the ahas_openapi.ListFlowRulesOfApp API asynchronously
func (client *Client) ListFlowRulesOfAppWithCallback(request *ListFlowRulesOfAppRequest, callback func(response *ListFlowRulesOfAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlowRulesOfAppResponse
		var err error
		defer close(result)
		response, err = client.ListFlowRulesOfApp(request)
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

// ListFlowRulesOfAppRequest is the request struct for api ListFlowRulesOfApp
type ListFlowRulesOfAppRequest struct {
	*requests.RpcRequest
	AhasRegionId string           `position:"Query" name:"AhasRegionId"`
	AppName      string           `position:"Query" name:"AppName"`
	Namespace    string           `position:"Query" name:"Namespace"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	PageIndex    requests.Integer `position:"Query" name:"PageIndex"`
}

// ListFlowRulesOfAppResponse is the response struct for api ListFlowRulesOfApp
type ListFlowRulesOfAppResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListFlowRulesOfAppRequest creates a request to invoke ListFlowRulesOfApp API
func CreateListFlowRulesOfAppRequest() (request *ListFlowRulesOfAppRequest) {
	request = &ListFlowRulesOfAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ahas-openapi", "2019-09-01", "ListFlowRulesOfApp", "ahas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListFlowRulesOfAppResponse creates a response to parse from ListFlowRulesOfApp response
func CreateListFlowRulesOfAppResponse() (response *ListFlowRulesOfAppResponse) {
	response = &ListFlowRulesOfAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}