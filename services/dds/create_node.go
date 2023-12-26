package dds

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

// CreateNode invokes the dds.CreateNode API synchronously
func (client *Client) CreateNode(request *CreateNodeRequest) (response *CreateNodeResponse, err error) {
	response = CreateCreateNodeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNodeWithChan invokes the dds.CreateNode API asynchronously
func (client *Client) CreateNodeWithChan(request *CreateNodeRequest) (<-chan *CreateNodeResponse, <-chan error) {
	responseChan := make(chan *CreateNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNode(request)
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

// CreateNodeWithCallback invokes the dds.CreateNode API asynchronously
func (client *Client) CreateNodeWithCallback(request *CreateNodeRequest, callback func(response *CreateNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNodeResponse
		var err error
		defer close(result)
		response, err = client.CreateNode(request)
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

// CreateNodeRequest is the request struct for api CreateNode
type CreateNodeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NodeType             string           `position:"Query" name:"NodeType"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ReadonlyReplicas     requests.Integer `position:"Query" name:"ReadonlyReplicas"`
	CouponNo             string           `position:"Query" name:"CouponNo"`
	NodeClass            string           `position:"Query" name:"NodeClass"`
	ShardDirect          requests.Boolean `position:"Query" name:"ShardDirect"`
	AccountName          string           `position:"Query" name:"AccountName"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	FromApp              string           `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	NodeStorage          requests.Integer `position:"Query" name:"NodeStorage"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AccountPassword      string           `position:"Query" name:"AccountPassword"`
}

// CreateNodeResponse is the response struct for api CreateNode
type CreateNodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	NodeId    string `json:"NodeId" xml:"NodeId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateNodeRequest creates a request to invoke CreateNode API
func CreateCreateNodeRequest() (request *CreateNodeRequest) {
	request = &CreateNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "CreateNode", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNodeResponse creates a response to parse from CreateNode response
func CreateCreateNodeResponse() (response *CreateNodeResponse) {
	response = &CreateNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}