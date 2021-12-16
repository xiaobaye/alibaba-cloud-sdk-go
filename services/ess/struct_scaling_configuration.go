package ess

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

// ScalingConfiguration is a nested struct in ess response
type ScalingConfiguration struct {
	ScalingConfigurationId          string               `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
	ScalingConfigurationName        string               `json:"ScalingConfigurationName" xml:"ScalingConfigurationName"`
	ScalingGroupId                  string               `json:"ScalingGroupId" xml:"ScalingGroupId"`
	InstanceName                    string               `json:"InstanceName" xml:"InstanceName"`
	ImageId                         string               `json:"ImageId" xml:"ImageId"`
	ImageName                       string               `json:"ImageName" xml:"ImageName"`
	HostName                        string               `json:"HostName" xml:"HostName"`
	InstanceType                    string               `json:"InstanceType" xml:"InstanceType"`
	Cpu                             int                  `json:"Cpu" xml:"Cpu"`
	Memory                          int                  `json:"Memory" xml:"Memory"`
	InstanceGeneration              string               `json:"InstanceGeneration" xml:"InstanceGeneration"`
	SecurityGroupId                 string               `json:"SecurityGroupId" xml:"SecurityGroupId"`
	IoOptimized                     string               `json:"IoOptimized" xml:"IoOptimized"`
	InternetChargeType              string               `json:"InternetChargeType" xml:"InternetChargeType"`
	InternetMaxBandwidthIn          int                  `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut         int                  `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	SystemDiskCategory              string               `json:"SystemDiskCategory" xml:"SystemDiskCategory"`
	SystemDiskSize                  int                  `json:"SystemDiskSize" xml:"SystemDiskSize"`
	SystemDiskName                  string               `json:"SystemDiskName" xml:"SystemDiskName"`
	SystemDiskDescription           string               `json:"SystemDiskDescription" xml:"SystemDiskDescription"`
	SystemDiskAutoSnapshotPolicyId  string               `json:"SystemDiskAutoSnapshotPolicyId" xml:"SystemDiskAutoSnapshotPolicyId"`
	SystemDiskPerformanceLevel      string               `json:"SystemDiskPerformanceLevel" xml:"SystemDiskPerformanceLevel"`
	LifecycleState                  string               `json:"LifecycleState" xml:"LifecycleState"`
	CreationTime                    string               `json:"CreationTime" xml:"CreationTime"`
	LoadBalancerWeight              int                  `json:"LoadBalancerWeight" xml:"LoadBalancerWeight"`
	UserData                        string               `json:"UserData" xml:"UserData"`
	KeyPairName                     string               `json:"KeyPairName" xml:"KeyPairName"`
	RamRoleName                     string               `json:"RamRoleName" xml:"RamRoleName"`
	DeploymentSetId                 string               `json:"DeploymentSetId" xml:"DeploymentSetId"`
	SecurityEnhancementStrategy     string               `json:"SecurityEnhancementStrategy" xml:"SecurityEnhancementStrategy"`
	SpotStrategy                    string               `json:"SpotStrategy" xml:"SpotStrategy"`
	PasswordInherit                 bool                 `json:"PasswordInherit" xml:"PasswordInherit"`
	ResourceGroupId                 string               `json:"ResourceGroupId" xml:"ResourceGroupId"`
	HpcClusterId                    string               `json:"HpcClusterId" xml:"HpcClusterId"`
	InstanceDescription             string               `json:"InstanceDescription" xml:"InstanceDescription"`
	CreditSpecification             string               `json:"CreditSpecification" xml:"CreditSpecification"`
	ImageFamily                     string               `json:"ImageFamily" xml:"ImageFamily"`
	ZoneId                          string               `json:"ZoneId" xml:"ZoneId"`
	DedicatedHostId                 string               `json:"DedicatedHostId" xml:"DedicatedHostId"`
	Affinity                        string               `json:"Affinity" xml:"Affinity"`
	Tenancy                         string               `json:"Tenancy" xml:"Tenancy"`
	PrivatePoolOptionsMatchCriteria string               `json:"PrivatePoolOptions.MatchCriteria" xml:"PrivatePoolOptions.MatchCriteria"`
	PrivatePoolOptionsId            string               `json:"PrivatePoolOptions.Id" xml:"PrivatePoolOptions.Id"`
	SpotInterruptionBehavior        string               `json:"SpotInterruptionBehavior" xml:"SpotInterruptionBehavior"`
	SpotDuration                    int                  `json:"SpotDuration" xml:"SpotDuration"`
	Ipv6AddressCount                int                  `json:"Ipv6AddressCount" xml:"Ipv6AddressCount"`
	InstanceTypes                   InstanceTypes        `json:"InstanceTypes" xml:"InstanceTypes"`
	WeightedCapacities              WeightedCapacities   `json:"WeightedCapacities" xml:"WeightedCapacities"`
	SecurityGroupIds                SecurityGroupIds     `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	SystemDiskCategories            SystemDiskCategories `json:"SystemDiskCategories" xml:"SystemDiskCategories"`
	SchedulerOptions                SchedulerOptions     `json:"SchedulerOptions" xml:"SchedulerOptions"`
	DataDisks                       DataDisks            `json:"DataDisks" xml:"DataDisks"`
	Tags                            Tags                 `json:"Tags" xml:"Tags"`
	SpotPriceLimit                  SpotPriceLimit       `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
	InstancePatternInfos            InstancePatternInfos `json:"InstancePatternInfos" xml:"InstancePatternInfos"`
}
