#!/bin/bash

if [[ $1 == "start" ]]; then
	kubectl apply -f <(istioctl kube-inject -f flight-istio-deployment.yaml) 
	kubectl apply -f <(istioctl kube-inject -f flight-gateway.yaml)
	kubectl apply -f <(istioctl kube-inject -f flight-route.yaml)
elif [[ $1 == "stop" ]]; then
	kubectl delete -f <(istioctl kube-inject -f flight-istio-deployment.yaml)
	kubectl delete -f <(istioctl kube-inject -f flight-gateway.yaml)
	kubectl delete -f <(istioctl kube-inject -f flight-route.yaml)
elif [[ $1 == "check" ]]; then
	istioctl get gateway
	istioctl get virtualservices
elif [[ $1 == "update" ]]; then
	istioctl replace -f $2
else
	echo "Please specify [start, stop, check, update]"
	exit 1
fi
