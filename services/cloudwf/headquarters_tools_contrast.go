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

// HeadquartersToolsContrast invokes the cloudwf.HeadquartersToolsContrast API synchronously
// api document: https://help.aliyun.com/api/cloudwf/headquarterstoolscontrast.html
func (client *Client) HeadquartersToolsContrast(request *HeadquartersToolsContrastRequest) (response *HeadquartersToolsContrastResponse, err error) {
	response = CreateHeadquartersToolsContrastResponse()
	err = client.DoAction(request, response)
	return
}

// HeadquartersToolsContrastWithChan invokes the cloudwf.HeadquartersToolsContrast API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/headquarterstoolscontrast.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) HeadquartersToolsContrastWithChan(request *HeadquartersToolsContrastRequest) (<-chan *HeadquartersToolsContrastResponse, <-chan error) {
	responseChan := make(chan *HeadquartersToolsContrastResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.HeadquartersToolsContrast(request)
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

// HeadquartersToolsContrastWithCallback invokes the cloudwf.HeadquartersToolsContrast API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/headquarterstoolscontrast.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) HeadquartersToolsContrastWithCallback(request *HeadquartersToolsContrastRequest, callback func(response *HeadquartersToolsContrastResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *HeadquartersToolsContrastResponse
		var err error
		defer close(result)
		response, err = client.HeadquartersToolsContrast(request)
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

// HeadquartersToolsContrastRequest is the request struct for api HeadquartersToolsContrast
type HeadquartersToolsContrastRequest struct {
	*requests.RpcRequest
	Bid requests.Integer `position:"Query" name:"Bid"`
}

// HeadquartersToolsContrastResponse is the response struct for api HeadquartersToolsContrast
type HeadquartersToolsContrastResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateHeadquartersToolsContrastRequest creates a request to invoke HeadquartersToolsContrast API
func CreateHeadquartersToolsContrastRequest() (request *HeadquartersToolsContrastRequest) {
	request = &HeadquartersToolsContrastRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersToolsContrast", "cloudwf", "openAPI")
	return
}

// CreateHeadquartersToolsContrastResponse creates a response to parse from HeadquartersToolsContrast response
func CreateHeadquartersToolsContrastResponse() (response *HeadquartersToolsContrastResponse) {
	response = &HeadquartersToolsContrastResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
