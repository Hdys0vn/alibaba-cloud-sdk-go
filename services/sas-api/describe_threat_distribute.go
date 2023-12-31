package sas_api

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

// DescribeThreatDistribute invokes the sas_api.DescribeThreatDistribute API synchronously
// api document: https://help.aliyun.com/api/sas-api/describethreatdistribute.html
func (client *Client) DescribeThreatDistribute(request *DescribeThreatDistributeRequest) (response *DescribeThreatDistributeResponse, err error) {
	response = CreateDescribeThreatDistributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeThreatDistributeWithChan invokes the sas_api.DescribeThreatDistribute API asynchronously
// api document: https://help.aliyun.com/api/sas-api/describethreatdistribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeThreatDistributeWithChan(request *DescribeThreatDistributeRequest) (<-chan *DescribeThreatDistributeResponse, <-chan error) {
	responseChan := make(chan *DescribeThreatDistributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeThreatDistribute(request)
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

// DescribeThreatDistributeWithCallback invokes the sas_api.DescribeThreatDistribute API asynchronously
// api document: https://help.aliyun.com/api/sas-api/describethreatdistribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeThreatDistributeWithCallback(request *DescribeThreatDistributeRequest, callback func(response *DescribeThreatDistributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeThreatDistributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeThreatDistribute(request)
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

// DescribeThreatDistributeRequest is the request struct for api DescribeThreatDistribute
type DescribeThreatDistributeRequest struct {
	*requests.RpcRequest
	EndDate   string           `position:"Query" name:"EndDate"`
	SourceIp  string           `position:"Query" name:"SourceIp"`
	HitDay    requests.Integer `position:"Query" name:"HitDay"`
	StartDate string           `position:"Query" name:"StartDate"`
	ApiType   requests.Integer `position:"Query" name:"ApiType"`
}

// DescribeThreatDistributeResponse is the response struct for api DescribeThreatDistribute
type DescribeThreatDistributeResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	Categories []string `json:"Categories" xml:"Categories"`
	Items      []Item   `json:"Items" xml:"Items"`
}

// CreateDescribeThreatDistributeRequest creates a request to invoke DescribeThreatDistribute API
func CreateDescribeThreatDistributeRequest() (request *DescribeThreatDistributeRequest) {
	request = &DescribeThreatDistributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas-api", "2017-07-05", "DescribeThreatDistribute", "sas-api", "openAPI")
	return
}

// CreateDescribeThreatDistributeResponse creates a response to parse from DescribeThreatDistribute response
func CreateDescribeThreatDistributeResponse() (response *DescribeThreatDistributeResponse) {
	response = &DescribeThreatDistributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
