package jarvis

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

// CreateCdnIp invokes the jarvis.CreateCdnIp API synchronously
// api document: https://help.aliyun.com/api/jarvis/createcdnip.html
func (client *Client) CreateCdnIp(request *CreateCdnIpRequest) (response *CreateCdnIpResponse, err error) {
	response = CreateCreateCdnIpResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCdnIpWithChan invokes the jarvis.CreateCdnIp API asynchronously
// api document: https://help.aliyun.com/api/jarvis/createcdnip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCdnIpWithChan(request *CreateCdnIpRequest) (<-chan *CreateCdnIpResponse, <-chan error) {
	responseChan := make(chan *CreateCdnIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCdnIp(request)
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

// CreateCdnIpWithCallback invokes the jarvis.CreateCdnIp API asynchronously
// api document: https://help.aliyun.com/api/jarvis/createcdnip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCdnIpWithCallback(request *CreateCdnIpRequest, callback func(response *CreateCdnIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCdnIpResponse
		var err error
		defer close(result)
		response, err = client.CreateCdnIp(request)
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

// CreateCdnIpRequest is the request struct for api CreateCdnIp
type CreateCdnIpRequest struct {
	*requests.RpcRequest
	CdnIpList       string           `position:"Query" name:"CdnIpList"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
	SourceCode      string           `position:"Query" name:"SourceCode"`
}

// CreateCdnIpResponse is the response struct for api CreateCdnIp
type CreateCdnIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
}

// CreateCreateCdnIpRequest creates a request to invoke CreateCdnIp API
func CreateCreateCdnIpRequest() (request *CreateCdnIpRequest) {
	request = &CreateCdnIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "CreateCdnIp", "jarvis", "openAPI")
	return
}

// CreateCreateCdnIpResponse creates a response to parse from CreateCdnIp response
func CreateCreateCdnIpResponse() (response *CreateCdnIpResponse) {
	response = &CreateCdnIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
