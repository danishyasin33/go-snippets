This project contains various go snippets as an example of my exemplry go skills. 

## Utils
Following is an explanation of each file: 
1. ec2.go
Contains the function `Ec2Address` which can be used to get the ec2 local ipv4. This function can only be executed on an EC2 instance in AWS. 
2. s3.go 
Contains `DownloadFromS3` which takes in a bucket and an item name and tries to download that object in `us-east-1` region.
3. utils.go
Contains various small functions that gets some specific job done. Each function has a corresponding comment that explains its use


## GRPC
It's a sample GRPC service that has a go service that communicates with a python grpc server to handle a `blurBackground` request. The python script uses a ML model to detect the background of an image and then blur behind it. 
The complexity of the blurring is hidden behind an `infer` method that is not included here. 

## TODO
1. Add tests to make it easier to check the functions