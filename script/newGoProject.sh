#!/bin/bash

#需要创建的文件夹列表(正常服务器项目)
folderDef=("api" "build" "cmd" "configs" "docs" "init" "internal" "pkg" "scripts" "test" "tools")

#需要创建的文件夹列表(web项目)
folderWeb=("api" "build" "cmd" "configs" "docs" "init" "internal" "pkg" "scripts" "test" "tools" "web")

#需要创建的文件夹列表(完整列表)
folderAll=("api" "assets" "build" "cmd" "configs" "deployments" "docs" "examples" "githooks" "init" "internal" "pkg" "scripts" "test" "third_party" "tools" "vendor" "web" "website")

selectF=$2

# 创建文件夹
CreateFile(){

	if [ ! -d "${1}" ]; then
		mkdir "${1}"
	fi

	echo "create folder ${1}"

	# shellcheck disable=SC2206
	folder=(${folderDef[*]})

	if [[ 1 == ${selectF} ]]; then
		# shellcheck disable=SC2206
		folder=(${folderWeb[*]})
		echo "select folderWeb create ! "
	elif [[ 2 == ${selectF} ]]; then
		# shellcheck disable=SC2206
		folder=(${folderAll[*]})
		echo "select folderAll create ! "
	else
		echo "select folderDef create ! "
	fi


	for f in ${folder[*]}; do
		if [ ! -d "${1}/${f}" ]; then
			mkdir "${1}/${f}"
			echo "create folder ${1}/${f} "
		else
			echo "exist folder ${1}/${f} "
		fi
	done
	echo "create ${1} success !"
}

# 输入验证
Input(){

	# shellcheck disable=SC2236
	if [ ! -n "${1}" ]; then
		echo "please input projectName !"
		echo "example: ./newGoProject.sh test"
		exit
	fi

	# shellcheck disable=SC2162
	read -p "Please confirm whether to create ${1},y/n?" confirm

	if [[ "y" != $confirm && "Y" != $confirm ]]; then
		echo "cancel create !"
		exit
	fi

	CreateFile "${1}"
}

Input "${1}"
