package dms_enterprise

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

// GetRuleNumLimitOfSLA invokes the dms_enterprise.GetRuleNumLimitOfSLA API synchronously
func (client *Client) GetRuleNumLimitOfSLA(request *GetRuleNumLimitOfSLARequest) (response *GetRuleNumLimitOfSLAResponse, err error) {
	response = CreateGetRuleNumLimitOfSLAResponse()
	err = client.DoAction(request, response)
	return
}

// GetRuleNumLimitOfSLAWithChan invokes the dms_enterprise.GetRuleNumLimitOfSLA API asynchronously
func (client *Client) GetRuleNumLimitOfSLAWithChan(request *GetRuleNumLimitOfSLARequest) (<-chan *GetRuleNumLimitOfSLAResponse, <-chan error) {
	responseChan := make(chan *GetRuleNumLimitOfSLAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRuleNumLimitOfSLA(request)
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

// GetRuleNumLimitOfSLAWithCallback invokes the dms_enterprise.GetRuleNumLimitOfSLA API asynchronously
func (client *Client) GetRuleNumLimitOfSLAWithCallback(request *GetRuleNumLimitOfSLARequest, callback func(response *GetRuleNumLimitOfSLAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRuleNumLimitOfSLAResponse
		var err error
		defer close(result)
		response, err = client.GetRuleNumLimitOfSLA(request)
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

// GetRuleNumLimitOfSLARequest is the request struct for api GetRuleNumLimitOfSLA
type GetRuleNumLimitOfSLARequest struct {
	*requests.RpcRequest
	DagId requests.Integer `position:"Query" name:"DagId"`
	Tid   requests.Integer `position:"Query" name:"Tid"`
}

// GetRuleNumLimitOfSLAResponse is the response struct for api GetRuleNumLimitOfSLA
type GetRuleNumLimitOfSLAResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	RuleNumLimit int    `json:"RuleNumLimit" xml:"RuleNumLimit"`
}

// CreateGetRuleNumLimitOfSLARequest creates a request to invoke GetRuleNumLimitOfSLA API
func CreateGetRuleNumLimitOfSLARequest() (request *GetRuleNumLimitOfSLARequest) {
	request = &GetRuleNumLimitOfSLARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetRuleNumLimitOfSLA", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetRuleNumLimitOfSLAResponse creates a response to parse from GetRuleNumLimitOfSLA response
func CreateGetRuleNumLimitOfSLAResponse() (response *GetRuleNumLimitOfSLAResponse) {
	response = &GetRuleNumLimitOfSLAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
