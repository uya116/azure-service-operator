# Azure SQL Operator

## Resources Supported

The Azure SQL operator suite consists of the following operators.

1. Azure SQL server - Deploys an Azure SQL server given the location and Resource group
2. Azure SQL database - Deploys an SQL database given the SQL server
3. Azure SQL firewall rule - Deploys a firewall rule to allow access to the SQL server from specific IPs
4. Azure SQL Action - Allows you to roll the password for the specified SQL server
5. Azure SQL failover group - Deploys a failover group on a specified Azure SQL server given the secondary server and the databases to failover
6. Azure SQL User - Creates an user on the specified Azure SQL database and stores the username/password as secrets

## Deploying SQL Resources

You can follow the steps [here](/docs/development.md) to either run the operator locally or in a real Kubernetes cluster.

You can use the YAML files in the `config/samples` folder to create the resources.

**Note**  Don't forget to set the Service Principal ID, Service Principal secret, Tenant ID and Subscription ID as environment variables

### Azure SQL server

For instance, this is the sample YAML for the Azure SQL server.

  ```yaml
    apiVersion: azure.microsoft.com/v1alpha1
    kind: AzureSqlServer
    metadata:
    name: sqlserver-sample
    spec:
     location: westus
     resourcegroup: resourceGroup1
  ```

The value for kind, `AzureSqlServer` is the Custom Resource Definition (CRD) name.
`sqlserver-sample` is the name of the SQL server resource that will be created.

The values under `spec` provide the values for the location where you want to create the Azure SQL server at and the Resource group in which you want to create it under.

Once you've updated the YAML with the settings you need, and you have the operator running, you can create a Custom SQL server resource using the command.

```bash
kubectl apply -f config/samples/azure_v1_sqlserver.yaml
```

Along with creating the SQL server, this operator also generates the admin username and password for the SQL server and stores it in a kube secret with the same name as the SQL server.

This secret contains the following fields.

- `fullyqualifiedservername` : Fully qualified name of the SQL server such as sqlservername.database.windows.net
- `sqlservername` : SQL server name
- `username` : Server admin
- `password` : Password for the server admin
- `fullyqualifiedusername` : Fully qualified user name that is required by some apps such as <username>@<sqlserver>

You can retrieve this secret using the following command for the sample YAML

```bash
kubectl get secret sqlserver-sample -o yaml
```

This would show you the details of the secret. `username` and `password` in the `data` section are the base64 encoded admin credentials to the SQL server.

```bash
apiVersion: v1alpha1
data:
  fullyqualifiedservername: c3Fsc2VydmVyLXNhbXBsZS04ODguZGF0YWJhc2Uud2luZG93cy5uZXQ=
  fullyqualifiedusername: aGFzMTUzMnVAc3Fsc2VydmVyLXNhbXBsZS04ODg=
  password: XTdpMmQqNsd7YlpFdEApMw==
  sqlservername: c3Fsc2VyfmVyLXNhbXBsZS04ODg=
  username: aGFzMTFzMnU=
kind: Secret
metadata:
  creationTimestamp: "2019-10-09T21:02:02Z"
  name: sqlserver-sample-888
  namespace: default
  ownerReferences:
  - apiVersion: azure.microsoft.com/v1
    blockOwnerDeletion: true
    controller: true
    kind: AzureSqlServer
    name: sqlserver-sample-888
    uid: 08fdbf42-ead8-11e9-91e0-025000000001
  resourceVersion: "131163"
  selfLink: /api/v1/namespaces/default/secrets/sqlserver-sample-888
  uid: 0aeb2429-ead8-11e9-91e0-025000000001
type: Opaque
```

### SQL Database

Below is the sample YAML for SQL database

```yaml
apiVersion: azure.microsoft.com/v1alpha1
kind: AzureSqlDatabase
metadata:
  name:  sqldatabase-sample
spec:
  location: westus
  resourcegroup: ResourceGroup1
  # Database Editions
  # Basic=0; Business=1; BusinessCritical=2; DataWarehouse=3; Free=4;
  # GeneralPurpose=5; Hyperscale=6; Premium=7; PremiumRS=8; Standard=9;
  # Stretch=10; System=11; System2=12; Web=13
  edition: 0
  server:  sqlserver-sample
```

