# This project is retired because I'm rewriting it in Go and React.

Feel free to submit feature requests for the new version.

On the roadmap:

* Support for more resource types
* Support for more cloud providers

# CloudTables

Cloud Resource Manager

Provides an easily searchable database of AWS resources across multiple accounts.

## Installation

---
**NOTE**

If you skip any of these steps, your install will explode like a hot pocket in the microwave.

---

### Set up AWS Credentials

Start by setting up your AWS credentials file.

```sh
mkdir $HOME/.aws
cat <<EOF > $HOME/.aws/credentials
[aws_acct_1]
aws_access_key_id=50M353C237!D
aws_secret_access_key=50M353C237P@55W02D

[aws_acct_2]
aws_access_key_id=50M353C237!D
aws_secret_access_key=50M353C237P@55W02D
EOF
```

You can add as many accounts as you want.  I've tested it with like 12(?).  If you manage more than 12 AWS accounts, I'm impressed and horrified all at the same time.

In the square brackets of the credentials file is the name of the AWS account.  And by "name" I mean whatever name you want to call it.  You can call it "prod" or "banana_phone"... it doesn't matter, and it doesn't correspond to anything in your AWS account.  It's completely arbitrary.  But it will show up in the "Account" column in CloudTables for every resource in that account, so try to pick a good identifier.

You can also change it later and re-sync to update the account name on those objects.

### Docker Compose

Provided is a simple docker compse file that will install the front end, back end, and DB containers on a standalone docker host.

I have scripts to install docker and docker-compose here:

* [CentOS 7](https://github.com/corysm1th/centos7_scripts/tree/master/docker)
* [Ubuntu 16.04](https://github.com/corysm1th/ubuntu_scripts/tree/master/docker)

First, clone the repository

```sh
git clone https://github.com/corysm1th/cloudtables-python.git
cd cloudtables-python
```

If you're just doing a PoC, you can `make self_signed` to generate the certs.  Otherwise, place your certificate and key somewhere on the server, and create symlinks for `cloudtables/ssl/nginx.crt` and `cloudtables/ssl/nginx.key` inside the repo.

```sh
cd cloudtables-python/ssl
ln -s /path/to/yourcert.crt nginx.crt
ln -s /path/to/yourcert.key nginx.key
```

Finally, `cd` back to the parent `cloudtables-python` repo folder and install the software:

```sh
make install
```

The front end will be exposed on port 443 of the host, and the inital sync should kick off automatically.