package arms

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

// CreatePrometheusAlertRule invokes the arms.CreatePrometheusAlertRule API synchronously
func (client *Client) CreatePrometheusAlertRule(request *CreatePrometheusAlertRuleRequest) (response *CreatePrometheusAlertRuleResponse, err error) {
	response = CreateCreatePrometheusAlertRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePrometheusAlertRuleWithChan invokes the arms.CreatePrometheusAlertRule API asynchronously
func (client *Client) CreatePrometheusAlertRuleWithChan(request *CreatePrometheusAlertRuleRequest) (<-chan *CreatePrometheusAlertRuleResponse, <-chan error) {
	responseChan := make(chan *CreatePrometheusAlertRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePrometheusAlertRule(request)
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

// CreatePrometheusAlertRuleWithCallback invokes the arms.CreatePrometheusAlertRule API asynchronously
func (client *Client) CreatePrometheusAlertRuleWithCallback(request *CreatePrometheusAlertRuleRequest, callback func(response *CreatePrometheusAlertRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePrometheusAlertRuleResponse
		var err error
		defer close(result)
		response, err = client.CreatePrometheusAlertRule(request)
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

// CreatePrometheusAlertRuleRequest is the request struct for api CreatePrometheusAlertRule
type CreatePrometheusAlertRuleRequest struct {
	*requests.RpcRequest
	ProductCode    string                           `position:"Query" name:"ProductCode"`
	Expression     string                           `position:"Query" name:"Expression"`
	AlertName      string                           `position:"Query" name:"AlertName"`
	Annotations    string                           `position:"Query" name:"Annotations"`
	ClusterId      string                           `position:"Query" name:"ClusterId"`
	DispatchRuleId requests.Integer                 `position:"Query" name:"DispatchRuleId"`
	ProxyUserId    string                           `position:"Query" name:"ProxyUserId"`
	Type           string                           `position:"Query" name:"Type"`
	Message        string                           `position:"Query" name:"Message"`
	Labels         string                           `position:"Query" name:"Labels"`
	Tags           *[]CreatePrometheusAlertRuleTags `position:"Query" name:"Tags"  type:"Repeated"`
	Duration       string                           `position:"Query" name:"Duration"`
	NotifyType     string                           `position:"Query" name:"NotifyType"`
}

// CreatePrometheusAlertRuleTags is a repeated param struct in CreatePrometheusAlertRuleRequest
type CreatePrometheusAlertRuleTags struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreatePrometheusAlertRuleResponse is the response struct for api CreatePrometheusAlertRule
type CreatePrometheusAlertRuleResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	PrometheusAlertRule PrometheusAlertRule `json:"PrometheusAlertRule" xml:"PrometheusAlertRule"`
}

// CreateCreatePrometheusAlertRuleRequest creates a request to invoke CreatePrometheusAlertRule API
func CreateCreatePrometheusAlertRuleRequest() (request *CreatePrometheusAlertRuleRequest) {
	request = &CreatePrometheusAlertRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "CreatePrometheusAlertRule", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePrometheusAlertRuleResponse creates a response to parse from CreatePrometheusAlertRule response
func CreateCreatePrometheusAlertRuleResponse() (response *CreatePrometheusAlertRuleResponse) {
	response = &CreatePrometheusAlertRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
