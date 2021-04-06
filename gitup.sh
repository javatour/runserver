#! /bin/bash

git add ./*
git commit -m $1
git push origin main

log=$(git log | grep commit)
echo ${log}
shortlog=$(echo ${log} | tr " " "\n")
n=0
for elem in ${shortlog}
do
	n=$((${n}+1))
	if [ ${n} -eq 2 ] ; then
		echo "commit log is ${elem}"	
	fi	
done

