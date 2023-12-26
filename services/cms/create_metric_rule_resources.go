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

// CreateMetricRuleResources invokes the cms.CreateMetricRuleResources API synchronously
func (client *Client) CreateMetricRuleResources(request *CreateMetricRuleResourcesRequest) (response *CreateMetricRuleResourcesResponse, err error) {
	response = CreateCreateMetricRuleResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMetricRuleResourcesWithChan invokes the cms.CreateMetricRuleResources API asynchronously
func (client *Client) CreateMetricRuleResourcesWithChan(request *CreateMetricRuleResourcesRequest) (<-chan *CreateMetricRuleResourcesResponse, <-chan error) {
	responseChan := make(chan *CreateMetricRuleResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMetricRuleResources(request)
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

// CreateMetricRuleResourcesWithCallback invokes the cms.CreateMetricRuleResources API asynchronously
func (client *Client) CreateMetricRuleResourcesWithCallback(request *CreateMetricRuleResourcesRequest, callback func(response *CreateMetricRuleResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMetricRuleResourcesResponse
		var err error
		defer close(result)
		response, err = client.CreateMetricRuleResources(request)
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

// CreateMetricRuleResourcesRequest is the request struct for api CreateMetricRuleResources
type CreateMetricRuleResourcesRequest struct {
	*requests.RpcRequest
	Resources string `position:"Query" name:"Resources"`
	RuleId    string `position:"Query" name:"RuleId"`
	Overwrite string `position:"Query" name:"Overwrite"`
}

// CreateMetricRuleResourcesResponse is the response struct for api CreateMetricRuleResources
type CreateMetricRuleResourcesResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateMetricRuleResourcesRequest creates a request to invoke CreateMetricRuleResources API
func CreateCreateMetricRuleResourcesRequest() (request *CreateMetricRuleResourcesRequest) {
	request = &CreateMetricRuleResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateMetricRuleResources", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMetricRuleResourcesResponse creates a response to parse from CreateMetricRuleResources response
func CreateCreateMetricRuleResourcesResponse() (response *CreateMetricRuleResourcesResponse) {
	response = &CreateMetricRuleResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}