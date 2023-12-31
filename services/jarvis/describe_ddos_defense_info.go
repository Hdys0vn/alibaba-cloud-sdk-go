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

// DescribeDdosDefenseInfo invokes the jarvis.DescribeDdosDefenseInfo API synchronously
// api document: https://help.aliyun.com/api/jarvis/describeddosdefenseinfo.html
func (client *Client) DescribeDdosDefenseInfo(request *DescribeDdosDefenseInfoRequest) (response *DescribeDdosDefenseInfoResponse, err error) {
	response = CreateDescribeDdosDefenseInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDdosDefenseInfoWithChan invokes the jarvis.DescribeDdosDefenseInfo API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describeddosdefenseinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDdosDefenseInfoWithChan(request *DescribeDdosDefenseInfoRequest) (<-chan *DescribeDdosDefenseInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDdosDefenseInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDdosDefenseInfo(request)
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

// DescribeDdosDefenseInfoWithCallback invokes the jarvis.DescribeDdosDefenseInfo API asynchronously
// api document: https://help.aliyun.com/api/jarvis/describeddosdefenseinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDdosDefenseInfoWithCallback(request *DescribeDdosDefenseInfoRequest, callback func(response *DescribeDdosDefenseInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDdosDefenseInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDdosDefenseInfo(request)
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

// DescribeDdosDefenseInfoRequest is the request struct for api DescribeDdosDefenseInfo
type DescribeDdosDefenseInfoRequest struct {
	*requests.RpcRequest
	SourceIp   string           `position:"Query" name:"SourceIp"`
	Lang       string           `position:"Query" name:"Lang"`
	SrcUid     requests.Integer `position:"Query" name:"srcUid"`
	SourceCode string           `position:"Query" name:"sourceCode"`
}

// DescribeDdosDefenseInfoResponse is the response struct for api DescribeDdosDefenseInfo
type DescribeDdosDefenseInfoResponse struct {
	*responses.BaseResponse
	RequestId            string                     `json:"RequestId" xml:"RequestId"`
	Module               string                     `json:"Module" xml:"Module"`
	BlackTimes           int                        `json:"BlackTimes" xml:"BlackTimes"`
	Duration             int                        `json:"Duration" xml:"Duration"`
	BgpPkgState          string                     `json:"BgpPkgState" xml:"BgpPkgState"`
	DdosDefenseThreshold []DdosDefenseThresholdItem `json:"DdosDefenseThreshold" xml:"DdosDefenseThreshold"`
}

// CreateDescribeDdosDefenseInfoRequest creates a request to invoke DescribeDdosDefenseInfo API
func CreateDescribeDdosDefenseInfoRequest() (request *DescribeDdosDefenseInfoRequest) {
	request = &DescribeDdosDefenseInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jarvis", "2018-02-06", "DescribeDdosDefenseInfo", "jarvis", "openAPI")
	return
}

// CreateDescribeDdosDefenseInfoResponse creates a response to parse from DescribeDdosDefenseInfo response
func CreateDescribeDdosDefenseInfoResponse() (response *DescribeDdosDefenseInfoResponse) {
	response = &DescribeDdosDefenseInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
