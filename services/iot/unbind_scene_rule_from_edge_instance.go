package iot

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

// UnbindSceneRuleFromEdgeInstance invokes the iot.UnbindSceneRuleFromEdgeInstance API synchronously
func (client *Client) UnbindSceneRuleFromEdgeInstance(request *UnbindSceneRuleFromEdgeInstanceRequest) (response *UnbindSceneRuleFromEdgeInstanceResponse, err error) {
	response = CreateUnbindSceneRuleFromEdgeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindSceneRuleFromEdgeInstanceWithChan invokes the iot.UnbindSceneRuleFromEdgeInstance API asynchronously
func (client *Client) UnbindSceneRuleFromEdgeInstanceWithChan(request *UnbindSceneRuleFromEdgeInstanceRequest) (<-chan *UnbindSceneRuleFromEdgeInstanceResponse, <-chan error) {
	responseChan := make(chan *UnbindSceneRuleFromEdgeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindSceneRuleFromEdgeInstance(request)
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

// UnbindSceneRuleFromEdgeInstanceWithCallback invokes the iot.UnbindSceneRuleFromEdgeInstance API asynchronously
func (client *Client) UnbindSceneRuleFromEdgeInstanceWithCallback(request *UnbindSceneRuleFromEdgeInstanceRequest, callback func(response *UnbindSceneRuleFromEdgeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindSceneRuleFromEdgeInstanceResponse
		var err error
		defer close(result)
		response, err = client.UnbindSceneRuleFromEdgeInstance(request)
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

// UnbindSceneRuleFromEdgeInstanceRequest is the request struct for api UnbindSceneRuleFromEdgeInstance
type UnbindSceneRuleFromEdgeInstanceRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	RuleId        string `position:"Query" name:"RuleId"`
}

// UnbindSceneRuleFromEdgeInstanceResponse is the response struct for api UnbindSceneRuleFromEdgeInstance
type UnbindSceneRuleFromEdgeInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUnbindSceneRuleFromEdgeInstanceRequest creates a request to invoke UnbindSceneRuleFromEdgeInstance API
func CreateUnbindSceneRuleFromEdgeInstanceRequest() (request *UnbindSceneRuleFromEdgeInstanceRequest) {
	request = &UnbindSceneRuleFromEdgeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UnbindSceneRuleFromEdgeInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateUnbindSceneRuleFromEdgeInstanceResponse creates a response to parse from UnbindSceneRuleFromEdgeInstance response
func CreateUnbindSceneRuleFromEdgeInstanceResponse() (response *UnbindSceneRuleFromEdgeInstanceResponse) {
	response = &UnbindSceneRuleFromEdgeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
