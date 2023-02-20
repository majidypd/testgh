#! /bin/bash

exitMessage() {
   echo "${2-exiting}"
   exit "${1}"
}



tag=${GITHUB_TAG-}
echo "recived:"
echo $tag
# if [[ $tag =~ ^v[0-9]{1}\.[0-9]{1,2}\.[0-9]{1,3}$ ]]; then
#    echo "Deploying ${tag} into production account in ${AWS_REGION_PROD}..."
# else
#    exitMessage 0 "not a valid tag, not deploying"
# fi

# echo "after xxxxx"
# echo $tag

go run main.go