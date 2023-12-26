package ebs

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

// CreateDiskReplicaPair invokes the ebs.CreateDiskReplicaPair API synchronously
func (client *Client) CreateDiskReplicaPair(request *CreateDiskReplicaPairRequest) (response *CreateDiskReplicaPairResponse, err error) {
	response = CreateCreateDiskReplicaPairResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDiskReplicaPairWithChan invokes the ebs.CreateDiskReplicaPair API asynchronously
func (client *Client) CreateDiskReplicaPairWithChan(request *CreateDiskReplicaPairRequest) (<-chan *CreateDiskReplicaPairResponse, <-chan error) {
	responseChan := make(chan *CreateDiskReplicaPairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDiskReplicaPair(request)
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

// CreateDiskReplicaPairWithCallback invokes the ebs.CreateDiskReplicaPair API asynchronously
func (client *Client) CreateDiskReplicaPairWithCallback(request *CreateDiskReplicaPairRequest, callback func(response *CreateDiskReplicaPairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDiskReplicaPairResponse
		var err error
		defer close(result)
		response, err = client.CreateDiskReplicaPair(request)
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

// CreateDiskReplicaPairRequest is the request struct for api CreateDiskReplicaPair
type CreateDiskReplicaPairRequest struct {
	*requests.RpcRequest
	PairName            string                      `position:"Query" name:"PairName"`
	ResourceGroupId     string                      `position:"Query" name:"ResourceGroupId"`
	Tag                 *[]CreateDiskReplicaPairTag `position:"Query" name:"Tag"  type:"Repeated"`
	DestinationCloudId  string                      `position:"Query" name:"DestinationCloudId"`
	Period              requests.Integer            `position:"Query" name:"Period"`
	PeriodUnit          string                      `position:"Query" name:"PeriodUnit"`
	DestinationDiskId   string                      `position:"Query" name:"DestinationDiskId"`
	ClientToken         string                      `position:"Query" name:"ClientToken"`
	DestinationRegionId string                      `position:"Query" name:"DestinationRegionId"`
	Description         string                      `position:"Query" name:"Description"`
	SourceZoneId        string                      `position:"Query" name:"SourceZoneId"`
	DestinationZoneId   string                      `position:"Query" name:"DestinationZoneId"`
	DiskId              string                      `position:"Query" name:"DiskId"`
	Bandwidth           requests.Integer            `position:"Query" name:"Bandwidth"`
	DestinationUid      string                      `position:"Query" name:"DestinationUid"`
	RPO                 requests.Integer            `position:"Query" name:"RPO"`
	ChargeType          string                      `position:"Query" name:"ChargeType"`
}

// CreateDiskReplicaPairTag is a repeated param struct in CreateDiskReplicaPairRequest
type CreateDiskReplicaPairTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateDiskReplicaPairResponse is the response struct for api CreateDiskReplicaPair
type CreateDiskReplicaPairResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ReplicaPairId string `json:"ReplicaPairId" xml:"ReplicaPairId"`
	OrderId       string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateDiskReplicaPairRequest creates a request to invoke CreateDiskReplicaPair API
func CreateCreateDiskReplicaPairRequest() (request *CreateDiskReplicaPairRequest) {
	request = &CreateDiskReplicaPairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ebs", "2021-07-30", "CreateDiskReplicaPair", "ebs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDiskReplicaPairResponse creates a response to parse from CreateDiskReplicaPair response
func CreateCreateDiskReplicaPairResponse() (response *CreateDiskReplicaPairResponse) {
	response = &CreateDiskReplicaPairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}