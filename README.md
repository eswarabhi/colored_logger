# Colored Logger
> A log activity that will write colored logs to console. 

![Alt text](https://i.imgur.com/XxavNUc.png)

## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```text
flogo install github.com/eswarabhi/colored_logger
```
## Schema
Inputs and Outputs:
```json
{
  "input":[
    {
      "name": "message",
      "type": "string",
      "value": ""
    },
    {
      "name": "flowInfo",
      "type": "boolean",
      "value": false
    },
    {
      "name": "addToFlow",
      "type": "boolean",
      "value": false
    },
    {
      "name": "level",
      "type": "string",
      "required": true,
      "allowed" : ["Trace", "Debug", "Info", "Print", "Warn", "Error", "Fatal"]
    }
  ],
  "output": [
    {
      "name": "message",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting         | Required      | Allowed Values  | Description |
| :-------------: |:-------------:| :---------------| ------------|
| message         | false         | string          |The message to log|
| flowInfo        | false         | boolean         |If set to true this will append the flow information to the log message|
| addToFlow       | false         | boolean         |If set to true this will add the log message to the 'message' output of the activity and make it available in further activities|
| level           | true          | ["Trace", "Debug", "Info", "Print","Warn", "Error", "Fatal"]| Allows us to set the log level manually|

## Examples
Below stated example will write "THIS IS AN ERROR!" in red color to console.
```json
      {
            "id": "colored_logger_6",
            "name": "Colored Logger (5)",
            "description": "ea-Simple Log Activity",
            "activity": {
              "ref": "github.com/eswarabhi/colored_logger",
              "input": {
                "message": "THIS IS AN ERROR!",
                "flowInfo": false,
                "addToFlow": false,
                "level": "Error"
              }
            }
        }
```
Below stated example will write "Find the trace" in cyan color to console.
```json
          {
            "id": "colored_logger_7",
            "name": "Colored Logger (6)",
            "description": "Activity that prints colored logs",
            "activity": {
              "ref": "github.com/eswarabhi/colored_logger",
              "input": {
                "message": "Find the trace",
                "flowInfo": false,
                "addToFlow": false,
                "level": "Trace"
              }
            }
          }
```
