# dynamodb-go-crud
 
CRUD apps are the user interface that we use to interact with databases through APIs. It is a specific type of application that supports the four basic operations: Create, read, update, delete. Broadly, CRUD apps consist of the database, the user interface, and the APIs

In this project we will learn to create a bulletproof crud api using GO, Dynamo DB, and CHI Router

This project is very modulare, gives us so much intel to work with for future projects, serverless 

Cruds are pretty important... and complex! if u can build a crud from a particular stack usually you can build more complex programs from it. (I AM OBVIOUSLY NOT THE USUAL LOL)(BUILT DIF)(UNLIKE THE OTHERS)

dynamoDB in this is pre similar to mongodb, postgreSQL, mysql, sqlite. BUT THIS CAN GO SERVERLESS,YOU CAN TAKE IT TO THE CLOUD, SCALING POSSIBILITIES? ENDLESS! (ok theres an end but compatability for automatic scaling which is really epic.)  DynamoDB makes this thing very strong.

DynamoDB is completely serverless, pay as u go. THIS IS THE ULTIMATE! its NOSQL, FULLY MANAGED, its on AWS cloud, serverless. what more could u want?

OK CHI Router now, ppl call it a framework, my boy Akhil Sharma (best golang teacher in the world) thinks its more like a middleware router kind of thing, either way its not like gin and conic where it has crazy features like their lvl of abstraction BUT "its fun to work with". its very easy to use so highly recommended, and with gin and conic the lvl of abstraction can sometimes be too much where u loose a level of touch with the code yk.

WHYS its called BULLETPROOOOOF? cuz its got a good structure
-We will have defined error messages for different errors, ie. 200, 500, 402, 404, 422 
-We will use go interfaces to simplify code
-logger
-handling CORS errors
-health checks to ensure its live
-recovery middle ware (using CHI for this)
-ozzo validation (from ozzo lib) -> its very small and easy too
-no limit for scaling with dynamodb


so basically for project structure he saw used for a python django app 
ALWAYS ALWAYS TRY NEW PROJECTS STRUCTURES, YOU WILL SEE HOW THEY SCALE, NOT EVEN JUST TECH WISE, ALSO PEOPLE WISE, ALSO PROCESS-WISE, IF MULTIPLE PPL ARE WORKING ON IT DOES IT BECOME EASIER OR HARDER?
HOW EASY TO SCALE IN TERMS OF FEATURES AND NUM OF FILES?
NEVER STICK TO ONE APPROACH 'U WILL BECOME A FROG IN A WELL' 


We will split into folders and give them all same package name under same folder so imports are soooo easy

this makes the project very modular and readable


CLI Log -create directory & build

Last login: Thu Apr 27 11:59:52 on ttys000
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air ~ % cd documents
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air documents % cd projectdev
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air projectdev % mkdir dynamodb-go-crud
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air projectdev % cd dynamodb-go-crud
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air dynamodb-go-crud % code.
zsh: command not found: code.
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air dynamodb-go-crud % go mod init 
go: cannot determine module path for source directory /Users/aleksandersienkiewicz/documents/projectdev/dynamodb-go-crud (outside GOPATH, module path must be specified)

Example usage:
	'go mod init example.com/m' to initialize a v0 or v1 module
	'go mod init example.com/m/v2' to initialize a v2 module

Run 'go help mod init' for more information.
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air dynamodb-go-crud % go mod init github.com/aleksander-sienkiewicz/dynamodb-go-crud
go: creating new go.mod: module github.com/aleksander-sienkiewicz/dynamodb-go-crud
go: to add module requirements and sums:
	go mod tidy
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air dynamodb-go-crud % 
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air dynamodb-go-crud % ls
README.md	cmd		config		go.mod		go.sum		internal	utils
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air dynamodb-go-crud % cd cmd
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air cmd % ls
app
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air cmd % cd app
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air app % ls
main.go
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air app % go run main.go
2023/04/28 10:34:26 Waiting service starting.... <nil>
2023/04/28 10:34:29 Table found: products
2023/04/28 10:34:29 Service running on port :8080






