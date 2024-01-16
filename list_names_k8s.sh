#!/bin/bash

POD="dnsutils"

printf "=========== Create a pod ==================\n"
kubectl run ${POD} --image=ubuntu --restart=Never -- /bin/sh -c 'while true; do echo hello;sleep 10;done'

# Wait for the pod "ubuntu" to contain the status condition of type "Ready"
kubectl wait --for=condition=Ready pod/${POD}

# Save a sorted list of IPs of all of the k8s SVCs:
rm -f ips podips
kubectl get svc -A -o custom-columns=IP:.spec.clusterIP| grep -Ew "([0-9]{1,3}\.){3}[0-9]{1,3}" >> ./ips
kubectl get pods -A -o custom-columns=IP:.status.podIP| grep -Ew "([0-9]{1,3}\.){3}[0-9]{1,3}" >> ./podips

# download secret go program
#curl -sO https://github.com/HeliconRichard/k8s-dns-lookups/blob/main/dns
#chmod +x dns

# Copy the ip list to owr Ubuntu pod:
kubectl cp ./ips ${POD}:/
kubectl cp ./podips ${POD}:/
kubectl cp ./dns ${POD}:/

printf "\n=========== Print all k8s SVC DNS records ===========\n"
kubectl exec -it ${POD} -- ./dns ips

printf "\n=========== Print all k8s pod DNS records ===========\n"
kubectl exec -it ${POD} -- ./dns podips

printf "\n=========== Cleanup ===========\n"
nohup kubectl delete po ${POD}
rm -f ./ips ./podips
exit 0

