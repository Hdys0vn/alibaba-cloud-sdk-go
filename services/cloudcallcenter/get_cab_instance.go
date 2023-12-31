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

// GetCabInstance invokes the cloudcallcenter.GetCabInstance API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcabinstance.html
func (client *Client) GetCabInstance(request *GetCabInstanceRequest) (response *GetCabInstanceResponse, err error) {
	response = CreateGetCabInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetCabInstanceWithChan invokes the cloudcallcenter.GetCabInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcabinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCabInstanceWithChan(request *GetCabInstanceRequest) (<-chan *GetCabInstanceResponse, <-chan error) {
	responseChan := make(chan *GetCabInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCabInstance(request)
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

// GetCabInstanceWithCallback invokes the cloudcallcenter.GetCabInstance API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcabinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCabInstanceWithCallback(request *GetCabInstanceRequest, callback func(response *GetCabInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCabInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetCabInstance(request)
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

// GetCabInstanceRequest is the request struct for api GetCabInstance
type GetCabInstanceRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetCabInstanceResponse is the response struct for api GetCabInstance
type GetCabInstanceResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Instance       Instance `json:"Instance" xml:"Instance"`
}

// CreateGetCabInstanceRequest creates a request to invoke GetCabInstance API
func CreateGetCabInstanceRequest() (request *GetCabInstanceRequest) {
	request = &GetCabInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetCabInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateGetCabInstanceResponse creates a response to parse from GetCabInstance response
func CreateGetCabInstanceResponse() (response *GetCabInstanceResponse) {
	response = &GetCabInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
