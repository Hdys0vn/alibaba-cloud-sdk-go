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

// CreateNetworkAcl invokes the vpc.CreateNetworkAcl API synchronously
func (client *Client) CreateNetworkAcl(request *CreateNetworkAclRequest) (response *CreateNetworkAclResponse, err error) {
	response = CreateCreateNetworkAclResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNetworkAclWithChan invokes the vpc.CreateNetworkAcl API asynchronously
func (client *Client) CreateNetworkAclWithChan(request *CreateNetworkAclRequest) (<-chan *CreateNetworkAclResponse, <-chan error) {
	responseChan := make(chan *CreateNetworkAclResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNetworkAcl(request)
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

// CreateNetworkAclWithCallback invokes the vpc.CreateNetworkAcl API asynchronously
func (client *Client) CreateNetworkAclWithCallback(request *CreateNetworkAclRequest, callback func(response *CreateNetworkAclResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNetworkAclResponse
		var err error
		defer close(result)
		response, err = client.CreateNetworkAcl(request)
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

// CreateNetworkAclRequest is the request struct for api CreateNetworkAcl
type CreateNetworkAclRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer       `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string                 `position:"Query" name:"ClientToken"`
	Description          string                 `position:"Query" name:"Description"`
	Tag                  *[]CreateNetworkAclTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount string                 `position:"Query" name:"ResourceOwnerAccount"`
	NetworkAclName       string                 `position:"Query" name:"NetworkAclName"`
	OwnerId              requests.Integer       `position:"Query" name:"OwnerId"`
	VpcId                string                 `position:"Query" name:"VpcId"`
}

// CreateNetworkAclTag is a repeated param struct in CreateNetworkAclRequest
type CreateNetworkAclTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateNetworkAclResponse is the response struct for api CreateNetworkAcl
type CreateNetworkAclResponse struct {
	*responses.BaseResponse
	NetworkAclId        string              `json:"NetworkAclId" xml:"NetworkAclId"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	NetworkAclAttribute NetworkAclAttribute `json:"NetworkAclAttribute" xml:"NetworkAclAttribute"`
}

// CreateCreateNetworkAclRequest creates a request to invoke CreateNetworkAcl API
func CreateCreateNetworkAclRequest() (request *CreateNetworkAclRequest) {
	request = &CreateNetworkAclRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateNetworkAcl", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNetworkAclResponse creates a response to parse from CreateNetworkAcl response
func CreateCreateNetworkAclResponse() (response *CreateNetworkAclResponse) {
	response = &CreateNetworkAclResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
