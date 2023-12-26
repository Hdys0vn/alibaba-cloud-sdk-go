package dms_enterprise

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

// PublishAndDeployTaskFlow invokes the dms_enterprise.PublishAndDeployTaskFlow API synchronously
func (client *Client) PublishAndDeployTaskFlow(request *PublishAndDeployTaskFlowRequest) (response *PublishAndDeployTaskFlowResponse, err error) {
	response = CreatePublishAndDeployTaskFlowResponse()
	err = client.DoAction(request, response)
	return
}

// PublishAndDeployTaskFlowWithChan invokes the dms_enterprise.PublishAndDeployTaskFlow API asynchronously
func (client *Client) PublishAndDeployTaskFlowWithChan(request *PublishAndDeployTaskFlowRequest) (<-chan *PublishAndDeployTaskFlowResponse, <-chan error) {
	responseChan := make(chan *PublishAndDeployTaskFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PublishAndDeployTaskFlow(request)
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

// PublishAndDeployTaskFlowWithCallback invokes the dms_enterprise.PublishAndDeployTaskFlow API asynchronously
func (client *Client) PublishAndDeployTaskFlowWithCallback(request *PublishAndDeployTaskFlowRequest, callback func(response *PublishAndDeployTaskFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PublishAndDeployTaskFlowResponse
		var err error
		defer close(result)
		response, err = client.PublishAndDeployTaskFlow(request)
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

// PublishAndDeployTaskFlowRequest is the request struct for api PublishAndDeployTaskFlow
type PublishAndDeployTaskFlowRequest struct {
	*requests.RpcRequest
	DagId           requests.Integer `position:"Query" name:"DagId"`
	Tid             requests.Integer `position:"Query" name:"Tid"`
	VersionComments string           `position:"Query" name:"VersionComments"`
}

// PublishAndDeployTaskFlowResponse is the response struct for api PublishAndDeployTaskFlow
type PublishAndDeployTaskFlowResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	DeployId     int64  `json:"DeployId" xml:"DeployId"`
}

// CreatePublishAndDeployTaskFlowRequest creates a request to invoke PublishAndDeployTaskFlow API
func CreatePublishAndDeployTaskFlowRequest() (request *PublishAndDeployTaskFlowRequest) {
	request = &PublishAndDeployTaskFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "PublishAndDeployTaskFlow", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePublishAndDeployTaskFlowResponse creates a response to parse from PublishAndDeployTaskFlow response
func CreatePublishAndDeployTaskFlowResponse() (response *PublishAndDeployTaskFlowResponse) {
	response = &PublishAndDeployTaskFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}