Update the `location` and the `resourcegroup` to where you want to provisiong the SQL database. `server` is the name of the Azure SQL server where you want to create the database in.
The `edition` represents the SQL database edition you want to use when creating the resource and should be one of the values above.

### SQL firewall

The SQL firewall operator allows you to add a SQL firewall rule to the SQL server.

Below is the sample YAML for SQL firewall rule

```yaml
apiVersion: azure.microsoft.com/v1alpha1
kind: AzureSqlFirewallRule
metadata:
  name: sqlf-allowazuresvcaccess
spec:
  resourcegroup: ResourceGroup1
  server:  sqlserver-sample
  
  # this IP range enables Azure Service access
  startipaddress: 0.0.0.0
  endipaddress: 0.0.0.0
```

The `server` indicates the SQL server on which you want to configure the new SQL firewall rule on and `resourcegroup` is the resource group of the SQL server. The `startipaddress` and `endipaddress` indicate the IP range of sources to allow access to the SQL server.

When the `startipadress` and `endipaddress` are 0.0.0.0, it is a special case that adds a firewall rule to allow all Azure services to access the SQL server.

### SQL Action

The SQL Action operator is used to trigger an action on the SQL server. Right now, the only action supported is `rollcreds` which rolls the password for the SQL server to a new one.

Below is a sample YAML for rolling the password

```yaml
apiVersion: azure.microsoft.com/v1alpha1
kind: AzureSqlAction
metadata:
  name: Sql-rollcreds-action
spec:
  resourcegroup: ResourceGroup1
  actionname: rollcreds
  servername: sqlserver-sample
```

The `name` is a name for the action that we want to trigger. The type of action is determined by the value of `actionname` in the spec which is `rollcreds` if you want to roll the password (Note: This action name should be exactly `rollcreds` for the password to be rolled). The `resourcegroup` and `servername` identify the SQL server on which the action should be triggered on.

Once you apply this, the kube secret with the same name as the SQL server is updated with the rolled password.

### SQL failover group

The SQL failover group operator is used to create a failover group on a specified primary Azure SQL server, given the secondary Azure SQL server (should be in a different location from the primary server) and the databases on the primary server that should failover.

Below is a sample YAML for creating a failover group

```yaml
apiVersion: azure.microsoft.com/v1alpha1
kind: AzureSqlFailoverGroup
metadata:
  name: azuresqlfailovergroup-sample
spec:
  location: eastus
  resourcegroup: resourcegroup-azure-operators
  server: sqlserver-samplepri
  failoverpolicy: automatic
  failovergraceperiod: 30
  secondaryserver: sqlserver-samplesec
  secondaryserverresourcegroup: resourcegroup-azure-operators
  databaselist:
    - "azuresqldatabase-sample"
    - "azuresqldatabase-sample2"
```

The `name` is a name for the failover group that we want to create. `server` is the primary SQL server on which the failover group is created, `location` and `resourcegroup` are the location and the resource group of the primary SQL server. `failoverpolicy` can be "automatic" or "manual". `failovergraceperiod` is the time in minutes. `secondaryserver` is the secondary SQL server to failover to and `secondaryserverresourcegroup` is the resource group that the server is in. `databaselist` is the list of databased on the primary SQL server that should replicate to the secondary SQL server, when there is a failover action.

Once you apply this, a secret with the same name as the SQL failovergroup is also stored. This secret contains the fields for primary/secondary failovergroup listener endpoints (`readwritelistenerendpoint` and `readonlylistenerendpoint`) and the primary/secondary SQL server names (`azuresqlprimaryservername` and `azuresqlsecondaryservername`)

### SQL database user

The SQL user operator is used to create a user on the specified Azure SQL database. This user is more restrictive than the admin user created on the SQL server and is so recommended to use. The operator creates the user on the database by auto generating a strong password, and also stores the username and password as a secret (name can be specified in the YAML), so applications can use them.

Below is a sample YAML for creating a database user

