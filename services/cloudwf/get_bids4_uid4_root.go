package cloudwf

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

// GetBids4Uid4Root invokes the cloudwf.GetBids4Uid4Root API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getbids4uid4root.html
func (client *Client) GetBids4Uid4Root(request *GetBids4Uid4RootRequest) (response *GetBids4Uid4RootResponse, err error) {
	response = CreateGetBids4Uid4RootResponse()
	err = client.DoAction(request, response)
	return
}

// GetBids4Uid4RootWithChan invokes the cloudwf.GetBids4Uid4Root API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getbids4uid4root.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBids4Uid4RootWithChan(request *GetBids4Uid4RootRequest) (<-chan *GetBids4Uid4RootResponse, <-chan error) {
	responseChan := make(chan *GetBids4Uid4RootResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBids4Uid4Root(request)
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

// GetBids4Uid4RootWithCallback invokes the cloudwf.GetBids4Uid4Root API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getbids4uid4root.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetBids4Uid4RootWithCallback(request *GetBids4Uid4RootRequest, callback func(response *GetBids4Uid4RootResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBids4Uid4RootResponse
		var err error
		defer close(result)
		response, err = client.GetBids4Uid4Root(request)
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

// GetBids4Uid4RootRequest is the request struct for api GetBids4Uid4Root
type GetBids4Uid4RootRequest struct {
	*requests.RpcRequest
	Uid requests.Integer `position:"Query" name:"Uid"`
}

// GetBids4Uid4RootResponse is the response struct for api GetBids4Uid4Root
type GetBids4Uid4RootResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetBids4Uid4RootRequest creates a request to invoke GetBids4Uid4Root API
func CreateGetBids4Uid4RootRequest() (request *GetBids4Uid4RootRequest) {
	request = &GetBids4Uid4RootRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetBids4Uid4Root", "cloudwf", "openAPI")
	return
}

// CreateGetBids4Uid4RootResponse creates a response to parse from GetBids4Uid4Root response
func CreateGetBids4Uid4RootResponse() (response *GetBids4Uid4RootResponse) {
	response = &GetBids4Uid4RootResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
