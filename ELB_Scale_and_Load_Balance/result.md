## **Task 1: Creating an AMI for auto scaling**
`Create an AMI from the existing Web Server 1.`
![alt text](./img/image-2.png)


## **Task 2: Create a target group**
![alt text](./img/image-3.png)


## **Task 3: Create a Load Balance**
![alt text](./img/image-4.png)

## **Task 4: Creating a launch template**
`Create a launch template for Auto Scaling group. A launch template is a template that an Auto Scaling group uses to launch EC2 instances.`
![alt text](./img/image-5.png)

## **Task 5: Creating an Auto Scaling group**
`Use a launch template to create an Auto Scaling group, which is configured to maintain an average CPU utilization of 50%.It automatically adds or removes EC2 instances as needed to keep the system's performance balanced and cost-effective`
![alt text](./img/image-6.png)

## **Task 6: Verifying that load balancing is working**
![alt text](./img/image-8.png)

`Registered Target`
![alt text](./img/image-7.png)

`Load Test`
![alt text](./img/image-9.png)

`Monitoring Dashboard`
![alt text](./img/image-10.png)


## **Task 7: Testing auto scaling**
`Target Tracking Alarm in the 'In alarm' state, indicating that average CPU utilization has exceeded the 50% threshold and is currently triggering the Auto Scaling Group to add more instances`
![alt text](./img/image-12.png)

![alt text](./img/image-11.png)


### Finished  !! 🎉