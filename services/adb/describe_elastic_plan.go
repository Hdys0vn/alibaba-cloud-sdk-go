package adb

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

// DescribeElasticPlan invokes the adb.DescribeElasticPlan API synchronously
func (client *Client) DescribeElasticPlan(request *DescribeElasticPlanRequest) (response *DescribeElasticPlanResponse, err error) {
	response = CreateDescribeElasticPlanResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeElasticPlanWithChan invokes the adb.DescribeElasticPlan API asynchronously
func (client *Client) DescribeElasticPlanWithChan(request *DescribeElasticPlanRequest) (<-chan *DescribeElasticPlanResponse, <-chan error) {
	responseChan := make(chan *DescribeElasticPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeElasticPlan(request)
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

// DescribeElasticPlanWithCallback invokes the adb.DescribeElasticPlan API asynchronously
func (client *Client) DescribeElasticPlanWithCallback(request *DescribeElasticPlanRequest, callback func(response *DescribeElasticPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeElasticPlanResponse
		var err error
		defer close(result)
		response, err = client.DescribeElasticPlan(request)
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

// DescribeElasticPlanRequest is the request struct for api DescribeElasticPlan
type DescribeElasticPlanRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ElasticPlanEnable    requests.Boolean `position:"Query" name:"ElasticPlanEnable"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ElasticPlanName      string           `position:"Query" name:"ElasticPlanName"`
	ResourcePoolName     string           `position:"Query" name:"ResourcePoolName"`
}

// DescribeElasticPlanResponse is the response struct for api DescribeElasticPlan
type DescribeElasticPlanResponse struct {
	*responses.BaseResponse
	RequestId       string            `json:"RequestId" xml:"RequestId"`
	ElasticPlanList []ElasticPlanInfo `json:"ElasticPlanList" xml:"ElasticPlanList"`
}

// CreateDescribeElasticPlanRequest creates a request to invoke DescribeElasticPlan API
func CreateDescribeElasticPlanRequest() (request *DescribeElasticPlanRequest) {
	request = &DescribeElasticPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "DescribeElasticPlan", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeElasticPlanResponse creates a response to parse from DescribeElasticPlan response
func CreateDescribeElasticPlanResponse() (response *DescribeElasticPlanResponse) {
	response = &DescribeElasticPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
