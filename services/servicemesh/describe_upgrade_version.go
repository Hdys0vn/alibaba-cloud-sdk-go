package servicemesh

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

// DescribeUpgradeVersion invokes the servicemesh.DescribeUpgradeVersion API synchronously
func (client *Client) DescribeUpgradeVersion(request *DescribeUpgradeVersionRequest) (response *DescribeUpgradeVersionResponse, err error) {
	response = CreateDescribeUpgradeVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUpgradeVersionWithChan invokes the servicemesh.DescribeUpgradeVersion API asynchronously
func (client *Client) DescribeUpgradeVersionWithChan(request *DescribeUpgradeVersionRequest) (<-chan *DescribeUpgradeVersionResponse, <-chan error) {
	responseChan := make(chan *DescribeUpgradeVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUpgradeVersion(request)
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

// DescribeUpgradeVersionWithCallback invokes the servicemesh.DescribeUpgradeVersion API asynchronously
func (client *Client) DescribeUpgradeVersionWithCallback(request *DescribeUpgradeVersionRequest, callback func(response *DescribeUpgradeVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUpgradeVersionResponse
		var err error
		defer close(result)
		response, err = client.DescribeUpgradeVersion(request)
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

// DescribeUpgradeVersionRequest is the request struct for api DescribeUpgradeVersion
type DescribeUpgradeVersionRequest struct {
	*requests.RpcRequest
	ServiceMeshId string `position:"Query" name:"ServiceMeshId"`
}

// DescribeUpgradeVersionResponse is the response struct for api DescribeUpgradeVersion
type DescribeUpgradeVersionResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Version   Version `json:"Version" xml:"Version"`
}

// CreateDescribeUpgradeVersionRequest creates a request to invoke DescribeUpgradeVersion API
func CreateDescribeUpgradeVersionRequest() (request *DescribeUpgradeVersionRequest) {
	request = &DescribeUpgradeVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("servicemesh", "2020-01-11", "DescribeUpgradeVersion", "servicemesh", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUpgradeVersionResponse creates a response to parse from DescribeUpgradeVersion response
func CreateDescribeUpgradeVersionResponse() (response *DescribeUpgradeVersionResponse) {
	response = &DescribeUpgradeVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
