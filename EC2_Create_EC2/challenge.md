
# Creating Amazon EC2 Instances
#### Lab overview
AWS provides multiple ways to launch Amazon Elastic Compute Cloud (Amazon EC2) instance. 

In this lab, you use the AWS Management Console to launch an EC2 instance and then use it as a bastion host to launch another EC2 instance, which will be a web server. You use EC2 Instance Connect to securely connect to the bastion host and use the AWS Command Line Interface (AWS CLI) to launch a web server instance.

The following diagram illustrates the final architecture that you will build:
![image.png](https://raw.githubusercontent.com/bucketio/img13/main/2026/03/16/1773644308139-29534b9b-88c0-42e0-bb4e-3ffeec32eb4a.png )

#### Objectives

After completing this lab, should be able to do the following:

- Launch an EC2 instance by using the AWS Management Console.
- Connect to the EC2 instance by using EC2 Instance Connect.
- Launch an EC2 instance by using the AWS CLI.