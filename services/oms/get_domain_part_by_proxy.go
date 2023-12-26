package oms

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

// GetDomainPartByProxy invokes the oms.GetDomainPartByProxy API synchronously
func (client *Client) GetDomainPartByProxy(request *GetDomainPartByProxyRequest) (response *GetDomainPartByProxyResponse, err error) {
	response = CreateGetDomainPartByProxyResponse()
	err = client.DoAction(request, response)
	return
}

// GetDomainPartByProxyWithChan invokes the oms.GetDomainPartByProxy API asynchronously
func (client *Client) GetDomainPartByProxyWithChan(request *GetDomainPartByProxyRequest) (<-chan *GetDomainPartByProxyResponse, <-chan error) {
	responseChan := make(chan *GetDomainPartByProxyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDomainPartByProxy(request)
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

// GetDomainPartByProxyWithCallback invokes the oms.GetDomainPartByProxy API asynchronously
func (client *Client) GetDomainPartByProxyWithCallback(request *GetDomainPartByProxyRequest, callback func(response *GetDomainPartByProxyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDomainPartByProxyResponse
		var err error
		defer close(result)
		response, err = client.GetDomainPartByProxy(request)
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

// GetDomainPartByProxyRequest is the request struct for api GetDomainPartByProxy
type GetDomainPartByProxyRequest struct {
	*requests.RpcRequest
	DomainCode     string           `position:"Query" name:"DomainCode"`
	DataType       string           `position:"Query" name:"DataType"`
	CompressEnable requests.Boolean `position:"Query" name:"CompressEnable"`
	Part           string           `position:"Query" name:"Part"`
}

// GetDomainPartByProxyResponse is the response struct for api GetDomainPartByProxy
type GetDomainPartByProxyResponse struct {
	*responses.BaseResponse
	Compressed bool   `json:"Compressed" xml:"Compressed"`
	Data       string `json:"Data" xml:"Data"`
	DataType   string `json:"DataType" xml:"DataType"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	DomainCode string `json:"DomainCode" xml:"DomainCode"`
}

// CreateGetDomainPartByProxyRequest creates a request to invoke GetDomainPartByProxy API
func CreateGetDomainPartByProxyRequest() (request *GetDomainPartByProxyRequest) {
	request = &GetDomainPartByProxyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oms", "2016-06-15", "GetDomainPartByProxy", "", "")
	request.Method = requests.POST
	return
}

// CreateGetDomainPartByProxyResponse creates a response to parse from GetDomainPartByProxy response
func CreateGetDomainPartByProxyResponse() (response *GetDomainPartByProxyResponse) {
	response = &GetDomainPartByProxyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
