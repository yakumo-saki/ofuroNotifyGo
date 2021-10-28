# ofuroNotifyGo

Bath notifier for aws lambda
Using golang, DynamoDB

## Usage

## Deploy

### First time

* Install lambroll (https://github.com/fujiwara/lambroll)
* Create lambda function in AWS Console
* lambroll --init
* Check `function.json` for Description, MemorySize, Timeout

### Deploy

* Exec deploy.sh

## Development

### First time setup

* Install docker
* cp testrun.sh debugrun.sh
* edit debugrun.sh (environment values)

### Run

* Run dynamodb-local conatiner `docker run --rm -d --name dynamodb -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -inMemory -sharedDb`
* exec this by `debugrun.sh`
