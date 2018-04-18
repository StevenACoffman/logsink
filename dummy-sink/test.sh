#!/bin/bash

echo " ***** TESTING UN-structured logs with kubernetes metadata ****"
curl -v -X POST \
  http://127.0.0.1:3000/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: a16e943a-11d4-96e5-d538-1665eb121320' \
  -d '[{
		"date": 1521100395.933043,
		"log": "E0315 07:53:15.932716       5 leaderelection.go:228] error initially creating leader election record: configmaps is forbidden: User \"system:serviceaccount:kangaroo:default\" cannot create configmaps in the namespace \"kangaroo\"",
		"stream": "stderr",
		"time": "2018-03-15T07:53:15.933043339Z",
		"kubernetes": {
			"pod_name": "scoffman-nginix-release-nginx-ingress-controller-d2d6b",
			"namespace_name": "kangaroo",
			"pod_id": "788f971a-f529-11e7-b200-0a7c245249c8",
			"labels": {
				"app": "nginx-ingress",
				"component": "controller",
				"controller-revision-hash": "45824362",
				"pod-template-generation": "1",
				"release": "scoffman-nginix-release"
			},
			"annotations": {
				"checksum/config": "7452805dfd3eb2bd112637b75acaae9b35e38809877bfee430ccddddc5c49212",
				"logging.topic": "k8s-myTopic",
				"kubernetes.io/created-by": "{\\\"kind\\\":\\\"SerializedReference\\\",\\\"apiVersion\\\":\\\"v1\\\",\\\"reference\\\":{\\\"kind\\\":\\\"DaemonSet\\\",\\\"namespace\\\":\\\"kangaroo\\\",\\\"name\\\":\\\"scoffman-nginix-release-nginx-ingress-controller\\\",\\\"uid\\\":\\\"f6c2ed6a-e69b-11e7-9d12-123d868e15b6\\\",\\\"apiVersion\\\":\\\"extensions\\\",\\\"resourceVersion\\\":\\\"29392226\\\"}}\\n"
			},
			"host": "ip-172-28-149-29.ec2.internal",
			"container_name": "nginx-ingress-controller",
			"docker_id": "ad79aacac8c68e76480bfb1f639d0553af4069ca84c4cd0b0db9553062c2fd7f"
		}
	},{
		"date": 1521100395.933043,
		"log": "E0315 07:53:15.932716       5 leaderelection.go:228] error initially creating leader election record: configmaps is forbidden: User \"system:serviceaccount:kangaroo:default\" cannot create configmaps in the namespace \"kangaroo\"",
		"stream": "stderr",
		"time": "2018-03-15T07:53:15.933043339Z",
		"kubernetes": {
			"pod_name": "scoffman-nginix-release-nginx-ingress-controller-d2d6b",
			"namespace_name": "kangaroo",
			"pod_id": "788f971a-f529-11e7-b200-0a7c245249c8",
			"labels": {
				"app": "nginx-ingress",
				"component": "controller",
				"controller-revision-hash": "45824362",
				"pod-template-generation": "1",
				"release": "scoffman-nginix-release"
			},
			"annotations": {
				"checksum/config": "7452805dfd3eb2bd112637b75acaae9b35e38809877bfee430ccddddc5c49212",
				"logging.topic": "k8s-myTopic2",
				"kubernetes.io/created-by": "{\\\"kind\\\":\\\"SerializedReference\\\",\\\"apiVersion\\\":\\\"v1\\\",\\\"reference\\\":{\\\"kind\\\":\\\"DaemonSet\\\",\\\"namespace\\\":\\\"kangaroo\\\",\\\"name\\\":\\\"scoffman-nginix-release-nginx-ingress-controller\\\",\\\"uid\\\":\\\"f6c2ed6a-e69b-11e7-9d12-123d868e15b6\\\",\\\"apiVersion\\\":\\\"extensions\\\",\\\"resourceVersion\\\":\\\"29392226\\\"}}\\n"
			},
			"host": "ip-172-28-149-29.ec2.internal",
			"container_name": "nginx-ingress-controller",
			"docker_id": "ad79aacac8c68e76480bfb1f639d0553af4069ca84c4cd0b0db9553062c2fd7f"
		}
	}
	]'

