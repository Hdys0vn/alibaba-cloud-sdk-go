package alidns

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

// DescribeDnsGtmInstance invokes the alidns.DescribeDnsGtmInstance API synchronously
func (client *Client) DescribeDnsGtmInstance(request *DescribeDnsGtmInstanceRequest) (response *DescribeDnsGtmInstanceResponse, err error) {
	response = CreateDescribeDnsGtmInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDnsGtmInstanceWithChan invokes the alidns.DescribeDnsGtmInstance API asynchronously
func (client *Client) DescribeDnsGtmInstanceWithChan(request *DescribeDnsGtmInstanceRequest) (<-chan *DescribeDnsGtmInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDnsGtmInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDnsGtmInstance(request)
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

// DescribeDnsGtmInstanceWithCallback invokes the alidns.DescribeDnsGtmInstance API asynchronously
func (client *Client) DescribeDnsGtmInstanceWithCallback(request *DescribeDnsGtmInstanceRequest, callback func(response *DescribeDnsGtmInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDnsGtmInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDnsGtmInstance(request)
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

// DescribeDnsGtmInstanceRequest is the request struct for api DescribeDnsGtmInstance
type DescribeDnsGtmInstanceRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// DescribeDnsGtmInstanceResponse is the response struct for api DescribeDnsGtmInstance
type DescribeDnsGtmInstanceResponse struct {
	*responses.BaseResponse
	ExpireTimestamp int64                          `json:"ExpireTimestamp" xml:"ExpireTimestamp"`
	RequestId       string                         `json:"RequestId" xml:"RequestId"`
	ResourceGroupId string                         `json:"ResourceGroupId" xml:"ResourceGroupId"`
	InstanceId      string                         `json:"InstanceId" xml:"InstanceId"`
	TaskQuota       int                            `json:"TaskQuota" xml:"TaskQuota"`
	CreateTime      string                         `json:"CreateTime" xml:"CreateTime"`
	SmsQuota        int                            `json:"SmsQuota" xml:"SmsQuota"`
	VersionCode     string                         `json:"VersionCode" xml:"VersionCode"`
	PaymentType     string                         `json:"PaymentType" xml:"PaymentType"`
	ExpireTime      string                         `json:"ExpireTime" xml:"ExpireTime"`
	CreateTimestamp int64                          `json:"CreateTimestamp" xml:"CreateTimestamp"`
	Config          ConfigInDescribeDnsGtmInstance `json:"Config" xml:"Config"`
	UsedQuota       UsedQuota                      `json:"UsedQuota" xml:"UsedQuota"`
}

// CreateDescribeDnsGtmInstanceRequest creates a request to invoke DescribeDnsGtmInstance API
func CreateDescribeDnsGtmInstanceRequest() (request *DescribeDnsGtmInstanceRequest) {
	request = &DescribeDnsGtmInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeDnsGtmInstance", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDnsGtmInstanceResponse creates a response to parse from DescribeDnsGtmInstance response
func CreateDescribeDnsGtmInstanceResponse() (response *DescribeDnsGtmInstanceResponse) {
	response = &DescribeDnsGtmInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
