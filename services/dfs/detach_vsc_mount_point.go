package dfs

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

// DetachVscMountPoint invokes the dfs.DetachVscMountPoint API synchronously
func (client *Client) DetachVscMountPoint(request *DetachVscMountPointRequest) (response *DetachVscMountPointResponse, err error) {
	response = CreateDetachVscMountPointResponse()
	err = client.DoAction(request, response)
	return
}

// DetachVscMountPointWithChan invokes the dfs.DetachVscMountPoint API asynchronously
func (client *Client) DetachVscMountPointWithChan(request *DetachVscMountPointRequest) (<-chan *DetachVscMountPointResponse, <-chan error) {
	responseChan := make(chan *DetachVscMountPointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachVscMountPoint(request)
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

// DetachVscMountPointWithCallback invokes the dfs.DetachVscMountPoint API asynchronously
func (client *Client) DetachVscMountPointWithCallback(request *DetachVscMountPointRequest, callback func(response *DetachVscMountPointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachVscMountPointResponse
		var err error
		defer close(result)
		response, err = client.DetachVscMountPoint(request)
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

// DetachVscMountPointRequest is the request struct for api DetachVscMountPoint
type DetachVscMountPointRequest struct {
	*requests.RpcRequest
	Description   string    `position:"Query" name:"Description"`
	InputRegionId string    `position:"Query" name:"InputRegionId"`
	MountPointId  string    `position:"Query" name:"MountPointId"`
	FileSystemId  string    `position:"Query" name:"FileSystemId"`
	VscIds        *[]string `position:"Query" name:"VscIds"  type:"Json"`
	InstanceIds   string    `position:"Query" name:"InstanceIds"`
}

// DetachVscMountPointResponse is the response struct for api DetachVscMountPoint
type DetachVscMountPointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachVscMountPointRequest creates a request to invoke DetachVscMountPoint API
func CreateDetachVscMountPointRequest() (request *DetachVscMountPointRequest) {
	request = &DetachVscMountPointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DFS", "2018-06-20", "DetachVscMountPoint", "alidfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachVscMountPointResponse creates a response to parse from DetachVscMountPoint response
func CreateDetachVscMountPointResponse() (response *DetachVscMountPointResponse) {
	response = &DetachVscMountPointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
