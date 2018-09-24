package cloudtables

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/pkg/errors"
)

// GetAWSAddresses fetches EC2 Elastic IP addresses.
func GetAWSAddresses(svc ec2iface.EC2API, account, region string) (*[]EC2EIPObj, int, error) {
	var EC2EIPObjs []EC2EIPObj
	input := ec2.DescribeAddressesInput{}
	obj := EC2EIPObj{}
	obj.Account = account
	obj.Region = region
	var count int
	addrs, err := svc.DescribeAddresses(&input)
	if err != nil {
		return &EC2EIPObjs, count, errors.Wrap(err, "DescribeAddresses request failed.")
	}
	count += len(addrs.Addresses)
	for _, ip := range addrs.Addresses {
		obj.PublicIP = *ip.PublicIp
		obj.Domain = *ip.Domain
		obj.InstanceID = *ip.InstanceId
		EC2EIPObjs = append(EC2EIPObjs, obj)
	}
	return &EC2EIPObjs, count, nil
}

// GetAWSInstances fetches EC2 Instances.
func GetAWSInstances(svc ec2iface.EC2API, account, region string) (*[]EC2InstObj, int, error) {
	var EC2InstObjs []EC2InstObj
	input := ec2.DescribeInstancesInput{}
	obj := EC2InstObj{}
	obj.Account = account
	obj.Region = region
	var count int
	instances, err := svc.DescribeInstances(&input)
	if err != nil {
		return &EC2InstObjs, count, errors.Wrap(err, "DescribeInstances request failed.")
	}
	for i, resv := range instances.Reservations {
		count += len(resv.Instances)
		for _, inst := range instances.Reservations[i].Instances {
			obj.AvailabilityZone = *inst.Placement.AvailabilityZone
			obj.ID = *inst.InstanceId
			obj.Key = *inst.KeyName
			obj.Name = getNameTag(inst.Tags)
			obj.PrivateIP = *inst.PrivateIpAddress
			obj.PublicIP = *inst.PublicIpAddress
			obj.Type = *inst.InstanceType
			EC2InstObjs = append(EC2InstObjs, obj)
		}
	}
	return &EC2InstObjs, count, nil
}

// Returns the value of an EC2 tag with a Key of "Name"
func getNameTag(tags []*ec2.Tag) string {
	for _, t := range tags {
		if *t.Key == "Name" {
			return *t.Value
		}
	}
	return ""
}
