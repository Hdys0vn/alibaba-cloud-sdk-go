package r_kvstore

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

// CreateInstances invokes the r_kvstore.CreateInstances API synchronously
func (client *Client) CreateInstances(request *CreateInstancesRequest) (response *CreateInstancesResponse, err error) {
	response = CreateCreateInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// CreateInstancesWithChan invokes the r_kvstore.CreateInstances API asynchronously
func (client *Client) CreateInstancesWithChan(request *CreateInstancesRequest) (<-chan *CreateInstancesResponse, <-chan error) {
	responseChan := make(chan *CreateInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateInstances(request)
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

// CreateInstancesWithCallback invokes the r_kvstore.CreateInstances API asynchronously
func (client *Client) CreateInstancesWithCallback(request *CreateInstancesRequest, callback func(response *CreateInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateInstancesResponse
		var err error
		defer close(result)
		response, err = client.CreateInstances(request)
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

// CreateInstancesRequest is the request struct for api CreateInstances
type CreateInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Instances            string           `position:"Query" name:"Instances"`
	CouponNo             string           `position:"Query" name:"CouponNo"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	RebuildInstance      requests.Boolean `position:"Query" name:"RebuildInstance"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	AgentId              string           `position:"Query" name:"AgentId"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Token                string           `position:"Query" name:"Token"`
	PrivateIpAddress     string           `position:"Query" name:"PrivateIpAddress"`
	AutoRenew            string           `position:"Query" name:"AutoRenew"`
}

// CreateInstancesResponse is the response struct for api CreateInstances
type CreateInstancesResponse struct {
	*responses.BaseResponse
	OrderId     string      `json:"OrderId" xml:"OrderId"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	InstanceIds InstanceIds `json:"InstanceIds" xml:"InstanceIds"`
}

// CreateCreateInstancesRequest creates a request to invoke CreateInstances API
func CreateCreateInstancesRequest() (request *CreateInstancesRequest) {
	request = &CreateInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateInstances", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateInstancesResponse creates a response to parse from CreateInstances response
func CreateCreateInstancesResponse() (response *CreateInstancesResponse) {
	response = &CreateInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
