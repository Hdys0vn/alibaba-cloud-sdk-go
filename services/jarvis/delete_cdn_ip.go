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

// DeleteCdnIp invokes the jarvis.DeleteCdnIp API synchronously
// api document: https://help.aliyun.com/api/jarvis/deletecdnip.html
func (client *Client) DeleteCdnIp(request *DeleteCdnIpRequest) (response *DeleteCdnIpResponse, err error) {
	response = CreateDeleteCdnIpResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCdnIpWithChan invokes the jarvis.DeleteCdnIp API asynchronously
// api document: https://help.aliyun.com/api/jarvis/deletecdnip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCdnIpWithChan(request *DeleteCdnIpRequest) (<-chan *DeleteCdnIpResponse, <-chan error) {
	responseChan := make(chan *DeleteCdnIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCdnIp(request)
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

// DeleteCdnIpWithCallback invokes the jarvis.DeleteCdnIp API asynchronously
// api document: https://help.aliyun.com/api/jarvis/deletecdnip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCdnIpWithCallback(request *DeleteCdnIpRequest, callback func(response *DeleteCdnIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCdnIpResponse
		var err error
		defer close(result)
		response, err = client.DeleteCdnIp(request)
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

// DeleteCdnIpRequest is the request struct for api DeleteCdnIp
type DeleteCdnIpRequest struct {
	*requests.RpcRequest
	ItemId          requests.Integer `position:"Query" name:"ItemId"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	CdnIp           string           `position:"Query" name:"CdnIp"`
	Lang            string           `position:"Query" name:"Lang"`
	SourceCode      string           `position:"Query" name:"SourceCode"`
}

// DeleteCdnIpResponse is the response struct for api DeleteCdnIp
type DeleteCdnIpResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
}

// CreateDeleteCdnIpRequest creates a request to invoke DeleteCdnIp API
func CreateDeleteCdnIpRequest() (request *DeleteCdnIpRequest) {
	request = &DeleteCdnIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DeleteCdnIp", "jarvis", "openAPI")
	return
}

// CreateDeleteCdnIpResponse creates a response to parse from DeleteCdnIp response
func CreateDeleteCdnIpResponse() (response *DeleteCdnIpResponse) {
	response = &DeleteCdnIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
