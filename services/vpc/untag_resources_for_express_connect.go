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

// UntagResourcesForExpressConnect invokes the vpc.UntagResourcesForExpressConnect API synchronously
func (client *Client) UntagResourcesForExpressConnect(request *UntagResourcesForExpressConnectRequest) (response *UntagResourcesForExpressConnectResponse, err error) {
	response = CreateUntagResourcesForExpressConnectResponse()
	err = client.DoAction(request, response)
	return
}

// UntagResourcesForExpressConnectWithChan invokes the vpc.UntagResourcesForExpressConnect API asynchronously
func (client *Client) UntagResourcesForExpressConnectWithChan(request *UntagResourcesForExpressConnectRequest) (<-chan *UntagResourcesForExpressConnectResponse, <-chan error) {
	responseChan := make(chan *UntagResourcesForExpressConnectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UntagResourcesForExpressConnect(request)
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

// UntagResourcesForExpressConnectWithCallback invokes the vpc.UntagResourcesForExpressConnect API asynchronously
func (client *Client) UntagResourcesForExpressConnectWithCallback(request *UntagResourcesForExpressConnectRequest, callback func(response *UntagResourcesForExpressConnectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UntagResourcesForExpressConnectResponse
		var err error
		defer close(result)
		response, err = client.UntagResourcesForExpressConnect(request)
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

// UntagResourcesForExpressConnectRequest is the request struct for api UntagResourcesForExpressConnect
type UntagResourcesForExpressConnectRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                      `position:"Query" name:"ResourceOwnerId"`
	Tag                  *[]UntagResourcesForExpressConnectTag `position:"Query" name:"Tag"  type:"Repeated"`
	All                  requests.Boolean                      `position:"Query" name:"All"`
	ResourceId           *[]string                             `position:"Query" name:"ResourceId"  type:"Repeated"`
	ResourceOwnerAccount string                                `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                                `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                      `position:"Query" name:"OwnerId"`
	ResourceType         string                                `position:"Query" name:"ResourceType"`
	TagKey               *[]string                             `position:"Query" name:"TagKey"  type:"Repeated"`
}

// UntagResourcesForExpressConnectTag is a repeated param struct in UntagResourcesForExpressConnectRequest
type UntagResourcesForExpressConnectTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// UntagResourcesForExpressConnectResponse is the response struct for api UntagResourcesForExpressConnect
type UntagResourcesForExpressConnectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUntagResourcesForExpressConnectRequest creates a request to invoke UntagResourcesForExpressConnect API
func CreateUntagResourcesForExpressConnectRequest() (request *UntagResourcesForExpressConnectRequest) {
	request = &UntagResourcesForExpressConnectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "UntagResourcesForExpressConnect", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUntagResourcesForExpressConnectResponse creates a response to parse from UntagResourcesForExpressConnect response
func CreateUntagResourcesForExpressConnectResponse() (response *UntagResourcesForExpressConnectResponse) {
	response = &UntagResourcesForExpressConnectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
