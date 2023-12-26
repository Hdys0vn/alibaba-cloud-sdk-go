package slb

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

// SetRule invokes the slb.SetRule API synchronously
func (client *Client) SetRule(request *SetRuleRequest) (response *SetRuleResponse, err error) {
	response = CreateSetRuleResponse()
	err = client.DoAction(request, response)
	return
}

// SetRuleWithChan invokes the slb.SetRule API asynchronously
func (client *Client) SetRuleWithChan(request *SetRuleRequest) (<-chan *SetRuleResponse, <-chan error) {
	responseChan := make(chan *SetRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetRule(request)
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

// SetRuleWithCallback invokes the slb.SetRule API asynchronously
func (client *Client) SetRuleWithCallback(request *SetRuleRequest, callback func(response *SetRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetRuleResponse
		var err error
		defer close(result)
		response, err = client.SetRule(request)
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

// SetRuleRequest is the request struct for api SetRule
type SetRuleRequest struct {
	*requests.RpcRequest
	AccessKeyId            string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HealthCheckTimeout     requests.Integer `position:"Query" name:"HealthCheckTimeout"`
	HealthCheckURI         string           `position:"Query" name:"HealthCheckURI"`
	RuleName               string           `position:"Query" name:"RuleName"`
	UnhealthyThreshold     requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	HealthyThreshold       requests.Integer `position:"Query" name:"HealthyThreshold"`
	Scheduler              string           `position:"Query" name:"Scheduler"`
	HealthCheck            string           `position:"Query" name:"HealthCheck"`
	ListenerSync           string           `position:"Query" name:"ListenerSync"`
	CookieTimeout          requests.Integer `position:"Query" name:"CookieTimeout"`
	StickySessionType      string           `position:"Query" name:"StickySessionType"`
	VServerGroupId         string           `position:"Query" name:"VServerGroupId"`
	Cookie                 string           `position:"Query" name:"Cookie"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	StickySession          string           `position:"Query" name:"StickySession"`
	HealthCheckDomain      string           `position:"Query" name:"HealthCheckDomain"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	Tags                   string           `position:"Query" name:"Tags"`
	HealthCheckInterval    requests.Integer `position:"Query" name:"HealthCheckInterval"`
	RuleId                 string           `position:"Query" name:"RuleId"`
	HealthCheckConnectPort requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	HealthCheckHttpCode    string           `position:"Query" name:"HealthCheckHttpCode"`
}

// SetRuleResponse is the response struct for api SetRule
type SetRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetRuleRequest creates a request to invoke SetRule API
func CreateSetRuleRequest() (request *SetRuleRequest) {
	request = &SetRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetRule", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetRuleResponse creates a response to parse from SetRule response
func CreateSetRuleResponse() (response *SetRuleResponse) {
	response = &SetRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
