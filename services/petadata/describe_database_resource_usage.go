package petadata

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

// DescribeDatabaseResourceUsage invokes the petadata.DescribeDatabaseResourceUsage API synchronously
// api document: https://help.aliyun.com/api/petadata/describedatabaseresourceusage.html
func (client *Client) DescribeDatabaseResourceUsage(request *DescribeDatabaseResourceUsageRequest) (response *DescribeDatabaseResourceUsageResponse, err error) {
	response = CreateDescribeDatabaseResourceUsageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDatabaseResourceUsageWithChan invokes the petadata.DescribeDatabaseResourceUsage API asynchronously
// api document: https://help.aliyun.com/api/petadata/describedatabaseresourceusage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDatabaseResourceUsageWithChan(request *DescribeDatabaseResourceUsageRequest) (<-chan *DescribeDatabaseResourceUsageResponse, <-chan error) {
	responseChan := make(chan *DescribeDatabaseResourceUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDatabaseResourceUsage(request)
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

// DescribeDatabaseResourceUsageWithCallback invokes the petadata.DescribeDatabaseResourceUsage API asynchronously
// api document: https://help.aliyun.com/api/petadata/describedatabaseresourceusage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDatabaseResourceUsageWithCallback(request *DescribeDatabaseResourceUsageRequest, callback func(response *DescribeDatabaseResourceUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDatabaseResourceUsageResponse
		var err error
		defer close(result)
		response, err = client.DescribeDatabaseResourceUsage(request)
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

// DescribeDatabaseResourceUsageRequest is the request struct for api DescribeDatabaseResourceUsage
type DescribeDatabaseResourceUsageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	DBName               string           `position:"Query" name:"DBName"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Interval             string           `position:"Query" name:"Interval"`
}

// DescribeDatabaseResourceUsageResponse is the response struct for api DescribeDatabaseResourceUsage
type DescribeDatabaseResourceUsageResponse struct {
	*responses.BaseResponse
	RequestId    string                                      `json:"RequestId" xml:"RequestId"`
	MonitorDatas MonitorDatasInDescribeDatabaseResourceUsage `json:"MonitorDatas" xml:"MonitorDatas"`
}

// CreateDescribeDatabaseResourceUsageRequest creates a request to invoke DescribeDatabaseResourceUsage API
func CreateDescribeDatabaseResourceUsageRequest() (request *DescribeDatabaseResourceUsageRequest) {
	request = &DescribeDatabaseResourceUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PetaData", "2016-01-01", "DescribeDatabaseResourceUsage", "petadata", "openAPI")
	return
}

// CreateDescribeDatabaseResourceUsageResponse creates a response to parse from DescribeDatabaseResourceUsage response
func CreateDescribeDatabaseResourceUsageResponse() (response *DescribeDatabaseResourceUsageResponse) {
	response = &DescribeDatabaseResourceUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
