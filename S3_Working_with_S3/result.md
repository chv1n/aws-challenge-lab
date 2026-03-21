## **Task 1: Connecting to the CLI Host EC2 instance and configuring the AWS CLI**  
![alt text](./img/image-1.png)

## **Task 2: Creating and initializing the S3 share bucket**  
In this task, using the AWS CLI to create the S3 share bucket and upload a few images. 

>Create bucket  
![alt text](./img/image-2.png)

>Load image into the bucket
![alt text](./img/image-3.png)

>Verify that the files ware synced to the S3 Bucket
![alt text](./img/image-4.png)

## **Task 3: Reviewing the IAM group and user permissions** 

>Reviewing the mediaco IAM group
![alt text](./img/image-5.png)

>Reviewing the mediacouser IAM user 
![alt text](./img/image-6.png)

>Create Access Key
![alt text](./img/image-14.png)

## **Task 4: Configuring event notifications on the S3 share bucket**
In this task, configure the S3 share bucket to generate an event notification to an SNS topic whenever the contents of the bucket change. The SNS topic then sends an email message to its subscribed users with the notification message.  
Perform the following steps:
- Create the s3NotificationTopic SNS topic.
- Grant Amazon S3 permission to publish to the topic.
- Subscribe to the topic.
- Add an event notification configuration to the S3 bucket.

>Creating and configuring the s3NotificationTopic SNS topic  
Edit Access policy
![alt text](./img/image-7.png)

Subscribe to the topic to recive the event notification.
![alt text](./img/image-9.png)

Confirm subscription
![alt text](./img/image-8.png)
![alt text](./img/image-10.png)

>Adding an event notification configuration to the S3 bucket
In this task, create an event notification configuration file that identifies the events that Amazon S3 will publish and the topic destination where Amazon S3 will send the event notifications. And then use the s3api CLI commands to associate this configuration file with the S3 share bucket.
![alt text](./img/image-11.png)

Associate the event configuration file with the S3 share bucket.   
`This command is " Setting S3 Bucket to save audit logs (Event Notification) Navigate" direction of travel to the next Bucket.`  
![alt text](./img/image-13.png)

`After that I received a notification email.`  
![alt text](./img/image-12.png)


## **Task 5: Testing the S3 share bucket event notifications**  
 In this task,  test the configuration of the S3 share bucket event notification by performing the use cases that mediacouser expects to perform on the bucket. These actions include putting objects into and deleting objects from the bucket, which send email notifications. Also test an unauthorized operation to verify that it is rejected. Use the AWS s3api CLI command to perform these operations on the S3 share bucket.

>Configure the CLI Host's AWS CLI client software to use the mediacouser credentials
![alt text](./img/image-15.png)
(Access key ID of mediacouser, Secret Access Key of mediacouser)

>Upload and delete this file  
` I tried uploading a file and then deleting the file. After that I got a notification.  `
![alt text](./img/image-17.png)
![alt text](./img/image-16.png)

### Finished  !! 🎉