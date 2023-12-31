package vpc

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

// CreateVirtualPhysicalConnection invokes the vpc.CreateVirtualPhysicalConnection API synchronously
func (client *Client) CreateVirtualPhysicalConnection(request *CreateVirtualPhysicalConnectionRequest) (response *CreateVirtualPhysicalConnectionResponse, err error) {
	response = CreateCreateVirtualPhysicalConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVirtualPhysicalConnectionWithChan invokes the vpc.CreateVirtualPhysicalConnection API asynchronously
func (client *Client) CreateVirtualPhysicalConnectionWithChan(request *CreateVirtualPhysicalConnectionRequest) (<-chan *CreateVirtualPhysicalConnectionResponse, <-chan error) {
	responseChan := make(chan *CreateVirtualPhysicalConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVirtualPhysicalConnection(request)
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

// CreateVirtualPhysicalConnectionWithCallback invokes the vpc.CreateVirtualPhysicalConnection API asynchronously
func (client *Client) CreateVirtualPhysicalConnectionWithCallback(request *CreateVirtualPhysicalConnectionRequest, callback func(response *CreateVirtualPhysicalConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVirtualPhysicalConnectionResponse
		var err error
		defer close(result)
		response, err = client.CreateVirtualPhysicalConnection(request)
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

// CreateVirtualPhysicalConnectionRequest is the request struct for api CreateVirtualPhysicalConnection
type CreateVirtualPhysicalConnectionRequest struct {
	*requests.RpcRequest
	VpconnAliUid             requests.Integer                      `position:"Query" name:"VpconnAliUid"`
	OrderMode                string                                `position:"Query" name:"OrderMode"`
	VlanId                   requests.Integer                      `position:"Query" name:"VlanId"`
	VpconnUidResourceGroupId string                                `position:"Query" name:"VpconnUidResourceGroupId"`
	Description              string                                `position:"Query" name:"Description"`
	Spec                     string                                `position:"Query" name:"Spec"`
	ResourceGroupId          string                                `position:"Query" name:"ResourceGroupId"`
	Tag                      *[]CreateVirtualPhysicalConnectionTag `position:"Query" name:"Tag"  type:"Repeated"`
	DryRun                   requests.Boolean                      `position:"Query" name:"DryRun"`
	Token                    string                                `position:"Query" name:"Token"`
	PhysicalConnectionId     string                                `position:"Query" name:"PhysicalConnectionId"`
	Name                     string                                `position:"Query" name:"Name"`
}

// CreateVirtualPhysicalConnectionTag is a repeated param struct in CreateVirtualPhysicalConnectionRequest
type CreateVirtualPhysicalConnectionTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateVirtualPhysicalConnectionResponse is the response struct for api CreateVirtualPhysicalConnection
type CreateVirtualPhysicalConnectionResponse struct {
	*responses.BaseResponse
	RequestId                 string `json:"RequestId" xml:"RequestId"`
	VirtualPhysicalConnection string `json:"VirtualPhysicalConnection" xml:"VirtualPhysicalConnection"`
}

// CreateCreateVirtualPhysicalConnectionRequest creates a request to invoke CreateVirtualPhysicalConnection API
func CreateCreateVirtualPhysicalConnectionRequest() (request *CreateVirtualPhysicalConnectionRequest) {
	request = &CreateVirtualPhysicalConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateVirtualPhysicalConnection", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVirtualPhysicalConnectionResponse creates a response to parse from CreateVirtualPhysicalConnection response
func CreateCreateVirtualPhysicalConnectionResponse() (response *CreateVirtualPhysicalConnectionResponse) {
	response = &CreateVirtualPhysicalConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
