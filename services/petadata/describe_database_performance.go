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

// DescribeDatabasePerformance invokes the petadata.DescribeDatabasePerformance API synchronously
// api document: https://help.aliyun.com/api/petadata/describedatabaseperformance.html
func (client *Client) DescribeDatabasePerformance(request *DescribeDatabasePerformanceRequest) (response *DescribeDatabasePerformanceResponse, err error) {
	response = CreateDescribeDatabasePerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDatabasePerformanceWithChan invokes the petadata.DescribeDatabasePerformance API asynchronously
// api document: https://help.aliyun.com/api/petadata/describedatabaseperformance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDatabasePerformanceWithChan(request *DescribeDatabasePerformanceRequest) (<-chan *DescribeDatabasePerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDatabasePerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDatabasePerformance(request)
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

// DescribeDatabasePerformanceWithCallback invokes the petadata.DescribeDatabasePerformance API asynchronously
// api document: https://help.aliyun.com/api/petadata/describedatabaseperformance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDatabasePerformanceWithCallback(request *DescribeDatabasePerformanceRequest, callback func(response *DescribeDatabasePerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDatabasePerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDatabasePerformance(request)
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

// DescribeDatabasePerformanceRequest is the request struct for api DescribeDatabasePerformance
type DescribeDatabasePerformanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	KeyList              string           `position:"Query" name:"KeyList"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	DBName               string           `position:"Query" name:"DBName"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	MonitorGroup         string           `position:"Query" name:"MonitorGroup"`
	Interval             string           `position:"Query" name:"Interval"`
}

// DescribeDatabasePerformanceResponse is the response struct for api DescribeDatabasePerformance
type DescribeDatabasePerformanceResponse struct {
	*responses.BaseResponse
	RequestId    string                                    `json:"RequestId" xml:"RequestId"`
	MonitorDatas MonitorDatasInDescribeDatabasePerformance `json:"MonitorDatas" xml:"MonitorDatas"`
}

// CreateDescribeDatabasePerformanceRequest creates a request to invoke DescribeDatabasePerformance API
func CreateDescribeDatabasePerformanceRequest() (request *DescribeDatabasePerformanceRequest) {
	request = &DescribeDatabasePerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PetaData", "2016-01-01", "DescribeDatabasePerformance", "petadata", "openAPI")
	return
}

// CreateDescribeDatabasePerformanceResponse creates a response to parse from DescribeDatabasePerformance response
func CreateDescribeDatabasePerformanceResponse() (response *DescribeDatabasePerformanceResponse) {
	response = &DescribeDatabasePerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
