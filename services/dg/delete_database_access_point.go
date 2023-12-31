package dg

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

// DeleteDatabaseAccessPoint invokes the dg.DeleteDatabaseAccessPoint API synchronously
func (client *Client) DeleteDatabaseAccessPoint(request *DeleteDatabaseAccessPointRequest) (response *DeleteDatabaseAccessPointResponse, err error) {
	response = CreateDeleteDatabaseAccessPointResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDatabaseAccessPointWithChan invokes the dg.DeleteDatabaseAccessPoint API asynchronously
func (client *Client) DeleteDatabaseAccessPointWithChan(request *DeleteDatabaseAccessPointRequest) (<-chan *DeleteDatabaseAccessPointResponse, <-chan error) {
	responseChan := make(chan *DeleteDatabaseAccessPointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDatabaseAccessPoint(request)
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

// DeleteDatabaseAccessPointWithCallback invokes the dg.DeleteDatabaseAccessPoint API asynchronously
func (client *Client) DeleteDatabaseAccessPointWithCallback(request *DeleteDatabaseAccessPointRequest, callback func(response *DeleteDatabaseAccessPointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDatabaseAccessPointResponse
		var err error
		defer close(result)
		response, err = client.DeleteDatabaseAccessPoint(request)
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

// DeleteDatabaseAccessPointRequest is the request struct for api DeleteDatabaseAccessPoint
type DeleteDatabaseAccessPointRequest struct {
	*requests.RpcRequest
	VpcAZone     string `position:"Body" name:"VpcAZone"`
	VpcRegionId  string `position:"Body" name:"VpcRegionId"`
	VSwitchId    string `position:"Body" name:"VSwitchId"`
	VpcId        string `position:"Body" name:"VpcId"`
	DbInstanceId string `position:"Body" name:"DbInstanceId"`
}

// DeleteDatabaseAccessPointResponse is the response struct for api DeleteDatabaseAccessPoint
type DeleteDatabaseAccessPointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateDeleteDatabaseAccessPointRequest creates a request to invoke DeleteDatabaseAccessPoint API
func CreateDeleteDatabaseAccessPointRequest() (request *DeleteDatabaseAccessPointRequest) {
	request = &DeleteDatabaseAccessPointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dg", "2019-03-27", "DeleteDatabaseAccessPoint", "dg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDatabaseAccessPointResponse creates a response to parse from DeleteDatabaseAccessPoint response
func CreateDeleteDatabaseAccessPointResponse() (response *DeleteDatabaseAccessPointResponse) {
	response = &DeleteDatabaseAccessPointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
