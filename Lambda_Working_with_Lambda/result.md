## **Task 1: Observing the salesAnalysisReport,salesAnalysisReportDERole IAM role settings**
`Analyze the IAM roles and the permissions that they grant to the salesAnalysisReport and salesAnalysisReportDataExtractor Lambda functions that you create later.`

>**SalesAnalysisReport**
![alt text](./img/image-5.png)
![alt text](./img/image-1.png)
>- **AmazonSNSFullAccess** provides full access to Amazon SNS resources.
>- **AmazonSSMReadOnlyAccess** provides read-only access to Systems Manager resources.
>- **AWSLambdaBasicRunRole** provides write permissions to CloudWatch logs (which are required by every Lambda function).
>- **AWSLambdaRole** gives a Lambda function the ability to invoke another Lambda function. 

>**SalesAnalysisReportDERole**
![alt text](./img/image-4.png)
![alt text](./img/image-3.png)
>- AWSLambdaBasicRunRole provides write permissions to CloudWatch logs.
>- AWSLambdaVPCAccessRunRole provides permissions to manage elastic network interfaces to connect a function to a virtual private cloud (VPC).


## **Task 2: Creating a Lambda layer and a data extractor Lambda function**
`in this task, first create a Lambda layer, and then create a Lambda function that uses the layer.`

>Creating a Lambda Layer
![alt text](./img/image-6.png)

>Creating a data extractor Lambda function and using the layer.
![alt text](./img/image-7.png)
![alt text](./img/image-8.png)

>Importing the code for the data extractor Lambda function
![alt text](./img/image-9.png)

>Configuring network settings for the function
![alt text](./img/image-10.png)

## **Task 3: Testing the data extractor Lambda function**
`The body field, which contains the report data that the function extracted, is empty because there is no order data in the database`
![alt text](./img/image-11.png)

>Placing an order and testing again
>
>`Access the café website and place some orders to populate data in the database. `
>![alt text](./img/image-12.png)
>![alt text](./img/image-13.png)
>`The returned JSON object now contains product quantity information in the body field similar to the following:`
![alt text](./img/image-15.png)


## **Task 4: Configuring notifications**
`Create an SNS topic and then subscribe an email address to the topic`

>Creating an SNS topic
![alt text](./img/image-16.png)

>Subscribing to the SNS topic
![alt text](./img/image-18.png)

>Confirm subscription
>
>`Check the inbox for the email address. I see an email from SARTopic with the subject "AWS Notification - Subscription Confirmation.`
>![alt text](./img/image-19.png)
>![alt text](./img/image-20.png)



## **Task 5: Creating the salesAnalysisReport Lambda function**

Create and configure the salesAnalysisReport Lambda function. This function is the main driver of the sales analysis report flow. It does the following:
- Retrieves the database connection information from Parameter Store
- Invokes the salesAnalysisReportDataExtractor Lambda function, which retrieves the report data from the database
- Formats and publishes a message containing the report data to the SNS topic

>Connecting to the CLI Host instance
![alt text](./img/image-21.png)

>Configuring the AWS CLI
![alt text](./img/image-23.png)

>Creating the salesAnalysisReport Lambda function using the AWS CLI
![alt text](./img/image-24.png)

>Configuring the salesAnalysisReport Lambda function
![alt text](./img/image-25.png)

>Testing the salesAnalysisReport Lambda function
>![alt text](./img/image-26.png)
>
>`I receive an email from AWS Notifications with the subject "Daily Sales Analysis Report." `
![alt text](./img/image-27.png)

>Adding a trigger to the salesAnalysisReport Lambda function
`Configure the report to be initiated Monday through Saturday at 8 PM each day`
![alt text](./img/image-28.png)


### Finished  !! 🎉