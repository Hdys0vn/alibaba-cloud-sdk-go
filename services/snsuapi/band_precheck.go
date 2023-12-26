package snsuapi

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

// BandPrecheck invokes the snsuapi.BandPrecheck API synchronously
// api document: https://help.aliyun.com/api/snsuapi/bandprecheck.html
func (client *Client) BandPrecheck(request *BandPrecheckRequest) (response *BandPrecheckResponse, err error) {
	response = CreateBandPrecheckResponse()
	err = client.DoAction(request, response)
	return
}

// BandPrecheckWithChan invokes the snsuapi.BandPrecheck API asynchronously
// api document: https://help.aliyun.com/api/snsuapi/bandprecheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BandPrecheckWithChan(request *BandPrecheckRequest) (<-chan *BandPrecheckResponse, <-chan error) {
	responseChan := make(chan *BandPrecheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BandPrecheck(request)
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

// BandPrecheckWithCallback invokes the snsuapi.BandPrecheck API asynchronously
// api document: https://help.aliyun.com/api/snsuapi/bandprecheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BandPrecheckWithCallback(request *BandPrecheckRequest, callback func(response *BandPrecheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BandPrecheckResponse
		var err error
		defer close(result)
		response, err = client.BandPrecheck(request)
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

// BandPrecheckRequest is the request struct for api BandPrecheck
type BandPrecheckRequest struct {
	*requests.RpcRequest
	IpAddress            string           `position:"Query" name:"IpAddress"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Port                 requests.Integer `position:"Query" name:"Port"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// BandPrecheckResponse is the response struct for api BandPrecheck
type BandPrecheckResponse struct {
	*responses.BaseResponse
	RequestId     string       `json:"RequestId" xml:"RequestId"`
	ResultCode    string       `json:"ResultCode" xml:"ResultCode"`
	ResultMessage string       `json:"ResultMessage" xml:"ResultMessage"`
	ResultModule  ResultModule `json:"ResultModule" xml:"ResultModule"`
}

// CreateBandPrecheckRequest creates a request to invoke BandPrecheck API
func CreateBandPrecheckRequest() (request *BandPrecheckRequest) {
	request = &BandPrecheckRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Snsuapi", "2018-07-09", "BandPrecheck", "snsuapi", "openAPI")
	return
}

// CreateBandPrecheckResponse creates a response to parse from BandPrecheck response
func CreateBandPrecheckResponse() (response *BandPrecheckResponse) {
	response = &BandPrecheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