```yaml
apiVersion: azure.microsoft.com/v1alpha1
kind: AzureSQLUser
metadata:
  name: sqluser-sample
spec:
  server: sqlserver-sample-777
  dbname: azuresqldatabase-sample
  adminsecret: sqlserver-sample-777
  # possible roles:
  # db_owner, db_securityadmin, db_accessadmin, db_backupoperator, db_ddladmin, db_datawriter, db_datareader, db_denydatawriter, db_denydatareader
  roles:
    - "db_owner"
```

The `name` is used to generate the username on the database. The exact name is not used but rather a UUID is appended to this to make it unique. `server` and `dbname` qualify the database on which you want to create the user on. `adminsecret` is the name of the secret where the username and password will be stored. `roles` specify the security roles that this user should have on the specified database.

Once you apply this, a secret with the name specified in `adminsecret` is stored with the following fields.

- `username` : Username of user created on the database
- `password` : Password for the user
- `sqlservernamespace` : Kube namespace where the SQL server is provisioned
- `sqlservername` : SQL server name

## View and Troubleshoot SQL Resources

To view your created SQL resources, such as sqlserver, run the following command:

```shell
kubectl get <CRD>
```

where CRD is the Custom Resource Definition name or `Kind` for the resource.

For instance, you can get the Azure SQL servers provisioned using the command

```shell
kubectl get AzureSqlServer
```

You should see the AzureSqlServer instances as below

```shell
NAME                  AGE
sqlserver-sample      1h
```

If you want to see more details about a particular resource instance such as the `Status` or `Events`, you can use the below command

```shell
kubectl describe <Kind> <instance name>
```

For instance, the below command is used to get more details about the `sqlserver-sample` instance

```shell
kubectl describe AzureSqlServer sqlserver-sample
```

```shell
Name:         sqlserver-sample234
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"azure.microsoft.com/v1alpha1","kind":"AzureSqlServer","metadata":{"annotations":{},"name":"sqlserver-sample234","namespace":"default"}...
API Version:  azure.microsoft.com/v1alpha1
Kind:         SqlServer
Metadata:
  Creation Timestamp:  2019-09-26T21:30:56Z
  Finalizers:
    azuresqlserver.finalizers.azure.com
  Generation:        1
  Resource Version:  20001
  Self Link:         /apis/azure.microsoft.com/v1/namespaces/default/azuresqlservers/sqlserver-sample234
  UID:               ed1c5d1d-e0a4-11e9-9ee8-52a5c765e9d7
Spec:
  Location:                 westus
  Resourcegroup:            resourceGroup1
Status:
  Provisioned:  true
  State:        Ready
Events:
  Type    Reason       Age                   From                  Message
  ----    ------       ----                  ----                  -------
  Normal  Updated      2m21s                 SqlServer-controller  finalizer azuresqlserver.finalizers.azure.com added
  Normal  Submitting   2m21s                 SqlServer-controller  starting resource reconciliation
  Normal  Checking     108s (x3 over 2m18s)  SqlServer-controller  instance in NotReady state
  Normal  Checking     76s (x2 over 78s)     SqlServer-controller  instance in Ready state
  Normal  Provisioned  75s (x2 over 76s)     SqlServer-controller  azuresqlserver sqlserver-sample234 provisioned
```

The `Status` section gives you the current state of the resource that it's Ready and is provisioned.

The `Events` have a chronological record of what occurred through the process of provisioning the resource.

## Delete a SQL Resource

To delete an existing resource from Kubernetes and Azure, use the following command.

```shell
kubectl delete <Kind> <instancename>
```

For instance, deleting the above SqlServer instance would look like this.

```shell
kubectl delete AzureSqlServer sqlserver-sample
```

The following message should appear:

`azuresqlserver.azure.microsoft.com sqlserver-sample deleted.`

## Demo

Watch this demo <https://bit.ly/2lUIX6Y> to observe how you would you use the Azure SQL Operator from a real application.

In this demo, we use YAML to deploy the application and the SQL server. Once the SQL server is provisioned, the connection details are stored as secrets which the application can use to access the SQL server.