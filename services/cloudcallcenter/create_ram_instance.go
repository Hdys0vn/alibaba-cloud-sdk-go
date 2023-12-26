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

// CreateRamInstance invokes the cloudcallcenter.CreateRamInstance API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createraminstance.html
func (client *Client) CreateRamInstance(request *CreateRamInstanceRequest) (response *CreateRamInstanceResponse, err error) {
	response = CreateCreateRamInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRamInstanceWithChan invokes the cloudcallcenter.CreateRamInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createraminstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRamInstanceWithChan(request *CreateRamInstanceRequest) (<-chan *CreateRamInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateRamInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRamInstance(request)
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

// CreateRamInstanceWithCallback invokes the cloudcallcenter.CreateRamInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createraminstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateRamInstanceWithCallback(request *CreateRamInstanceRequest, callback func(response *CreateRamInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRamInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateRamInstance(request)
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

// CreateRamInstanceRequest is the request struct for api CreateRamInstance
type CreateRamInstanceRequest struct {
	*requests.RpcRequest
	ProductCode      string `position:"Query" name:"ProductCode"`
	Role             string `position:"Query" name:"Role"`
	SubscriptionType string `position:"Query" name:"SubscriptionType"`
}

// CreateRamInstanceResponse is the response struct for api CreateRamInstance
type CreateRamInstanceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCreateRamInstanceRequest creates a request to invoke CreateRamInstance API
func CreateCreateRamInstanceRequest() (request *CreateRamInstanceRequest) {
	request = &CreateRamInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateRamInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateRamInstanceResponse creates a response to parse from CreateRamInstance response
func CreateCreateRamInstanceResponse() (response *CreateRamInstanceResponse) {
	response = &CreateRamInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
