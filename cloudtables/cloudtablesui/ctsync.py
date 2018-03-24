import json
from .models import ec2, eip, s3

def syncec2(ec2conn, section, region):
	'''Sync EC2 instances to Django'''
	print "Getting instances for Account %s, Region %s" % (section, region)
	reservations = ec2conn.get_all_instances()

	instances = [i for r in reservations for i in r.instances]
	print "Updating instances"
	for i in instances:
		if i.state != 'terminated':
			ec2obj, created = ec2.objects.get_or_create(ec2id = i.__dict__['id'])
			ec2obj.name = i.__dict__['tags'].get('Name', None)
			ec2obj.ec2type = i.__dict__['instance_type']
			ec2obj.keyname = i.__dict__['key_name']
			ec2obj.pubip = i.__dict__['ip_address']
			ec2obj.privip = i.__dict__['private_ip_address']
			ec2obj.awsacct = section
			ec2obj.availzone = i.__dict__['_placement']
			ec2obj.ec2region = region
			ec2obj.save()

	'''Clean up terminated instances'''
	print "Removing old instances"
	for obj in ec2.objects.filter(awsacct=section,ec2region=region):
		objexists = False
		for i in instances:
			if i.__dict__['id'] == obj.ec2id:
				objexists = True
		if not objexists:
			print "Deleting instance %s" % obj.ec2id
			obj.delete()

def synceip(ec2conn, section, region):
	'''Sync Elastic IPs to Django'''
	print "Getting EIPs"
	eips = ec2conn.get_all_addresses()
	print "Updating EIPs"
	for i in eips:
		eipobj, created = eip.objects.get_or_create(publicip = i.__dict__['public_ip'])
		eipobj.domain = i.__dict__['domain']
		eipobj.instanceid = i.__dict__.get('instance_id', None)
		eipobj.awsacct = section
		eipobj.region = region
		eipobj.save()

	'''Clean up released EIPs'''
	print "Cleaning up old EIPs"
	for obj in eip.objects.filter(awsacct=section,region=region):
		objexists = False
		for i in eips:
			if i.__dict__['public_ip'] == obj.publicip:
				objexists = True
		if not objexists:
			print "Deleting EIP %s" % obj.publicip
			obj.delete()

def syncs3(s3conn, section):
	'''Sync S3 buckets to Django'''
	print "Getting S3 Buckets"
	buckets = s3conn.get_all_buckets()
	print "Updating S3 Buckets"
	for b in buckets:
		bucketobj, created = s3.objects.get_or_create(bucket = b.__dict__['name'])
		bucketobj.awsacct = section
		bucketobj.save()

	'''Clean up deleted buckets'''
	print "Cleaning up deleted buckets"
	for obj in s3.objects.filter(awsacct=section):
		objexists = False
		for b in buckets:
			if b.__dict__['name'] == obj.bucket:
				objexists = True
		if not objexists:
			print "Deleting Bucket %s" % obj.bucket
			obj.delete()