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

// GetContactFlowVersion invokes the cloudcallcenter.GetContactFlowVersion API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcontactflowversion.html
func (client *Client) GetContactFlowVersion(request *GetContactFlowVersionRequest) (response *GetContactFlowVersionResponse, err error) {
	response = CreateGetContactFlowVersionResponse()
	err = client.DoAction(request, response)
	return
}

// GetContactFlowVersionWithChan invokes the cloudcallcenter.GetContactFlowVersion API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcontactflowversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetContactFlowVersionWithChan(request *GetContactFlowVersionRequest) (<-chan *GetContactFlowVersionResponse, <-chan error) {
	responseChan := make(chan *GetContactFlowVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetContactFlowVersion(request)
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

// GetContactFlowVersionWithCallback invokes the cloudcallcenter.GetContactFlowVersion API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getcontactflowversion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetContactFlowVersionWithCallback(request *GetContactFlowVersionRequest, callback func(response *GetContactFlowVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetContactFlowVersionResponse
		var err error
		defer close(result)
		response, err = client.GetContactFlowVersion(request)
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

// GetContactFlowVersionRequest is the request struct for api GetContactFlowVersion
type GetContactFlowVersionRequest struct {
	*requests.RpcRequest
	InstanceId           string `position:"Query" name:"InstanceId"`
	ContactFlowVersionId string `position:"Query" name:"ContactFlowVersionId"`
}

// GetContactFlowVersionResponse is the response struct for api GetContactFlowVersion
type GetContactFlowVersionResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Code           string      `json:"Code" xml:"Code"`
	Message        string      `json:"Message" xml:"Message"`
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ContactFlow    ContactFlow `json:"ContactFlow" xml:"ContactFlow"`
}

// CreateGetContactFlowVersionRequest creates a request to invoke GetContactFlowVersion API
func CreateGetContactFlowVersionRequest() (request *GetContactFlowVersionRequest) {
	request = &GetContactFlowVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetContactFlowVersion", "", "")
	request.Method = requests.POST
	return
}

// CreateGetContactFlowVersionResponse creates a response to parse from GetContactFlowVersion response
func CreateGetContactFlowVersionResponse() (response *GetContactFlowVersionResponse) {
	response = &GetContactFlowVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
