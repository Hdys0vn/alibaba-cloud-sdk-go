package ddosbgp

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

// ConfigSchedruleOnDemand invokes the ddosbgp.ConfigSchedruleOnDemand API synchronously
func (client *Client) ConfigSchedruleOnDemand(request *ConfigSchedruleOnDemandRequest) (response *ConfigSchedruleOnDemandResponse, err error) {
	response = CreateConfigSchedruleOnDemandResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigSchedruleOnDemandWithChan invokes the ddosbgp.ConfigSchedruleOnDemand API asynchronously
func (client *Client) ConfigSchedruleOnDemandWithChan(request *ConfigSchedruleOnDemandRequest) (<-chan *ConfigSchedruleOnDemandResponse, <-chan error) {
	responseChan := make(chan *ConfigSchedruleOnDemandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigSchedruleOnDemand(request)
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

// ConfigSchedruleOnDemandWithCallback invokes the ddosbgp.ConfigSchedruleOnDemand API asynchronously
func (client *Client) ConfigSchedruleOnDemandWithCallback(request *ConfigSchedruleOnDemandRequest, callback func(response *ConfigSchedruleOnDemandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigSchedruleOnDemandResponse
		var err error
		defer close(result)
		response, err = client.ConfigSchedruleOnDemand(request)
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

// ConfigSchedruleOnDemandRequest is the request struct for api ConfigSchedruleOnDemand
type ConfigSchedruleOnDemandRequest struct {
	*requests.RpcRequest
	TimeZone          string `position:"Query" name:"TimeZone"`
	RuleName          string `position:"Query" name:"RuleName"`
	RuleConditionMbps string `position:"Query" name:"RuleConditionMbps"`
	RuleAction        string `position:"Query" name:"RuleAction"`
	SourceIp          string `position:"Query" name:"SourceIp"`
	RuleUndoMode      string `position:"Query" name:"RuleUndoMode"`
	RuleUndoEndTime   string `position:"Query" name:"RuleUndoEndTime"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	RuleUndoBeginTime string `position:"Query" name:"RuleUndoBeginTime"`
	RuleConditionCnt  string `position:"Query" name:"RuleConditionCnt"`
	RuleSwitch        string `position:"Query" name:"RuleSwitch"`
	RuleConditionKpps string `position:"Query" name:"RuleConditionKpps"`
}

// ConfigSchedruleOnDemandResponse is the response struct for api ConfigSchedruleOnDemand
type ConfigSchedruleOnDemandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigSchedruleOnDemandRequest creates a request to invoke ConfigSchedruleOnDemand API
func CreateConfigSchedruleOnDemandRequest() (request *ConfigSchedruleOnDemandRequest) {
	request = &ConfigSchedruleOnDemandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddosbgp", "2018-07-20", "ConfigSchedruleOnDemand", "", "")
	request.Method = requests.POST
	return
}

// CreateConfigSchedruleOnDemandResponse creates a response to parse from ConfigSchedruleOnDemand response
func CreateConfigSchedruleOnDemandResponse() (response *ConfigSchedruleOnDemandResponse) {
	response = &ConfigSchedruleOnDemandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
