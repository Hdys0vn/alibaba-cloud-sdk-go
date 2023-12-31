package cassandra

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

// ModifyDataCenter invokes the cassandra.ModifyDataCenter API synchronously
func (client *Client) ModifyDataCenter(request *ModifyDataCenterRequest) (response *ModifyDataCenterResponse, err error) {
	response = CreateModifyDataCenterResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDataCenterWithChan invokes the cassandra.ModifyDataCenter API asynchronously
func (client *Client) ModifyDataCenterWithChan(request *ModifyDataCenterRequest) (<-chan *ModifyDataCenterResponse, <-chan error) {
	responseChan := make(chan *ModifyDataCenterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDataCenter(request)
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

// ModifyDataCenterWithCallback invokes the cassandra.ModifyDataCenter API asynchronously
func (client *Client) ModifyDataCenterWithCallback(request *ModifyDataCenterRequest, callback func(response *ModifyDataCenterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDataCenterResponse
		var err error
		defer close(result)
		response, err = client.ModifyDataCenter(request)
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

// ModifyDataCenterRequest is the request struct for api ModifyDataCenter
type ModifyDataCenterRequest struct {
	*requests.RpcRequest
	DataCenterId   string `position:"Query" name:"DataCenterId"`
	ClusterId      string `position:"Query" name:"ClusterId"`
	DataCenterName string `position:"Query" name:"DataCenterName"`
}

// ModifyDataCenterResponse is the response struct for api ModifyDataCenter
type ModifyDataCenterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDataCenterRequest creates a request to invoke ModifyDataCenter API
func CreateModifyDataCenterRequest() (request *ModifyDataCenterRequest) {
	request = &ModifyDataCenterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "ModifyDataCenter", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDataCenterResponse creates a response to parse from ModifyDataCenter response
func CreateModifyDataCenterResponse() (response *ModifyDataCenterResponse) {
	response = &ModifyDataCenterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
