package smartag

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

// DeleteCloudConnectNetwork invokes the smartag.DeleteCloudConnectNetwork API synchronously
func (client *Client) DeleteCloudConnectNetwork(request *DeleteCloudConnectNetworkRequest) (response *DeleteCloudConnectNetworkResponse, err error) {
	response = CreateDeleteCloudConnectNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCloudConnectNetworkWithChan invokes the smartag.DeleteCloudConnectNetwork API asynchronously
func (client *Client) DeleteCloudConnectNetworkWithChan(request *DeleteCloudConnectNetworkRequest) (<-chan *DeleteCloudConnectNetworkResponse, <-chan error) {
	responseChan := make(chan *DeleteCloudConnectNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCloudConnectNetwork(request)
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

// DeleteCloudConnectNetworkWithCallback invokes the smartag.DeleteCloudConnectNetwork API asynchronously
func (client *Client) DeleteCloudConnectNetworkWithCallback(request *DeleteCloudConnectNetworkRequest, callback func(response *DeleteCloudConnectNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCloudConnectNetworkResponse
		var err error
		defer close(result)
		response, err = client.DeleteCloudConnectNetwork(request)
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

// DeleteCloudConnectNetworkRequest is the request struct for api DeleteCloudConnectNetwork
type DeleteCloudConnectNetworkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CcnId                string           `position:"Query" name:"CcnId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteCloudConnectNetworkResponse is the response struct for api DeleteCloudConnectNetwork
type DeleteCloudConnectNetworkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCloudConnectNetworkRequest creates a request to invoke DeleteCloudConnectNetwork API
func CreateDeleteCloudConnectNetworkRequest() (request *DeleteCloudConnectNetworkRequest) {
	request = &DeleteCloudConnectNetworkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DeleteCloudConnectNetwork", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCloudConnectNetworkResponse creates a response to parse from DeleteCloudConnectNetwork response
func CreateDeleteCloudConnectNetworkResponse() (response *DeleteCloudConnectNetworkResponse) {
	response = &DeleteCloudConnectNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
