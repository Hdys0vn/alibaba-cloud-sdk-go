package cloudcallcenter

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

// RollbackVnInstance invokes the cloudcallcenter.RollbackVnInstance API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/rollbackvninstance.html
func (client *Client) RollbackVnInstance(request *RollbackVnInstanceRequest) (response *RollbackVnInstanceResponse, err error) {
	response = CreateRollbackVnInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RollbackVnInstanceWithChan invokes the cloudcallcenter.RollbackVnInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/rollbackvninstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RollbackVnInstanceWithChan(request *RollbackVnInstanceRequest) (<-chan *RollbackVnInstanceResponse, <-chan error) {
	responseChan := make(chan *RollbackVnInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RollbackVnInstance(request)
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

// RollbackVnInstanceWithCallback invokes the cloudcallcenter.RollbackVnInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/rollbackvninstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RollbackVnInstanceWithCallback(request *RollbackVnInstanceRequest, callback func(response *RollbackVnInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RollbackVnInstanceResponse
		var err error
		defer close(result)
		response, err = client.RollbackVnInstance(request)
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

// RollbackVnInstanceRequest is the request struct for api RollbackVnInstance
type RollbackVnInstanceRequest struct {
	*requests.RpcRequest
	TargetVersion string `position:"Query" name:"TargetVersion"`
	InstanceId    string `position:"Query" name:"InstanceId"`
}

// RollbackVnInstanceResponse is the response struct for api RollbackVnInstance
type RollbackVnInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateRollbackVnInstanceRequest creates a request to invoke RollbackVnInstance API
func CreateRollbackVnInstanceRequest() (request *RollbackVnInstanceRequest) {
	request = &RollbackVnInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "RollbackVnInstance", "", "")
	request.Method = requests.GET
	return
}

// CreateRollbackVnInstanceResponse creates a response to parse from RollbackVnInstance response
func CreateRollbackVnInstanceResponse() (response *RollbackVnInstanceResponse) {
	response = &RollbackVnInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}