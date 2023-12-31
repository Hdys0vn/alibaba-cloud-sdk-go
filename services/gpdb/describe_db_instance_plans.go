package gpdb

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

// DescribeDBInstancePlans invokes the gpdb.DescribeDBInstancePlans API synchronously
func (client *Client) DescribeDBInstancePlans(request *DescribeDBInstancePlansRequest) (response *DescribeDBInstancePlansResponse, err error) {
	response = CreateDescribeDBInstancePlansResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstancePlansWithChan invokes the gpdb.DescribeDBInstancePlans API asynchronously
func (client *Client) DescribeDBInstancePlansWithChan(request *DescribeDBInstancePlansRequest) (<-chan *DescribeDBInstancePlansResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstancePlansResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstancePlans(request)
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

// DescribeDBInstancePlansWithCallback invokes the gpdb.DescribeDBInstancePlans API asynchronously
func (client *Client) DescribeDBInstancePlansWithCallback(request *DescribeDBInstancePlansRequest, callback func(response *DescribeDBInstancePlansResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstancePlansResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstancePlans(request)
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

// DescribeDBInstancePlansRequest is the request struct for api DescribeDBInstancePlans
type DescribeDBInstancePlansRequest struct {
	*requests.RpcRequest
	PlanType         string           `position:"Query" name:"PlanType"`
	PlanCreateDate   string           `position:"Query" name:"PlanCreateDate"`
	DBInstanceId     string           `position:"Query" name:"DBInstanceId"`
	PlanDesc         string           `position:"Query" name:"PlanDesc"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	PlanScheduleType string           `position:"Query" name:"PlanScheduleType"`
	PlanId           string           `position:"Query" name:"PlanId"`
}

// DescribeDBInstancePlansResponse is the response struct for api DescribeDBInstancePlans
type DescribeDBInstancePlansResponse struct {
	*responses.BaseResponse
	Status           string                         `json:"Status" xml:"Status"`
	ErrorMessage     string                         `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId        string                         `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int                            `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int                            `json:"PageRecordCount" xml:"PageRecordCount"`
	PageNumber       int                            `json:"PageNumber" xml:"PageNumber"`
	Items            ItemsInDescribeDBInstancePlans `json:"Items" xml:"Items"`
}

// CreateDescribeDBInstancePlansRequest creates a request to invoke DescribeDBInstancePlans API
func CreateDescribeDBInstancePlansRequest() (request *DescribeDBInstancePlansRequest) {
	request = &DescribeDBInstancePlansRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDBInstancePlans", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstancePlansResponse creates a response to parse from DescribeDBInstancePlans response
func CreateDescribeDBInstancePlansResponse() (response *DescribeDBInstancePlansResponse) {
	response = &DescribeDBInstancePlansResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
