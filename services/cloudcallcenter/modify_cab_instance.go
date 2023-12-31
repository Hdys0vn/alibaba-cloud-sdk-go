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

// ModifyCabInstance invokes the cloudcallcenter.ModifyCabInstance API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifycabinstance.html
func (client *Client) ModifyCabInstance(request *ModifyCabInstanceRequest) (response *ModifyCabInstanceResponse, err error) {
	response = CreateModifyCabInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCabInstanceWithChan invokes the cloudcallcenter.ModifyCabInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifycabinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCabInstanceWithChan(request *ModifyCabInstanceRequest) (<-chan *ModifyCabInstanceResponse, <-chan error) {
	responseChan := make(chan *ModifyCabInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCabInstance(request)
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

// ModifyCabInstanceWithCallback invokes the cloudcallcenter.ModifyCabInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifycabinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCabInstanceWithCallback(request *ModifyCabInstanceRequest, callback func(response *ModifyCabInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCabInstanceResponse
		var err error
		defer close(result)
		response, err = client.ModifyCabInstance(request)
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

// ModifyCabInstanceRequest is the request struct for api ModifyCabInstance
type ModifyCabInstanceRequest struct {
	*requests.RpcRequest
	MaxConcurrentConversation requests.Integer `position:"Query" name:"MaxConcurrentConversation"`
	InstanceId                string           `position:"Query" name:"InstanceId"`
	InstanceName              string           `position:"Query" name:"InstanceName"`
	CallCenterInstanceId      string           `position:"Query" name:"CallCenterInstanceId"`
	InstanceDescription       string           `position:"Query" name:"InstanceDescription"`
}

// ModifyCabInstanceResponse is the response struct for api ModifyCabInstance
type ModifyCabInstanceResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Instance       Instance `json:"Instance" xml:"Instance"`
}

// CreateModifyCabInstanceRequest creates a request to invoke ModifyCabInstance API
func CreateModifyCabInstanceRequest() (request *ModifyCabInstanceRequest) {
	request = &ModifyCabInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyCabInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyCabInstanceResponse creates a response to parse from ModifyCabInstance response
func CreateModifyCabInstanceResponse() (response *ModifyCabInstanceResponse) {
	response = &ModifyCabInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
