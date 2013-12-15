# notify-by-opsgenie

## Overview
Simple script to notify about nagios alerts via opsgenie

## Usage
The script can be called standalone like this:
```
% ./notify-by-opsgenie --help
Usage of ./notify-by-opsgenie:
  -H="": the host to alert for
  -S="": the service to alert for
  -T="": time and date in long format
  -c="opsgenie_key": the opsgenie customer key
  -d="": description of the alert
  -s="": the state
  -t="": the alert type
  -u="": recipient to send to
```
But you probably want to use it in Nagios. For this define commands like this:
```
define command{
  command_name	notify-host-opsgenie
  command_line  $USER1$/notify-by-opsgenie -H="$HOSTNAME$" -s="$HOSTSTATE$" -c="OPSGENIEKEY" -t="$NOTIFICATIONTYPE$" -d="$HOSTOUTPUT$" -T="$LONGDATETIME$" -u="$CONTACTEMAIL$"
}

define command{
  command_name	notify-service-opsgenie
  command_line  $USER1$/notify-by-opsgenie -H="$HOSTNAME$" -S="$SERVICEDESC$" -s="$SERVICESTATE$" -c="OPSGENIEKEY" -t="$NOTIFICATIONTYPE$" -d="$SERVICEOUTPUT$" -T="$LONGDATETIME$" -u="$CONTACTEMAIL$"
}
```
And then set up a contact to use the commands for notifications like this:
```
define contact{
        contact_name                   opsgenie
        use                            generic-contact
        alias                          Opsgenie contact
        email                          your@email.com
        service_notification_commands  notify-service-opsgenie
        host_notification_commands     notify-host-opsgenie
        }
```

## How to contribute
1. Fork the repo
2. Hack away
3. Push the branch up to GitHub
4. Send a pull request

