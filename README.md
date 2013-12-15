# notify-by-opsgenie

## Overview
Simple script to notify about nagios alerts via opsgenie

## Usage

```
notify-by-opsgenie -H=$HOSTNAME$ -s=$HOSTSTATE$ -c=$CUSTOMERKEY -t=$NOTIFICATIONTYPE$ -d=$HOSTOUTPUT$ -T=$LONGDATETIME$ -u=$CONTACTEMAIL$
notify-by-opsgenie -H=$HOSTNAME$ -S=$SERVICEDESC$ -s=$SERVICESTATE$ -c=$CUSTOMERKEY -t=$NOTIFICATIONTYPE$ -d=$SERVICEOUTPUT$ -T=$LONGDATETIME$ -u=$CONTACTEMAIL$
```

## How to contribute
1. Fork the repo
2. Hack away
3. Push the branch up to GitHub
4. Send a pull request

