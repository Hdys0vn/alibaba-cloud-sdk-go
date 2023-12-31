package csb

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

// GetMetaServerSLO invokes the csb.GetMetaServerSLO API synchronously
// api document: https://help.aliyun.com/api/csb/getmetaserverslo.html
func (client *Client) GetMetaServerSLO(request *GetMetaServerSLORequest) (response *GetMetaServerSLOResponse, err error) {
	response = CreateGetMetaServerSLOResponse()
	err = client.DoAction(request, response)
	return
}

// GetMetaServerSLOWithChan invokes the csb.GetMetaServerSLO API asynchronously
// api document: https://help.aliyun.com/api/csb/getmetaserverslo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMetaServerSLOWithChan(request *GetMetaServerSLORequest) (<-chan *GetMetaServerSLOResponse, <-chan error) {
	responseChan := make(chan *GetMetaServerSLOResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMetaServerSLO(request)
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

// GetMetaServerSLOWithCallback invokes the csb.GetMetaServerSLO API asynchronously
// api document: https://help.aliyun.com/api/csb/getmetaserverslo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMetaServerSLOWithCallback(request *GetMetaServerSLORequest, callback func(response *GetMetaServerSLOResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMetaServerSLOResponse
		var err error
		defer close(result)
		response, err = client.GetMetaServerSLO(request)
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

// GetMetaServerSLORequest is the request struct for api GetMetaServerSLO
type GetMetaServerSLORequest struct {
	*requests.RpcRequest
}

// GetMetaServerSLOResponse is the response struct for api GetMetaServerSLO
type GetMetaServerSLOResponse struct {
	*responses.BaseResponse
	Code      int                    `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Data      map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateGetMetaServerSLORequest creates a request to invoke GetMetaServerSLO API
func CreateGetMetaServerSLORequest() (request *GetMetaServerSLORequest) {
	request = &GetMetaServerSLORequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "GetMetaServerSLO", "", "")
	request.Method = requests.GET
	return
}

// CreateGetMetaServerSLOResponse creates a response to parse from GetMetaServerSLO response
func CreateGetMetaServerSLOResponse() (response *GetMetaServerSLOResponse) {
	response = &GetMetaServerSLOResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
