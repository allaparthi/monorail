Source: https://bugs.chromium.org/p/monorail/issues/attachmentText?aid=450747

# Provision CloudSQL databases

_If you're only updating an instance you can skip this section._

Create a master `mysql v5.6` db instance in cloud sql with the name: `your_company-monorail-master-01`

*__NOTE__: Do NOT set a password on the DB*

Set the following flags on the DB using the GCP console:
`sql_mode = ERROR_FOR_DIVISION_BY_ZERO NO_AUTO_CREATE_USER NO_ENGINE_SUBSTITUTION`

Then create two read replicas of the master database, they must have the following names: `your_company-monorail-master-replica-01` and `your_company-monorail-master-replica-01`

Turn on high availability for the master DB.

# Create an AppEngine app

_If you're only updating an instance you can skip this section._
Goto cloud console for AppEngine and create an APP in the us-central region.

# Setup virtual environment

Now SSH to the a clean debian instance you want to build from. Here is where you'll stage monorail for deployment. 

Run the following commands:

`sudo apt-get install wget automake make g++ python-mysqldb python-dev git`

`mkdir mysql_packages && cd mysql_packages`

`wget https://dev.mysql.com/get/Downloads/MySQL-5.6/mysql-server_5.6.40-1ubuntu14.04_amd64.deb-bundle.tar`

`tar -xf mysql-server_5.6.40-1ubuntu14.04_amd64.deb-bundle.tar`

`sudo dpkg --install mysql-common*`

```
sudo apt install ./mysql-community-client_5.6.40-1ubuntu14.04_amd64.deb --fix-broken`
sudo apt install ./mysql-client_5.6.40-1ubuntu14.04_amd64.deb --fix-broken
sudo apt install ./libmysqlclient-dev_5.6.40-1ubuntu14.04_amd64.deb ./libmysqlclient18_5.6.40-1ubuntu14.04_amd64.deb
```

`cd $HOME && git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git`

`export PATH=$PATH:$HOME/depot_tools`

`mkdir working && cd working`

`fetch infra && cd infra`

`gclient sync`

`cd appengine/monorail`


Edit the following files and make sure they make sense for you. Pay particular attention the size and # of idle instances in each of the m4 files.
If you run monorail with it's shipped defaults, you'll spend well over a thousand dollars a day!!!!

Edit:

```
Makefile
app-cloud.yaml.m4
app.yaml.m4
module-api-cloud.yaml.m4
module-api.yaml.m4
module-besearch-cloud.yaml.m4
module-besearch.yaml.m4
module-latency-insensitive-cloud.yaml.m4
module-latency-insensitive.yaml.m4
settings.py
```

In the m4 files any line that references `instance_class` should probably set to `F1` at first and any line that references `min_idle_instances` should be set to `1` as well.
In `Makefile` set PRODID to the GAE AppID of your project, update DEVID and STAGEID if necessary as well.
Then update everything in `settings.py` 

You must setup your virtual monorail dev environment, so run this command
```eval `../../go/env.py` ```

The virtual environment has permissions that won't let the next step succeed, so fix that now:
`chmod -R 755 ~/.vpython-root/`

You must now install some dependencies into our virtual environment, so run this exactly:
``` `which python` `which pip` install six MySQL-Python ```

**If you're only updating an existing monorail deployment, then skip down to the section below: [To Deploy](#todeploy)**

If you're setting up a new deployment, then next we must setup the database tables.

Use the GCP Console to whitelist the IP address of your current machine so that it can communicate with your Cloud SQL database, in my case the IP address was: `35.188.90.49` in your case it will be different, so replace it as needed in the steps below.
 ``` 
DBIP=35.188.90.49
mysql --user=root -h $DBIP -e 'CREATE DATABASE monorail;'
mysql --user=root -h $DBIP monorail < schema/framework.sql
mysql --user=root -h $DBIP monorail < schema/project.sql
mysql --user=root -h $DBIP monorail < schema/tracker.sql
```
  
# Final GCP configuration

_If you're only updating an instance you can skip this section._

Next, turn on the IAM Service Account Credentials API at this URL:
https://console.developers.google.com/apis/api/iamcredentials.googleapis.com/overview


# AppEngine GSuite authentication

_If you're only updating an instance you can skip this section._

The following steps must be executed in this specific order...
__FIRST__ turn on authentication for GAE to use GSuite auth for your_company.com
__SECOND__ map bugs.your_company.io to the GAE app in the appengine settings panel.


We need to create an empty ClientConfig  in CloudDatastore, goto: https://console.cloud.google.com/datastore/entities

1. Create an entity called ClientConfig in the default namespace.
1. It should have a custom name identifier.
1. The custom name must be: api_client_configs
1. Give it a property called: configs
1. Set the type to Text.
1. Click create.

To make email work with monorail, make `monorail@your_company.com` as a Google Group via the GSuite admin console.  Make sure it accepts emails from the Internet (everyone).


# <a name="todeploy"></a>To Deploy

Then, to build monorail and push it to our prod instance: `make deploy_prod` 

If this is a new instance, log into the site as yourself and create an account. Then use the following command to make yourself an admin. Future admins can be made in the UI.

```mysql --user=root -h $DBIP monorail -e "UPDATE User SET is_site_admin = TRUE WHERE email = 'YOUR_EMAIL_GOES_HERE@your_company.com'"```


# Important Cleanup
Go back to the Cloud SQL console and unwhitelist your IP address.

About Monorail User Guide Release Notes Feedback on Monorail Terms Privacy
