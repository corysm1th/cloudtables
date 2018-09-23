package cloudtables

import (
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/pkg/errors"
)

// SyncDescribeInstances fetches EC2 Instances, and stores them in the database.
func SyncDescribeInstances(svc ec2iface.EC2API, account, region string) error {
	input := ec2.DescribeInstancesInput{}
	obj := EC2Obj{}
	obj.Account = account
	obj.Region = region
	var instCount int
	instances, err := svc.DescribeInstances(&input)
	if err != nil {
		return errors.Wrap(err, "DescribeInstances request failed.")
	}
	log.Println("EC2")
	for i, resv := range instances.Reservations {
		instCount += len(resv.Instances)
		for _, inst := range instances.Reservations[i].Instances {
			obj.AvailabilityZone = *inst.Placement.AvailabilityZone
			obj.ID = *inst.InstanceId
			obj.Key = *inst.KeyName
			obj.Name = getNameTag(inst.Tags)
			obj.PrivateIP = *inst.PrivateIpAddress
			obj.PublicIP = *inst.PublicIpAddress
			obj.Type = *inst.InstanceType
			// TODO: Store instance in ec2_instance_obj
			log.Printf("Account: %s  Region: %s  ID: %s  ", obj.Account, obj.Region, obj.Name)
		}
	}
	// TODO: Store count in metric
	return nil
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
