package sgw

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

// CreateExpressSync invokes the sgw.CreateExpressSync API synchronously
func (client *Client) CreateExpressSync(request *CreateExpressSyncRequest) (response *CreateExpressSyncResponse, err error) {
	response = CreateCreateExpressSyncResponse()
	err = client.DoAction(request, response)
	return
}

// CreateExpressSyncWithChan invokes the sgw.CreateExpressSync API asynchronously
func (client *Client) CreateExpressSyncWithChan(request *CreateExpressSyncRequest) (<-chan *CreateExpressSyncResponse, <-chan error) {
	responseChan := make(chan *CreateExpressSyncResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateExpressSync(request)
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

// CreateExpressSyncWithCallback invokes the sgw.CreateExpressSync API asynchronously
func (client *Client) CreateExpressSyncWithCallback(request *CreateExpressSyncRequest, callback func(response *CreateExpressSyncResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateExpressSyncResponse
		var err error
		defer close(result)
		response, err = client.CreateExpressSync(request)
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

// CreateExpressSyncRequest is the request struct for api CreateExpressSync
type CreateExpressSyncRequest struct {
	*requests.RpcRequest
	BucketRegion  string `position:"Query" name:"BucketRegion"`
	Description   string `position:"Query" name:"Description"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	BucketName    string `position:"Query" name:"BucketName"`
	Name          string `position:"Query" name:"Name"`
	BucketPrefix  string `position:"Query" name:"BucketPrefix"`
}

// CreateExpressSyncResponse is the response struct for api CreateExpressSync
type CreateExpressSyncResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Success       bool   `json:"Success" xml:"Success"`
	Code          string `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ExpressSyncId string `json:"ExpressSyncId" xml:"ExpressSyncId"`
}

// CreateCreateExpressSyncRequest creates a request to invoke CreateExpressSync API
func CreateCreateExpressSyncRequest() (request *CreateExpressSyncRequest) {
	request = &CreateExpressSyncRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "CreateExpressSync", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateExpressSyncResponse creates a response to parse from CreateExpressSync response
func CreateCreateExpressSyncResponse() (response *CreateExpressSyncResponse) {
	response = &CreateExpressSyncResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}