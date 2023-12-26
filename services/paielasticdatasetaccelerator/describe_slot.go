package paielasticdatasetaccelerator

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

// DescribeSlot invokes the paielasticdatasetaccelerator.DescribeSlot API synchronously
func (client *Client) DescribeSlot(request *DescribeSlotRequest) (response *DescribeSlotResponse, err error) {
	response = CreateDescribeSlotResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSlotWithChan invokes the paielasticdatasetaccelerator.DescribeSlot API asynchronously
func (client *Client) DescribeSlotWithChan(request *DescribeSlotRequest) (<-chan *DescribeSlotResponse, <-chan error) {
	responseChan := make(chan *DescribeSlotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSlot(request)
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

// DescribeSlotWithCallback invokes the paielasticdatasetaccelerator.DescribeSlot API asynchronously
func (client *Client) DescribeSlotWithCallback(request *DescribeSlotRequest, callback func(response *DescribeSlotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSlotResponse
		var err error
		defer close(result)
		response, err = client.DescribeSlot(request)
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

// DescribeSlotRequest is the request struct for api DescribeSlot
type DescribeSlotRequest struct {
	*requests.RoaRequest
	SlotId string `position:"Path" name:"SlotId"`
}

// DescribeSlotResponse is the response struct for api DescribeSlot
type DescribeSlotResponse struct {
	*responses.BaseResponse
	RequestId       string     `json:"RequestId" xml:"RequestId"`
	UserId          string     `json:"UserId" xml:"UserId"`
	OwnerId         string     `json:"OwnerId" xml:"OwnerId"`
	GmtCreateTime   string     `json:"GmtCreateTime" xml:"GmtCreateTime"`
	GmtModifiedTime string     `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
	Uuid            string     `json:"Uuid" xml:"Uuid"`
	InstanceId      string     `json:"InstanceId" xml:"InstanceId"`
	Name            string     `json:"Name" xml:"Name"`
	Description     string     `json:"Description" xml:"Description"`
	StorageType     string     `json:"StorageType" xml:"StorageType"`
	StorageUri      string     `json:"StorageUri" xml:"StorageUri"`
	Capacity        string     `json:"Capacity" xml:"Capacity"`
	IoType          string     `json:"IoType" xml:"IoType"`
	LifeCycle       LifeCycle  `json:"LifeCycle" xml:"LifeCycle"`
	Status          Status     `json:"Status" xml:"Status"`
	Tags            []TagsItem `json:"Tags" xml:"Tags"`
}

// CreateDescribeSlotRequest creates a request to invoke DescribeSlot API
func CreateDescribeSlotRequest() (request *DescribeSlotRequest) {
	request = &DescribeSlotRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "DescribeSlot", "/api/v1/slots/[SlotId]", "datasetacc", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeSlotResponse creates a response to parse from DescribeSlot response
func CreateDescribeSlotResponse() (response *DescribeSlotResponse) {
	response = &DescribeSlotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
