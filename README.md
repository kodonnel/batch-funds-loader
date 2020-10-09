# Batch Funds Loader

batch-funds-loader is a Golang program that accepts or declines attempts to load funds into customers' accounts in real-time.

## Prerequisites

    golang version with go modules (developed and tested on version 1.13.4)

## Installation

    $ git clone git@github.com:kodonnel/batch-funds-loader.git
    $ cd batch-funds-loader
    $ make build
    $ ./batch-funds-loader

## Usage

    $ ./batch-funds-loader [-h] [-v] [-i inputfile] [-o outputfile] 

    h - help
    v - verbose
    i - inputfile, default input.txt
    o - outputfile, default output.txt

## Example

### *command*

    $ ./batch-funds-loader -i inputfile.txt -o outputfile.txt

### *inputfile.txt contents (before running command)*

    {"id":"15888","customer_id":"528","load_amount":"$4000.47","time":"2000-01-02T00:00:00Z"}
    {"id":"15889","customer_id":"528","load_amount":"$4000.47","time":"2000-01-03T00:00:00Z"}
    {"id":"15890","customer_id":"528","load_amount":"$4000.47","time":"2000-01-04T00:00:00Z"}
    {"id":"15891","customer_id":"528","load_amount":"$4000.47","time":"2000-01-05T00:00:00Z"}
    {"id":"15892","customer_id":"528","load_amount":"$4000.47","time":"2000-01-06T00:00:00Z"}
    {"id":"15893","customer_id":"528","load_amount":"$4000.47","time":"2000-01-07T00:00:00Z"}
    {"id":"15895","customer_id":"528","load_amount":"$4000.47","time":"2000-01-09T00:00:00Z"}
    {"id":"30081","customer_id":"154","load_amount":"$1413.18","time":"2000-01-01T01:01:22Z"}
    {"id":"26540","customer_id":"426","load_amount":"$404.56","time":"2000-01-01T02:02:44Z"}
    {"id":"26544","customer_id":"426","load_amount":"$5000.56","time":"2000-01-01T02:02:44Z"}
    {"id":"26541","customer_id":"426","load_amount":"$404.56","time":"2000-01-01T02:02:44Z"}
    {"id":"26542","customer_id":"426","load_amount":"$404.56","time":"2000-01-01T02:02:44Z"}
    {"id":"26543","customer_id":"426","load_amount":"$404.56","time":"2000-01-01T02:02:44Z"}
    {"id":"10694","customer_id":"1","load_amount":"$785.11","time":"2000-01-01T03:04:06Z"}
    {"id":"10694","customer_id":"1","load_amount":"$400.11","time":"2000-01-01T03:04:06Z"}

### *outputfile.txt contents (after running command)*

    {"id":"15888","customer_id":"528","accepted":true}
    {"id":"15889","customer_id":"528","accepted":true}
    {"id":"15890","customer_id":"528","accepted":true}
    {"id":"15891","customer_id":"528","accepted":true}
    {"id":"15892","customer_id":"528","accepted":true}
    {"id":"15893","customer_id":"528","accepted":false}
    {"id":"15895","customer_id":"528","accepted":false}
    {"id":"30081","customer_id":"154","accepted":true}
    {"id":"26540","customer_id":"426","accepted":true}
    {"id":"26544","customer_id":"426","accepted":false}
    {"id":"26541","customer_id":"426","accepted":true}
    {"id":"26542","customer_id":"426","accepted":true}
    {"id":"26543","customer_id":"426","accepted":false}
    {"id":"10694","customer_id":"1","accepted":true}

## Assumptions

If the given output file does not exist, it will be created. Otherwise it will be appended to.
load_amount values are provided in CAD.


load example:
    {
        "id": "1234",
        "customer_id": "1234",
        "load_amount": "$123.45",
        "time": "2018-01-01T00:00:00Z"
    }

For a valid load:
- id will be a string consisting only of numeric characters
- id will be greater than or equal to 1 and less than 4294967295 (the max for golangs uint32 data type)
- customer_id will be a string consisting only of numeric characters
- customer_id will be greater than or equal to 1 and less than 4294967295 (the max for golangs uint32 data type)
- load_amount will always match the pattern ^\$\d+\.\d{2}$
- load amount will not be negative and will be less than 4294967295 (the max for golangs uint32 data type)
- time will always be provided in the ISO 8601 format in the America/Toronto timezone

Invalid load funds requests will be ignored.

Loads that were not accepted do not count against the maximum for the follow requirements:
- A maximum of $5,000 can be loaded per day. 
- A maximum of $20,000 can be loaded per week.
- A maximum of 3 loads can be performed per day, regardless of amount.


Both accepted and not accepted loads (but not ignored loads) apply for the following requirement:
- If a load ID is observed more than once for a particular user, all but the first instance can be ignored.

Input arrives in ascending chronological order.

## Testing

### Run the tests
    $ make test

### Coverage report
    $ make coverage

## Documentation

    $ go doc [pkg]

    packages include data, utils, and handlers

## Future Enhancements

- performance tests
- parameterized tests
- real database
- caching the days loads for quick access
- generated markdown apidocs

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)