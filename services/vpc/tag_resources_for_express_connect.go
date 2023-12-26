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

// TagResourcesForExpressConnect invokes the vpc.TagResourcesForExpressConnect API synchronously
func (client *Client) TagResourcesForExpressConnect(request *TagResourcesForExpressConnectRequest) (response *TagResourcesForExpressConnectResponse, err error) {
	response = CreateTagResourcesForExpressConnectResponse()
	err = client.DoAction(request, response)
	return
}

// TagResourcesForExpressConnectWithChan invokes the vpc.TagResourcesForExpressConnect API asynchronously
func (client *Client) TagResourcesForExpressConnectWithChan(request *TagResourcesForExpressConnectRequest) (<-chan *TagResourcesForExpressConnectResponse, <-chan error) {
	responseChan := make(chan *TagResourcesForExpressConnectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TagResourcesForExpressConnect(request)
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

// TagResourcesForExpressConnectWithCallback invokes the vpc.TagResourcesForExpressConnect API asynchronously
func (client *Client) TagResourcesForExpressConnectWithCallback(request *TagResourcesForExpressConnectRequest, callback func(response *TagResourcesForExpressConnectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TagResourcesForExpressConnectResponse
		var err error
		defer close(result)
		response, err = client.TagResourcesForExpressConnect(request)
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

// TagResourcesForExpressConnectRequest is the request struct for api TagResourcesForExpressConnect
type TagResourcesForExpressConnectRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                    `position:"Query" name:"ResourceOwnerId"`
	Tag                  *[]TagResourcesForExpressConnectTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceId           *[]string                           `position:"Query" name:"ResourceId"  type:"Repeated"`
	ResourceOwnerAccount string                              `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                              `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                    `position:"Query" name:"OwnerId"`
	ResourceType         string                              `position:"Query" name:"ResourceType"`
}

// TagResourcesForExpressConnectTag is a repeated param struct in TagResourcesForExpressConnectRequest
type TagResourcesForExpressConnectTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// TagResourcesForExpressConnectResponse is the response struct for api TagResourcesForExpressConnect
type TagResourcesForExpressConnectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateTagResourcesForExpressConnectRequest creates a request to invoke TagResourcesForExpressConnect API
func CreateTagResourcesForExpressConnectRequest() (request *TagResourcesForExpressConnectRequest) {
	request = &TagResourcesForExpressConnectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "TagResourcesForExpressConnect", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTagResourcesForExpressConnectResponse creates a response to parse from TagResourcesForExpressConnect response
func CreateTagResourcesForExpressConnectResponse() (response *TagResourcesForExpressConnectResponse) {
	response = &TagResourcesForExpressConnectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
