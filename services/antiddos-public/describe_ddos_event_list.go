package antiddos_public

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

// DescribeDdosEventList invokes the antiddos_public.DescribeDdosEventList API synchronously
func (client *Client) DescribeDdosEventList(request *DescribeDdosEventListRequest) (response *DescribeDdosEventListResponse, err error) {
	response = CreateDescribeDdosEventListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDdosEventListWithChan invokes the antiddos_public.DescribeDdosEventList API asynchronously
func (client *Client) DescribeDdosEventListWithChan(request *DescribeDdosEventListRequest) (<-chan *DescribeDdosEventListResponse, <-chan error) {
	responseChan := make(chan *DescribeDdosEventListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDdosEventList(request)
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

// DescribeDdosEventListWithCallback invokes the antiddos_public.DescribeDdosEventList API asynchronously
func (client *Client) DescribeDdosEventListWithCallback(request *DescribeDdosEventListRequest, callback func(response *DescribeDdosEventListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDdosEventListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDdosEventList(request)
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

// DescribeDdosEventListRequest is the request struct for api DescribeDdosEventList
type DescribeDdosEventListRequest struct {
	*requests.RpcRequest
	InternetIp   string           `position:"Query" name:"InternetIp"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	DdosRegionId string           `position:"Query" name:"DdosRegionId"`
	InstanceType string           `position:"Query" name:"InstanceType"`
	Lang         string           `position:"Query" name:"Lang"`
	CurrentPage  requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
}

// DescribeDdosEventListResponse is the response struct for api DescribeDdosEventList
type DescribeDdosEventListResponse struct {
	*responses.BaseResponse
	Total         int           `json:"Total" xml:"Total"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	DdosEventList DdosEventList `json:"DdosEventList" xml:"DdosEventList"`
}

// CreateDescribeDdosEventListRequest creates a request to invoke DescribeDdosEventList API
func CreateDescribeDdosEventListRequest() (request *DescribeDdosEventListRequest) {
	request = &DescribeDdosEventListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("antiddos-public", "2017-05-18", "DescribeDdosEventList", "ddosbasic", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDdosEventListResponse creates a response to parse from DescribeDdosEventList response
func CreateDescribeDdosEventListResponse() (response *DescribeDdosEventListResponse) {
	response = &DescribeDdosEventListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
