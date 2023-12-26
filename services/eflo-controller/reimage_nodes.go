package eflo_controller

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

// ReimageNodes invokes the eflo_controller.ReimageNodes API synchronously
func (client *Client) ReimageNodes(request *ReimageNodesRequest) (response *ReimageNodesResponse, err error) {
	response = CreateReimageNodesResponse()
	err = client.DoAction(request, response)
	return
}

// ReimageNodesWithChan invokes the eflo_controller.ReimageNodes API asynchronously
func (client *Client) ReimageNodesWithChan(request *ReimageNodesRequest) (<-chan *ReimageNodesResponse, <-chan error) {
	responseChan := make(chan *ReimageNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReimageNodes(request)
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

// ReimageNodesWithCallback invokes the eflo_controller.ReimageNodes API asynchronously
func (client *Client) ReimageNodesWithCallback(request *ReimageNodesRequest, callback func(response *ReimageNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReimageNodesResponse
		var err error
		defer close(result)
		response, err = client.ReimageNodes(request)
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

// ReimageNodesRequest is the request struct for api ReimageNodes
type ReimageNodesRequest struct {
	*requests.RpcRequest
	IgnoreFailedNodeTasks requests.Boolean     `position:"Body" name:"IgnoreFailedNodeTasks"`
	ClusterId             string               `position:"Body" name:"ClusterId"`
	Nodes                 *[]ReimageNodesNodes `position:"Body" name:"Nodes"  type:"Json"`
}

// ReimageNodesNodes is a repeated param struct in ReimageNodesRequest
type ReimageNodesNodes struct {
	Hostname      string `name:"Hostname"`
	ImageId       string `name:"ImageId"`
	LoginPassword string `name:"LoginPassword"`
	NodeId        string `name:"NodeId"`
}

// ReimageNodesResponse is the response struct for api ReimageNodes
type ReimageNodesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateReimageNodesRequest creates a request to invoke ReimageNodes API
func CreateReimageNodesRequest() (request *ReimageNodesRequest) {
	request = &ReimageNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo-controller", "2022-12-15", "ReimageNodes", "", "")
	request.Method = requests.POST
	return
}

// CreateReimageNodesResponse creates a response to parse from ReimageNodes response
func CreateReimageNodesResponse() (response *ReimageNodesResponse) {
	response = &ReimageNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
