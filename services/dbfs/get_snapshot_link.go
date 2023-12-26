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

// GetSnapshotLink invokes the dbfs.GetSnapshotLink API synchronously
func (client *Client) GetSnapshotLink(request *GetSnapshotLinkRequest) (response *GetSnapshotLinkResponse, err error) {
	response = CreateGetSnapshotLinkResponse()
	err = client.DoAction(request, response)
	return
}

// GetSnapshotLinkWithChan invokes the dbfs.GetSnapshotLink API asynchronously
func (client *Client) GetSnapshotLinkWithChan(request *GetSnapshotLinkRequest) (<-chan *GetSnapshotLinkResponse, <-chan error) {
	responseChan := make(chan *GetSnapshotLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSnapshotLink(request)
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

// GetSnapshotLinkWithCallback invokes the dbfs.GetSnapshotLink API asynchronously
func (client *Client) GetSnapshotLinkWithCallback(request *GetSnapshotLinkRequest, callback func(response *GetSnapshotLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSnapshotLinkResponse
		var err error
		defer close(result)
		response, err = client.GetSnapshotLink(request)
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

// GetSnapshotLinkRequest is the request struct for api GetSnapshotLink
type GetSnapshotLinkRequest struct {
	*requests.RpcRequest
	LinkId string `position:"Query" name:"LinkId"`
}

// GetSnapshotLinkResponse is the response struct for api GetSnapshotLink
type GetSnapshotLinkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetSnapshotLinkRequest creates a request to invoke GetSnapshotLink API
func CreateGetSnapshotLinkRequest() (request *GetSnapshotLinkRequest) {
	request = &GetSnapshotLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "GetSnapshotLink", "dbfs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetSnapshotLinkResponse creates a response to parse from GetSnapshotLink response
func CreateGetSnapshotLinkResponse() (response *GetSnapshotLinkResponse) {
	response = &GetSnapshotLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}