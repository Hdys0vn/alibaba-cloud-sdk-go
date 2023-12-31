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

// ListUpgradeImg invokes the cloudwf.ListUpgradeImg API synchronously
// api document: https://help.aliyun.com/api/cloudwf/listupgradeimg.html
func (client *Client) ListUpgradeImg(request *ListUpgradeImgRequest) (response *ListUpgradeImgResponse, err error) {
	response = CreateListUpgradeImgResponse()
	err = client.DoAction(request, response)
	return
}

// ListUpgradeImgWithChan invokes the cloudwf.ListUpgradeImg API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listupgradeimg.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUpgradeImgWithChan(request *ListUpgradeImgRequest) (<-chan *ListUpgradeImgResponse, <-chan error) {
	responseChan := make(chan *ListUpgradeImgResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUpgradeImg(request)
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

// ListUpgradeImgWithCallback invokes the cloudwf.ListUpgradeImg API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/listupgradeimg.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListUpgradeImgWithCallback(request *ListUpgradeImgRequest, callback func(response *ListUpgradeImgResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUpgradeImgResponse
		var err error
		defer close(result)
		response, err = client.ListUpgradeImg(request)
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

// ListUpgradeImgRequest is the request struct for api ListUpgradeImg
type ListUpgradeImgRequest struct {
	*requests.RpcRequest
	Length    requests.Integer `position:"Query" name:"Length"`
	PageIndex requests.Integer `position:"Query" name:"PageIndex"`
}

// ListUpgradeImgResponse is the response struct for api ListUpgradeImg
type ListUpgradeImgResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateListUpgradeImgRequest creates a request to invoke ListUpgradeImg API
func CreateListUpgradeImgRequest() (request *ListUpgradeImgRequest) {
	request = &ListUpgradeImgRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "ListUpgradeImg", "cloudwf", "openAPI")
	return
}

// CreateListUpgradeImgResponse creates a response to parse from ListUpgradeImg response
func CreateListUpgradeImgResponse() (response *ListUpgradeImgResponse) {
	response = &ListUpgradeImgResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
