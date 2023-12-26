package dbfs

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

// GetAutoSnapshotPolicy invokes the dbfs.GetAutoSnapshotPolicy API synchronously
func (client *Client) GetAutoSnapshotPolicy(request *GetAutoSnapshotPolicyRequest) (response *GetAutoSnapshotPolicyResponse, err error) {
	response = CreateGetAutoSnapshotPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// GetAutoSnapshotPolicyWithChan invokes the dbfs.GetAutoSnapshotPolicy API asynchronously
func (client *Client) GetAutoSnapshotPolicyWithChan(request *GetAutoSnapshotPolicyRequest) (<-chan *GetAutoSnapshotPolicyResponse, <-chan error) {
	responseChan := make(chan *GetAutoSnapshotPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAutoSnapshotPolicy(request)
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

// GetAutoSnapshotPolicyWithCallback invokes the dbfs.GetAutoSnapshotPolicy API asynchronously
func (client *Client) GetAutoSnapshotPolicyWithCallback(request *GetAutoSnapshotPolicyRequest, callback func(response *GetAutoSnapshotPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAutoSnapshotPolicyResponse
		var err error
		defer close(result)
		response, err = client.GetAutoSnapshotPolicy(request)
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

// GetAutoSnapshotPolicyRequest is the request struct for api GetAutoSnapshotPolicy
type GetAutoSnapshotPolicyRequest struct {
	*requests.RpcRequest
	PolicyId string `position:"Query" name:"PolicyId"`
}

// GetAutoSnapshotPolicyResponse is the response struct for api GetAutoSnapshotPolicy
type GetAutoSnapshotPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetAutoSnapshotPolicyRequest creates a request to invoke GetAutoSnapshotPolicy API
func CreateGetAutoSnapshotPolicyRequest() (request *GetAutoSnapshotPolicyRequest) {
	request = &GetAutoSnapshotPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "GetAutoSnapshotPolicy", "dbfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAutoSnapshotPolicyResponse creates a response to parse from GetAutoSnapshotPolicy response
func CreateGetAutoSnapshotPolicyResponse() (response *GetAutoSnapshotPolicyResponse) {
	response = &GetAutoSnapshotPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}