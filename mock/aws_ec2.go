package mock

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// EC2Client mocks the AWS API EC2 endpoint.
type EC2Client struct{ ec2iface.EC2API }

// DescribeInstances returns a mock EC2 instance.
func (e *EC2Client) DescribeInstances(i *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	nameKey, nameValue, id, itype, key, privIP, pubIP, zone := "Name", "or2ccenk8p123", "i-cjdu38ok3jd839ck3", "r5d.24xlarge", "ops-production",
		"10.1.2.100", "192.0.2.100", "us-west-2c"
	place := ec2.Placement{AvailabilityZone: &zone}
	nameTag := ec2.Tag{Key: &nameKey, Value: &nameValue}
	i1 := ec2.Instance{
		InstanceId:       &id, // 17 char
		InstanceType:     &itype,
		KeyName:          &key,
		Tags:             []*ec2.Tag{&nameTag},
		PrivateIpAddress: &privIP,
		PublicIpAddress:  &pubIP,
		Placement:        &place,
	}
	res1 := ec2.Reservation{
		Instances: []*ec2.Instance{&i1},
	}
	r := ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{&res1},
	}
	return &r, nil
}

// DescribeAddresses returns a mock Elastic IP address
func (e *EC2Client) DescribeAddresses(input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	domain, inst, pubIP := "devops4.life", "i-cjdu38ok3jd839ck3", "192.0.2.100"
	addr := ec2.Address{Domain: &domain, InstanceId: &inst, PublicIp: &pubIP}
	r := ec2.DescribeAddressesOutput{Addresses: []*ec2.Address{&addr}}
	return &r, nil
}
