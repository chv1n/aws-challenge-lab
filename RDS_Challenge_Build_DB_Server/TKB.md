# AWS RDS & EC2 Troubleshooting Knowledge Base

## Issue #1: Security and Compliance - Network Timeout
**Issue Description:** Unable to connect to Amazon RDS instance from an EC2 instance.

* **Symptoms:**
    * Connection hangs indefinitely.
    * Ends with a `Connection timed out` error message.
* **Root Cause Analysis:**
    * The RDS Security Group does not have an Inbound Rule to allow traffic on port 3306 from the EC2 instance's Private IP or its Security Group ID.
* **Resolution Procedures:**
    1.  Navigate to **RDS Console** > **Connectivity & security**.
    2.  Click on the active **VPC security groups**.
    3.  Select **Edit inbound rules**.
    4.  Add a new rule: **Type:** `MySQL/Aurora`, **Port:** `3306`.
    5.  Set **Source** to the EC2's Security Group ID (Best Practice) or specific IP range.
* **Helpful Tools or Resources:**
    * AWS Management Console
    * `nc -zv <endpoint> 3306` (Netcat) to verify port availability.
* **Comments:** Always verify the network path and security group rules before troubleshooting database-level credentials.

---

## Issue #2: Storage and Data Management - Parameter Group & Authentication
**Issue Description:** Unable to apply database configuration changes (e.g., changing authentication plugins) despite updating the Parameter Group settings.

* **Symptoms:**
    1.  Database still returns `ERROR 2059 (HY000)` after updating the Parameter Group.
    2.  The connection continues to fail with `caching_sha2_password` error even though `mysql_native_password` was set.
* **Root Cause Analysis:**
    1.  **Static vs. Dynamic:** Some parameters are "Static," meaning they require a manual reboot to take effect.
    2.  **Pending-Reboot:** The RDS instance status shows "pending-reboot" after the Parameter Group is associated.
    3.  **Cached Credentials:** Existing users (like `admin`) remain mapped to the old authentication plugin until their password is re-hashed.
* **Resolution Procedures:**
    1.  **Create Custom Group:** Since the `default` group is read-only, create a custom Parameter Group.
    2.  **Modify Parameter:** Set `default_authentication_plugin` to `mysql_native_password`.
    3.  **Apply & Reboot:** Attach the new group to the RDS instance and perform a **Manual Reboot**.
    4.  **Force Password Re-hash:** Use **Modify RDS** to change the Master Password (even to the same one) to force the DB to store it using the new plugin.
* **Helpful Tools or Resources:**
    * AWS RDS Console (Parameter Groups tab)
    * SQL Command: `SHOW VARIABLES LIKE 'default_authentication_plugin';`
* **Comments:** Remember that updating a Parameter Group is a two-step process: **Apply Changes** and then **Manual Reboot**.

---