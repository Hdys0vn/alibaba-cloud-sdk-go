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

// EnableVnInstance invokes the cloudcallcenter.EnableVnInstance API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/enablevninstance.html
func (client *Client) EnableVnInstance(request *EnableVnInstanceRequest) (response *EnableVnInstanceResponse, err error) {
	response = CreateEnableVnInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// EnableVnInstanceWithChan invokes the cloudcallcenter.EnableVnInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/enablevninstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableVnInstanceWithChan(request *EnableVnInstanceRequest) (<-chan *EnableVnInstanceResponse, <-chan error) {
	responseChan := make(chan *EnableVnInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableVnInstance(request)
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

// EnableVnInstanceWithCallback invokes the cloudcallcenter.EnableVnInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/enablevninstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableVnInstanceWithCallback(request *EnableVnInstanceRequest, callback func(response *EnableVnInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableVnInstanceResponse
		var err error
		defer close(result)
		response, err = client.EnableVnInstance(request)
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

// EnableVnInstanceRequest is the request struct for api EnableVnInstance
type EnableVnInstanceRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// EnableVnInstanceResponse is the response struct for api EnableVnInstance
type EnableVnInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateEnableVnInstanceRequest creates a request to invoke EnableVnInstance API
func CreateEnableVnInstanceRequest() (request *EnableVnInstanceRequest) {
	request = &EnableVnInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "EnableVnInstance", "", "")
	request.Method = requests.GET
	return
}

// CreateEnableVnInstanceResponse creates a response to parse from EnableVnInstance response
func CreateEnableVnInstanceResponse() (response *EnableVnInstanceResponse) {
	response = &EnableVnInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}