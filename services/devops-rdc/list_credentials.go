package devops_rdc

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

// ListCredentials invokes the devops_rdc.ListCredentials API synchronously
func (client *Client) ListCredentials(request *ListCredentialsRequest) (response *ListCredentialsResponse, err error) {
	response = CreateListCredentialsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCredentialsWithChan invokes the devops_rdc.ListCredentials API asynchronously
func (client *Client) ListCredentialsWithChan(request *ListCredentialsRequest) (<-chan *ListCredentialsResponse, <-chan error) {
	responseChan := make(chan *ListCredentialsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCredentials(request)
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

// ListCredentialsWithCallback invokes the devops_rdc.ListCredentials API asynchronously
func (client *Client) ListCredentialsWithCallback(request *ListCredentialsRequest, callback func(response *ListCredentialsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCredentialsResponse
		var err error
		defer close(result)
		response, err = client.ListCredentials(request)
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

// ListCredentialsRequest is the request struct for api ListCredentials
type ListCredentialsRequest struct {
	*requests.RpcRequest
	UserPk string `position:"Body" name:"UserPk"`
	OrgId  string `position:"Body" name:"OrgId"`
}

// ListCredentialsResponse is the response struct for api ListCredentials
type ListCredentialsResponse struct {
	*responses.BaseResponse
	RequestId    string                   `json:"RequestId" xml:"RequestId"`
	ErrorCode    string                   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                   `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool                     `json:"Success" xml:"Success"`
	Object       []map[string]interface{} `json:"Object" xml:"Object"`
}

// CreateListCredentialsRequest creates a request to invoke ListCredentials API
func CreateListCredentialsRequest() (request *ListCredentialsRequest) {
	request = &ListCredentialsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "ListCredentials", "", "")
	request.Method = requests.POST
	return
}

// CreateListCredentialsResponse creates a response to parse from ListCredentials response
func CreateListCredentialsResponse() (response *ListCredentialsResponse) {
	response = &ListCredentialsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