echo " ***** TESTING structured logs without kubernetes metadata ****"
curl -v -X POST \
  http://127.0.0.1:3000/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 830daadc-515b-a55a-85cf-7c8910c3c5d4' \
  -d '[{"log":"awaiting signals\n","stream":"stdout","time":"2018-01-29T21:45:25.10786953Z"},{"log":"21:13:10.264 [http-bio-8888-exec-1] INFO  o.c.s.s.StatsdClient - {\"dests\":[\"DEFAULT\"],\"origin\":\"mylists-service\",\"host\":\"mylists-service01\",\"level\":\"INFO\",\"date\":\"2018-02-02 21:13:10\",\"ver\":\"1.1\",\"msg\":\"Stat sent: myjstor.userid.f01d4422-c9ff-4489-ab21-9baa671a9ca0:1|c\"}","stream":"stdout","time":"2018-01-29T21:45:55.107691376Z"},{"log":"{\"dests\":[\"k8s-pirate-message\"],\"_lb0\":{\"k8s-pirate-message\":\"07b33d65-f4b4-4095-9c21-95343e479eb7\"},\"eventid\":\"0cc0c074-2a2c-4f1c-8fd0-12fcfa612988\",\"requestid\":\"ad8ba3b3-27fe-4171-b1a1-d90c85989a1b\",\"origin\":\"local-pirate-spew.test\",\"eventtype\":\"foo\",\"tstamp_usec\":1517263525107326}\n","stream":"stdout","time":"2018-01-29T22:05:25.107703878Z"},{"log":"{\"dests\":[\"k8s-pirate-message\"],\"_lb0\":{\"k8s-pirate-message\":\"cce69a84-242b-4429-a43b-2cdc43f6b001\"},\"eventid\":\"176007c6-87e3-476e-ab00-a293276e4de7\",\"requestid\":\"33217961-0264-4976-8556-b4856a624d71\",\"origin\":\"local-pirate-spew.test\",\"eventtype\":\"foo\",\"tstamp_usec\":1517263555107339}\n","stream":"stdout","time":"2018-01-29T22:05:55.107728539Z"},{"log":"{\"dests\":[\"k8s-pirate-message\"],\"_lb0\":{\"k8s-pirate-message\":\"0228440b-ecc4-4541-9aaf-00a97685af1e\"},\"eventid\":\"0038566e-5fe8-43e5-8d44-0b60292458be\",\"requestid\":\"2d45a59f-6aa0-4cd1-912c-395627248ec1\",\"origin\":\"local-pirate-spew.test\",\"eventtype\":\"foo\",\"tstamp_usec\":1517263585107336}\n","stream":"stdout","time":"2018-01-29T22:06:25.107708785Z"},{"log":"{\"userids\": {\"1\": [\"2\", \"3\"]}, \"eventtype\": \"view_item\", \"origin\": \"lukes-test-service\", \"uri\": \"\/stable\/10.2307\/12345\", \"tstamp_usec\": 1400604584596834, \"dests\": [\"K8S_CAPTAINS_LOG\",\"YOWZA\"], \"requestid\": \"some-requestid:123456\", \"item_id\": \"4e92511e-a591-4ec4-876e-94274db4af07\", \"sessionid\":\"123454\", \"item_doi\":\"10.2307\/12345\"}","stream":"stdout","time":"2018-01-29T22:06:25.107708785Z"},{"log":"{\"eventtype\": \"barebones_event\", \"origin\": \"lukes-test-service\", \"requestid\": \"some-requestid:123456\", \"tstamp_usec\": 1400604684112158, \"dests\": [\"K8S_CAPTAINS_LOG\",\"YOWZA\"]}\n","stream":"stdout","time":"2018-01-29T22:06:25.107708785Z"}]'