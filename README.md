# XLD Command Line Interface
 ## _EXPERIMENTAL_

## xld login

Provide login details for XL Deploy server. Credentials are stored base64 encoded in a .xld file in the root of your user profile for reuse in subsequent requests.

Usage:

 - xld login <server> <username> <password>

Example:

 - xld login localhost:4516 admin $ecr3tP@ssw0rd


## xld deploy

Executes a deployment, either initial or update. If you need to make a distinction, use xld initial or xld update instead.

Usage:

 - xld deploy <app id> <env id>

 - xld deploy <app id> <env id> -orchestrator <orchestrator(s)>

Examples:

 - xld deploy app/MyApp/2.0 env/MyEnv

 - xld deploy app/MyComp/3.0 env/MyEnv -orchestrator parallel-by-container parallel-by-composite-package


## xld undeploy

Uninstall an application from a container.

Usage:

 - xld undeploy <deployed app id>

Example:

 - xld undeploy env/MyEnv/MyApp



## xld initial

Executes an initial deployment. This explicitly does *not* work for upgrade deployments. Use xld deploy to deploy regardless of intitial/upgrade.

Usage:

 - xld initial <app id> <env id>

 - xld initial <app id> <env id> -orchestrator <orchestrator(s)>

Examples:

 - xld initial app/MyApp/1.0 env/MyEnv

 - xld initial app/MyComp/1.0 env/MyEnv -orchestrator parallel-by-container parallel-by-composite-package


## xld update 

Executes an update deployment. This explicitly does *not* work for initial deployments. Use xld deploy to deploy regardless of intitial/upgrade.

Usage:

 - xld update <app id> <env id>

Examples:

 - xld update app/MyApp/2.0 env/MyEnv


## xld plan-initial

Show the steps for an initial deployment without executing. For execution, see xld initial

Usage:

 - xld plan-initial <app id> <env id>

Example:

 - xld plan-initial app/MyApp/1.0 env/MyEnv


## xld plan-update

Show the steps for an update deployment without executing. For execution, see xld update

Usage:

 - xld plan-update <app id> <env id>

Example:

 - xld plan-update app/MyApp/2.0 env/MyEnv


## xld create

Create items in XL Deploy from command line.

#### Basic usage:

xld create <type> <id> -<key> <value(s)>...

#### Advanced usage:

 - To enter key-value pairs, you can pipe JSON or CSV as input:

	<output json map> | xld create <type> <id> -<key> stdin:json

	<ouput csv file> | xld create <type> <id> -<key> stdin:csv

 - To enter the full content, you can pipe JSON:

	<output json map> | xld create <type> <id> stdin

 - To enter the full content, type and ID, you can pipe JSON:

	<output json map> | xld create stdin

Examples:

xld create overthere.LocalHost inf/MyServer -os UNIX -tags one two three -temporaryDirectoryPath /tmp

xld create dict env/MyDict -entries key1=value1 key2=value2

xld create env env/MyEnv -members inf/MyServer -dictionaries env/MyDict


Take a file myentries.json with the following content:

{
	"key1": "value1",
	"key2": "value2"
}

and type:

cat myentries.json | xld create dict env/MyDict -entries stdin:json

Take a file mydict.json with the following content:

{
	"entries": {
		"key1": "value1",
		"key2": "value2"
	}
}

and type:

cat myentries.json | xld create dict env/MyDict stdin:json

Take a file myitem.json with the following content:

{
    "content": {
        "entries": {
            "key1": "value1",
            "key2": "value2"
        }
    },
    "id": "env/MyDict",
    "type": "dict"
}

and type:

cat myentries.json | xld create stdin:json

Abbreviations

XLD allows the following abbreviations for item types:

env -> udm.Environment
dict -> udm.Dictionary
dir -> udm.Directory

XLD allows the following abbreviations for ID paths:

app -> Applications
env -> Environments
inf -> Infrastructure
conf -> Configuration



## xld read

Read a configuraton item from the repository and output JSON format.

Usage:

- xld read <id>

Examples:

- xld read env/MyEnv
- xld read inf/MyServer/MyTomcat

Note: env and inf are abbreviations for Environments and Infrastructure. You can also use the full names:

- xld read Infrastructure/MyServer

You can also use the abbreviation "latest" to automatically find the newest version of an application:

- xld read app/MyApp/latest


## xld modify

Not yet implemented


## xld remove

Delete an item from the repository.

Usage:

 - xld remove <item id(s)

Examples:

 - xld remove env/MyEnv

 - xld remove $(xld list -like %My%)


## xld list

Search for items in the repository

Usage:

 - xld list <item id> -type <type> -like <query> -before <time indication> -after <..> -page <##> -pagesize <##>

Example:

For all the direct children of Applications, type:

 - xld list app

For all the direct children and descendants of Applications, type:

 - xld list app/*

For all items with "Csv" in the name, type:

 - xld list -like %Csv%


## xld describe 

Print properties and property type for item type(s).

Usage:

 - xld describe <item type(s)

Examples:

 - xld describe jee.War

 - xld describe tomcat.Server udm.Directory

 - xld describe $(xld types | grep tomcat)


## xld types

Prints the list of item types installed in the XL Deploy server you connected to

Usage:
 
 - xld types

Examples

 - xld types | grep tomcat


