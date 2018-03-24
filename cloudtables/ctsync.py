import os
import django
os.environ["DJANGO_SETTINGS_MODULE"] = 'cloudtables.settings'
django.setup()
from cloudtablesui.models import ec2, eip, s3

def syncec2(ec2conn, section, region):
	'''Sync EC2 instances to Django'''
	reservations = ec2conn.get_all_instances()
	instances = [i for r in reservations for i in r.instances]
	for i in instances:
		if i.state != 'terminated':
			printf(i.__dict__)
			ec2obj, created = ec2.objects.get_or_create(ec2id = i.__dict__['id'])
			ec2obj.name = i.__dict__['tags'].get('Name', None)
			ec2obj.ec2type = i.__dict__['instance_type']
			ec2obj.keyname = i.__dict__['key_name']
			ec2obj.pubip = i.__dict__['ip_address']
			ec2obj.privip = i.__dict__['private_ip_address']
			ec2obj.awsacct = section
			ec2obj.availzone = i.__dict__['_placement']
			ec2obj.save()

def synceip(ec2conn, section, region):
	'''Sync Elastic IPs to Django'''
	eips = ec2conn.get_all_addresses()
	for i in eips:
		eipobj, created = eip.objects.get_or_create(publicip = i.__dict__['public_ip'])
		eipobj.domain = i.__dict__['domain']
		eipobj.instanceid = i.__dict__.get('instance_id', None)
		eipobj.awsacct = section
		eipobj.region = region
		eipobj.save()

def syncs3(s3conn, section):
	'''Sync S3 buckets to Django'''
	buckets = s3conn.get_all_buckets()
	for b in buckets:
		bucketobj, created = s3.objects.get_or_create(bucket = b.__dict__['name'])
		bucketobj.awsacct = section
		bucketobj.save()