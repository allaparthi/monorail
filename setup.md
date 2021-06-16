#### Commands to set this up and deploy:

##### Download and install google-cloud-sdk
_Note: Install latest version in wget command_

```
wget google-cloud-sdk-303.0.0-darwin-x86_64.tar.gz 
./google-cloud-sdk/install.sh
source ~/.bash_profile
gcloud auth login
gcloud config set project monorail-280007
```

##### For monorail:
```
cd $HOME && git clone https://chromium.googlesource.com/chromium/tools/depot_tools.git
export PATH=$PATH:$HOME/depot_tools
mkdir monorail_local && cd monorail_local/
fetch infra && cd infra
git remote set-url origin git@github.com:epiFi/monorail.git
git fetch
git reset --hard origin/master
cd appengine/monorail/
eval `../../go/env.py`
chmod -R 755 ~/.vpython-root/
`which python` `which pip` install six
cd ~/monorail_local/infra/luci/appengine/components/components/
`which python` `which pip` install sendgrid -t third_party/
cd ~/monorail_local/appengine/monorail/
make deploy_prod
```


##### For subsequent deploys: 
- You can run `make deploy_prod`. 
Make sure the repo wherever is this command running from, is in total sync with remote repo. 
Any local commits, or uncommitted changes etc will lead to deploying `tainted-<your_machine_name>-<random_num>.appspot.com` which isn't a good production deployment.

##### For updating with latest remote repo:

- Look out for architectural changes or new infrastructure components before you update.

- Make a local copy of file appengine/monorail/features/notify.py. We've added following code for sendgrid mail support.
```
from sendgrid import SendGridAPIClient
from sendgrid.helpers.mail import Mail as sg_mail
from sendgrid.helpers.mail import To as sg_to
from sendgrid.helpers.mail import From as sg_from
from sendgrid.helpers.mail import Content as sg_content
......
class OutboundEmailTask(jsonfeed.InternalTask):
  def HandleRequest(self, mr):
    ......
    else:
        sg = SendGridAPIClient(os.environ['SENDGRID_API_KEY'])
        retry_count = 3
        logging.info('build_sendgrid_message')
        sendgrid_from = sg_from(sender)
        sendgrid_to = sg_to(to)
        sendgrid_content = sg_content("text/plain", body)
        sendgrid_message = sg_mail(
        from_email=sendgrid_from, to_emails=sendgrid_to, subject=subject, html_content=sendgrid_content)
        logging.info('sendgrid_message: %s', sendgrid_message)

        for i in range(retry_count):
          try:
            sg_response = sg.send(sendgrid_message)
            logging.info('viaSendgrid status_code: %s, body: %s, headers: %s', sg_response.status_code, sg_response.body, sg_response.headers)
            #message.send()
            break
          except apiproxy_errors.DeadlineExceededError as ex:
            logging.warning('Sending email timed out on try: %d', i)
            #logging.warning(str(ex))
            logging.warning(e.message)
```

- Make a copy of changes in following files:
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
