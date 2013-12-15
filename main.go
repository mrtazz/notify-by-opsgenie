// simple notification script to use opsgenie from nagios
//
//     notify-by-opsgenie -H=$HOSTNAME$ -s=$HOSTSTATE$ -c=$CUSTOMERKEY -t=$NOTIFICATIONTYPE$ -d=$HOSTOUTPUT$ -T=$LONGDATETIME$ -u=$CONTACTEMAIL$
//     notify-by-opsgenie -H=$HOSTNAME$ -S=$SERVICEDESC$ -s=$SERVICESTATE$ -c=$CUSTOMERKEY -t=$NOTIFICATIONTYPE$ -d=$SERVICEOUTPUT$ -T=$LONGDATETIME$ -u=$CONTACTEMAIL$
//
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type AlertDetails map[string]string
type Alert map[string]interface{}

func main() {
	// note, that variables are pointers
	host := flag.String("H", "", "the host to alert for")
	service := flag.String("S", "", "the service to alert for")
	state := flag.String("s", "", "the state")
	key := flag.String("c", "opsgenie_key", "the opsgenie customer key")
	alert_type := flag.String("t", "", "the alert type")
	description := flag.String("d", "", "description of the alert")
	longdate := flag.String("T", "", "time and date in long format")
	recipient := flag.String("u", "", "recipient to send to")

	flag.Parse()

	var label, message string
	if *service != "" {
		label = fmt.Sprintf("%s/%s", *host, *service)
		message = fmt.Sprintf("%s on %s is %s", *service, *host, *state)
	} else {
		label = *host
		message = fmt.Sprintf("%s is %s", *host, *state)
	}

	a := Alert{"message": message, "recipients": []string{*recipient}, "alias": label,
		"customerKey": *key, "note": fmt.Sprintf("%s\nAlert Date: %s",
			*description, *longdate)}
	switch *alert_type {
	case "PROBLEM":
		create_alert(a)
	case "RECOVERY":
		close_alert(a)
	case "ACKNOWLEDGEMENT":
		ack_alert(a)
	}

}

// helper functions
func create_alert(a Alert) error {
	data, _ := json.Marshal(a)
	err := http_post("alert", data)
	return err
}

func close_alert(a Alert) error {
	data, _ := json.Marshal(a)
	err := http_post("alert/close", data)
	return err
}

func ack_alert(a Alert) error {
	data, _ := json.Marshal(a)
	err := http_post("alert/acknowledge", data)
	return err
}

func http_post(urlpath string, buf []byte) error {
	body := bytes.NewBuffer(buf)
	resp, err := http.Post("https://api.opsgenie.com/v1/json/"+urlpath,
		"application/json", body)
	defer resp.Body.Close()
	return err
}
