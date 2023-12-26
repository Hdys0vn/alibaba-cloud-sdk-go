package cms

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

// DeleteMetricRuleResources invokes the cms.DeleteMetricRuleResources API synchronously
func (client *Client) DeleteMetricRuleResources(request *DeleteMetricRuleResourcesRequest) (response *DeleteMetricRuleResourcesResponse, err error) {
	response = CreateDeleteMetricRuleResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMetricRuleResourcesWithChan invokes the cms.DeleteMetricRuleResources API asynchronously
func (client *Client) DeleteMetricRuleResourcesWithChan(request *DeleteMetricRuleResourcesRequest) (<-chan *DeleteMetricRuleResourcesResponse, <-chan error) {
	responseChan := make(chan *DeleteMetricRuleResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMetricRuleResources(request)
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

// DeleteMetricRuleResourcesWithCallback invokes the cms.DeleteMetricRuleResources API asynchronously
func (client *Client) DeleteMetricRuleResourcesWithCallback(request *DeleteMetricRuleResourcesRequest, callback func(response *DeleteMetricRuleResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMetricRuleResourcesResponse
		var err error
		defer close(result)
		response, err = client.DeleteMetricRuleResources(request)
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

// DeleteMetricRuleResourcesRequest is the request struct for api DeleteMetricRuleResources
type DeleteMetricRuleResourcesRequest struct {
	*requests.RpcRequest
	Resources string `position:"Query" name:"Resources"`
	RuleId    string `position:"Query" name:"RuleId"`
}

// DeleteMetricRuleResourcesResponse is the response struct for api DeleteMetricRuleResources
type DeleteMetricRuleResourcesResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteMetricRuleResourcesRequest creates a request to invoke DeleteMetricRuleResources API
func CreateDeleteMetricRuleResourcesRequest() (request *DeleteMetricRuleResourcesRequest) {
	request = &DeleteMetricRuleResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteMetricRuleResources", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteMetricRuleResourcesResponse creates a response to parse from DeleteMetricRuleResources response
func CreateDeleteMetricRuleResourcesResponse() (response *DeleteMetricRuleResourcesResponse) {
	response = &DeleteMetricRuleResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
